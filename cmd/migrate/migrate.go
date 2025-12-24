package main

import (
	"github.com/mardini1414/simple-procurement-system-be/internal/config"
	database "github.com/mardini1414/simple-procurement-system-be/internal/database"
)

func main() {
	cfg := config.LoadConfig()
	db := database.LoadDB(cfg)

	database.Migrate(db)
}
