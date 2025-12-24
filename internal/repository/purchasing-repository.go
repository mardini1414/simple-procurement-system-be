package repository

import (
	"github.com/google/uuid"
	"github.com/mardini1414/simple-procurement-system-be/internal/dto"
	"github.com/mardini1414/simple-procurement-system-be/internal/model"
	"gorm.io/gorm"
)

type PurchasingRepository struct {
	db *gorm.DB
}

func NewPurchasingRepository(db *gorm.DB) *PurchasingRepository {
	return &PurchasingRepository{
		db: db,
	}
}

func (r *PurchasingRepository) Create(tx *gorm.DB, p dto.CreatePurchase) error {
	purchase := model.Purchasing{
		ID:         p.ID,
		UserID:     p.UserID,
		SupplierID: p.SupplierID,
		Date:       p.Date,
		GrandTotal: p.GrandTotal,
	}

	if err := tx.Create(&purchase).Error; err != nil {
		return err
	}

	return nil
}

func (r *PurchasingRepository) UpdateGrandTotal(tx *gorm.DB, id uuid.UUID, grandTotal int64) error {
	var purchase model.Purchasing
	return tx.Model(&purchase).Where("id = ?", id).Update("grand_total", grandTotal).Error
}

func (r *PurchasingRepository) CreateDetail(tx *gorm.DB, pd dto.CreatePurchaseDetail) error {
	detail := model.PurchasingDetail{
		PurchasingID: pd.PurchaseID,
		ItemID:       pd.ItemID,
		Qty:          pd.Qty,
		SubTotal:     pd.SubTotal,
	}

	if err := tx.Create(&detail).Error; err != nil {
		return err
	}

	return nil
}

func (r *PurchasingRepository) FindAll() (*[]dto.GetPurchasingResponse, error) {
	var purchasings []dto.GetPurchasingResponse

	err := r.db.Table("purchasings").
		Select("purchasings.id, users.username, suppliers.name as supplier_name, suppliers.email as supplier_email, purchasings.date, purchasings.grand_total").
		Joins("left join users on users.id = purchasings.user_id").
		Joins("left join suppliers on suppliers.id = purchasings.supplier_id").
		Scan(&purchasings).Error

	if err != nil {
		return nil, err
	}

	return &purchasings, nil
}

func (r *PurchasingRepository) FindById(id uuid.UUID) ([]dto.PurchasingJoinRow, error) {
	var rows []dto.PurchasingJoinRow
	err := r.db.Table("purchasings p").
		Select(`
		p.id AS purchasing_id,
		p.date,
		p.grand_total,
		u.username ,
		s.name AS supplier_name,
		s.email AS supplier_email,
		d.id AS detail_id,
		d.item_id,
		i.name AS item_name,
		d.qty,
		d.sub_total
	`).
		Joins("LEFT JOIN users u ON u.id = p.user_id").
		Joins("LEFT JOIN suppliers s ON s.id = p.supplier_id").
		Joins("LEFT JOIN purchasing_details d ON d.purchasing_id = p.id").
		Joins("LEFT JOIN items i ON i.id = d.item_id").
		Where("p.id = ?", id).
		Scan(&rows).Error

	if err != nil {
		return nil, err
	}

	return rows, nil
}
