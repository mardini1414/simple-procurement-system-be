package model

import (
	"time"

	"github.com/google/uuid"
)

type Purchasing struct {
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`

	Date time.Time `gorm:"not null"`

	SupplierID uuid.UUID
	Supplier   Supplier

	UserID uuid.UUID
	User   User

	GrandTotal int64 `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time

	Details []PurchasingDetail `gorm:"foreignKey:PurchasingID"`
}
