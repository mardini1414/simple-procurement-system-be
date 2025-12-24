package model

import (
	"time"

	"github.com/google/uuid"
)

type Item struct {
	ID    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name  string    `gorm:"size:150;not null" json:"name"`
	Stock int       `gorm:"not null" json:"stock"`
	Price int64     `gorm:"not null" json:"price"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	PurchaseDetails []PurchasingDetail `gorm:"foreignKey:ItemID" json:"purchase_details,omitempty"`
}
