package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Conn *gorm.DB

func InitDatabase() {
	var err error
	dsn := "host=postgres-go user=postgres password=postgres dbname=postgres port=5432"
	Conn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}
