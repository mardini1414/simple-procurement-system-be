package model

import (
	"time"

	"github.com/google/uuid"
)

type Supplier struct {
	ID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name    string    `gorm:"size:150;not null" json:"name"`
	Email   string    `gorm:"size:150;unique" json:"email,omitempty"`
	Address string    `gorm:"type:text" json:"address,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Purchases []Purchasing `gorm:"foreignKey:SupplierID" json:"purchases,omitempty"`
}
