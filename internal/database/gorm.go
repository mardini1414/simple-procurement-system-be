package database

import (
	"fmt"
	"log"

	"github.com/mardini1414/simple-procurement-system-be/internal/config"
	"github.com/mardini1414/simple-procurement-system-be/internal/dto"
	"gorm.io/gorm"

	"gorm.io/driver/postgres"
)

func LoadDB(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		cfg.DbHost,
		cfg.DbUser,
		cfg.DbPassword,
		cfg.DbName,
		cfg.DbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error loading database")
	}
	return db
}

func Paginate(req dto.Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := req.Page

		if page <= 0 {
			page = 1
		}

		pageSize := req.Size

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize

		return db.Offset(offset).Limit(pageSize)
	}

}
