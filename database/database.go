package database

import (
	"github.com/WeDias/golang-test-api/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Conn *gorm.DB

func InitDatabase() {
	var err error
	Conn, err = gorm.Open(sqlite.Open("database/database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	Conn.AutoMigrate(&entities.Product{})
}
