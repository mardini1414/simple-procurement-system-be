package main

import (
	"github.com/google/uuid"
	"github.com/mardini1414/simple-procurement-system-be/internal/config"
	database "github.com/mardini1414/simple-procurement-system-be/internal/database"
	"github.com/mardini1414/simple-procurement-system-be/internal/model"
	"github.com/mardini1414/simple-procurement-system-be/pkg"
)

func main() {
	cfg := config.LoadConfig()
	db := database.LoadDB(cfg)

	userId := uuid.New()

	user := model.User{
		ID:       userId,
		Username: "admin",
		Password: pkg.HashPassword("admin123"),
		Role:     "user",
	}

	db.Create(&user)
}
