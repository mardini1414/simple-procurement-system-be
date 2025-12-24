package dto

import (
	"time"

	"github.com/google/uuid"
)

type (
	CreatePurchasingRequest struct {
		SupplierID uuid.UUID     `json:"supplier_id" validate:"required"`
		Items      []ItemRequest `json:"items" validate:"required,dive"`
	}

	ItemRequest struct {
		ID  uuid.UUID `json:"id" validate:"required"`
		Qty int       `json:"qty" validate:"required,gte=0"`
	}

	PurchasingJoinRow struct {
		PurchasingID uuid.UUID `json:"purchasing_id"`
		Date         time.Time `json:"date"`
		GrandTotal   int64     `json:"grand_total"`

		Username      string `json:"username"`
		SupplierName  string `json:"supplier_name"`
		SupplierEmail string `json:"supplier_email"`

		DetailID uuid.UUID `json:"detail_id"`
		ItemID   uuid.UUID `json:"item_id"`
		ItemName string    `json:"item_name"`
		Qty      int       `json:"qty"`
		SubTotal int64     `json:"sub_total"`
	}

	GetPurchasingResponse struct {
		ID            uuid.UUID `json:"id"`
		Username      string    `json:"username"`
		SupplierName  string    `json:"supplier_name"`
		SupplierEmail string    `json:"supplier_email"`
		Date          time.Time `json:"date"`
		GrandTotal    int64     `json:"grand_total"`
	}

	GetPurchasingDetailResponse struct {
		ID            uuid.UUID        `json:"id"`
		Username      string           `json:"username"`
		SupplierName  string           `json:"supplier_name"`
		SupplierEmail string           `json:"supplier_email"`
		Date          time.Time        `json:"date"`
		GrandTotal    int64            `json:"grand_total"`
		Details       []DetailResponse `json:"details"`
	}

	DetailResponse struct {
		ID       uuid.UUID `json:"id"`
		ItemName string    `json:"item_name"`
		Qty      int       `json:"qty"`
		SubTotal int64     `json:"sub_total"`
	}

	CreatePurchase struct {
		ID         uuid.UUID
		UserID     uuid.UUID
		SupplierID uuid.UUID
		Date       time.Time
		GrandTotal int64
	}

	CreatePurchaseDetail struct {
		PurchaseID uuid.UUID
		ItemID     uuid.UUID
		Qty        int
		SubTotal   int64
	}

	PurchaseWebhookPayload struct {
		Event      string          `json:"event"`
		PurchaseID uuid.UUID       `json:"purchase_id"`
		UserID     uuid.UUID       `json:"user_id"`
		SupplierID uuid.UUID       `json:"supplier_id"`
		Date       time.Time       `json:"date"`
		GrandTotal int64           `json:"grand_total"`
		Details    []DetailPayload `json:"details"`
	}

	DetailPayload struct {
		ID       uuid.UUID `json:"id"`
		Qty      int       `json:"qty"`
		SubTotal int64     `json:"sub_total"`
	}
)
