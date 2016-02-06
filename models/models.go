package models

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

	"github.com/Gr1N/pacman/modules/settings"
)

var (
	g *gorm.DB
)

// Model represents base model.
type Model struct {
	ID        int64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Init initializes application models.
func Init() {
	dbm, err := gorm.Open(settings.S.DB.Driver, settings.S.DB.Spec)
	if err != nil {
		panic(err)
	}

	dbm.DB().Ping()
	dbm.DB().SetMaxIdleConns(settings.S.DB.MaxIdleConns)
	dbm.DB().SetMaxOpenConns(settings.S.DB.MaxOpenConns)
	dbm.LogMode(settings.S.DB.LogMode)
	// dbm.SetLogger()

	dbm.AutoMigrate(
		&User{},
		&Token{},
		&Service{},
		&Repo{})

	g = &dbm
}
