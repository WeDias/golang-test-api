package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Conn *gorm.DB

func InitDatabase() {
	var err error

	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v",
		os.Getenv("PG_HOST"), os.Getenv("PG_USER"), os.Getenv("PG_PASS"),
		os.Getenv("PG_DBNM"), os.Getenv("PG_PORT"),
	)

	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	Conn, err = gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		panic("failed to connect database")
	}
}
