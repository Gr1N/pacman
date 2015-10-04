package app

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"

	"github.com/revel/revel"

	"github.com/Gr1N/pacman/app/models"
)

var (
	DB *gorm.DB
)

func InitDB() {
	var (
		driver string
		spec   string
		found  bool
	)

	// Read configuration
	if driver, found = revel.Config.String("db.driver"); !found {
		revel.ERROR.Fatal("No db.driver found.")
	}
	if spec, found = revel.Config.String("db.spec"); !found {
		revel.ERROR.Fatal("No db.spec found.")
	}

	// Initializae `gorm`
	dbm, err := gorm.Open(driver, spec)
	if err != nil {
		revel.ERROR.Fatal(err)
	}

	DB = &dbm

	dbm.DB().Ping()
	dbm.DB().SetMaxIdleConns(10)
	dbm.DB().SetMaxOpenConns(100)
	dbm.SingularTable(true)

	// Migrate
	dbm.AutoMigrate(&models.User{})
}
