package database

import (
	"log"

	"github.com/mardini1414/simple-procurement-system-be/internal/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.User{},
		&model.Supplier{},
		&model.Item{},
		&model.Purchasing{},
		&model.PurchasingDetail{},
	)

	if err != nil {
		log.Fatal("Migate failed", err.Error())
	}

}
