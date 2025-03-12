package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gpjen/simapro/config"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := config.AppConfig.DB_URL
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	log.Println("Database Connected successfully")

}
