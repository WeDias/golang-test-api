package database

import (
	"fmt"
	"os"

	"github.com/WeDias/golang-test-api/utils"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var Conn *gorm.DB

func init() {
	godotenv.Load()

	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v",
		os.Getenv("PG_HOST"), os.Getenv("PG_USER"), os.Getenv("PG_PASS"),
		os.Getenv("PG_DBNM"), os.Getenv("PG_PORT"),
	)

	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Silent),
	}

	var err error
	utils.InfoLog.Println("connecting to database")
	Conn, err = gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		utils.ErrorLog.Panicln(err.Error())
	}
}
