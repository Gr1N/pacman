package app

import (
	gorm "github.com/Gr1N/revel-gorm/app"

	"github.com/Gr1N/pacman/app/models"
)

func initDB() {
	// Initialize GORM...
	dbm := gorm.InitDB()

	// ...and migrate
	dbm.AutoMigrate(&models.User{}, &models.Service{})
}
