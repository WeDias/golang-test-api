package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Conn *gorm.DB

func InitDatabase() {
	var err error
	dsn := fmt.Sprintf(
		"host=%s user=postgres password=postgres dbname=postgres port=5432",
		os.Getenv("PG_HOST"))
	Conn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}
