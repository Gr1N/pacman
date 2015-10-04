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

	maxIdleConns := revel.Config.IntDefault("db.max_idle_conns", 10)
	maxOpenConns := revel.Config.IntDefault("db.max_open_conns", 100)

	// Initialize `gorm`
	dbm, err := gorm.Open(driver, spec)
	if err != nil {
		revel.ERROR.Fatal(err)
	}

	DB = &dbm

	dbm.DB().Ping()
	dbm.DB().SetMaxIdleConns(maxIdleConns)
	dbm.DB().SetMaxOpenConns(maxOpenConns)
	dbm.SingularTable(true)

	// Migrate
	dbm.AutoMigrate(&models.User{})
}
