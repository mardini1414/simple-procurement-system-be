package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mardini1414/simple-procurement-system-be/internal/config"
	"github.com/mardini1414/simple-procurement-system-be/internal/dto"
	"github.com/mardini1414/simple-procurement-system-be/internal/repository"
	"github.com/mardini1414/simple-procurement-system-be/pkg"
	"gorm.io/gorm"
)

type PurchasingService struct {
	r        *repository.PurchasingRepository
	itemRepo *repository.ItemRepository
	db       *gorm.DB
	cfg      *config.Config
}

func NewPurchasingService(r *repository.PurchasingRepository,
	itemRepo *repository.ItemRepository,
	db *gorm.DB,
	cfg *config.Config) *PurchasingService {
	return &PurchasingService{
		r:        r,
		itemRepo: itemRepo,
		db:       db,
		cfg:      cfg,
	}
}

func (s *PurchasingService) Create(req dto.CreatePurchasingRequest, c *fiber.Ctx) error {
	return s.db.Transaction(func(tx *gorm.DB) error {

		if len(req.Items) == 0 {
			return pkg.NewApiError(fiber.StatusBadRequest, "Item tidak boleh kosong", nil)
		}

		userId, ok := c.Locals("user_id").(uuid.UUID)

		if !ok {
			return pkg.NewApiError(fiber.StatusBadRequest, "Invalid user id", nil)
		}

		var grandTotal int64
		purchaseId := uuid.New()

		purchase := dto.CreatePurchase{
			ID:         purchaseId,
			UserID:     userId,
			SupplierID: req.SupplierID,
			Date:       time.Now(),
			GrandTotal: grandTotal,
		}

		if err := s.r.Create(tx, purchase); err != nil {
			return err
		}

		details := []dto.DetailPayload{}

		for _, it := range req.Items {
			item, err := s.itemRepo.FindById(it.ID)

			if err != nil {
				return pkg.NewApiError(fiber.StatusNotFound, "Item tidak dtemukan", nil)
			}

			subTotal := int64(it.Qty) * item.Price
			grandTotal += subTotal

			purchaseDetail := dto.CreatePurchaseDetail{
				PurchaseID: purchaseId,
				ItemID:     it.ID,
				Qty:        it.Qty,
				SubTotal:   subTotal,
			}

			if err := s.r.CreateDetail(tx, purchaseDetail); err != nil {
				return err
			}

			if err := s.itemRepo.IncreaseStock(
				tx, it.ID, it.Qty,
			); err != nil {
				return err
			}

			details = append(details, dto.DetailPayload{
				ID:       item.ID,
				Qty:      it.Qty,
				SubTotal: subTotal,
			})
		}

		if err := s.r.UpdateGrandTotal(tx, purchaseId, grandTotal); err != nil {
			return err
		}

		webhookUrl := s.cfg.WebhookURL
		payload := dto.PurchaseWebhookPayload{
			Event:      "purchase.created",
			PurchaseID: purchase.ID,
			UserID:     userId,
			SupplierID: purchase.SupplierID,
			Date:       purchase.Date,
			GrandTotal: grandTotal,
			Details:    details,
		}

		go s.SendPurchaseWebhook(webhookUrl, payload)

		return nil

	})
}

func (s *PurchasingService) SendPurchaseWebhook(webhookURL string, payload dto.PurchaseWebhookPayload) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		webhookURL,
		bytes.NewBuffer(body),
	)

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return fmt.Errorf("webhook failed with status %d", resp.StatusCode)
	}

	return nil
}

func (s *PurchasingService) FindAll() (*[]dto.GetPurchasingResponse, error) {
	return s.r.FindAll()
}

func (s *PurchasingService) FindById(id uuid.UUID) (*dto.GetPurchasingDetailResponse, error) {
	rows, err := s.r.FindById(id)

	if rows == nil || err != nil {
		return nil, pkg.NewApiError(fiber.StatusNotFound, "Purchasing tidak di temukan", nil)
	}

	first := rows[0]

	purchasing := dto.GetPurchasingDetailResponse{
		ID:            first.PurchasingID,
		Date:          first.Date,
		GrandTotal:    first.GrandTotal,
		Username:      first.Username,
		SupplierName:  first.SupplierName,
		SupplierEmail: first.SupplierEmail,
	}

	for _, r := range rows {
		if r.DetailID == uuid.Nil {
			continue
		}

		purchasing.Details = append(purchasing.Details, dto.DetailResponse{
			ID:       r.DetailID,
			ItemName: r.ItemName,
			Qty:      r.Qty,
			SubTotal: r.SubTotal,
		})
	}

	return &purchasing, nil

}
