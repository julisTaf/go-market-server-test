package database

import (
	"Go-market-test/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// GlobalDB a global db object will be used across different packages
var GlobalDB *gorm.DB

// InitDatabase creates a sqlite db
func InitDatabase() (err error) {
	cfg, err := config.SetConfig()
	if err != nil {
		return
	}
	GlobalDB, err = gorm.Open(sqlite.Open(cfg.DataBase), &gorm.Config{})
	if err != nil {
		return
	}
	return
}
