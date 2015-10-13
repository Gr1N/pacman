package app

import (
	g "github.com/Gr1N/revel-gorm/app"

	"github.com/Gr1N/pacman/app/models"
)

func initDB() {
	// Initialize GORM...
	dbm := g.InitDB()

	// ...and migrate
	dbm.AutoMigrate(&models.User{}, &models.Service{})
}
