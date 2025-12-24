package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Username string    `gorm:"size:100;unique;not null"`
	Password string    `gorm:"size:255;not null"`
	Role     string    `gorm:"size:50;not null"`

	CreatedAt time.Time
	UpdatedAt time.Time

	Purchases []Purchasing `gorm:"foreignKey:UserID"`
}
