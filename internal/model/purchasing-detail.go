package model

import (
	"time"

	"github.com/google/uuid"
)

type PurchasingDetail struct {
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`

	PurchasingID uuid.UUID
	Purchasing   Purchasing

	ItemID uuid.UUID
	Item   Item

	Qty      int   `gorm:"not null"`
	SubTotal int64 `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
