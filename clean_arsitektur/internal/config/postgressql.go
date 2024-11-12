package config

import (
	"clean_arsitektur/internal/domain/entity"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbTimezone := os.Getenv("DB_TIMEZONE")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbTimezone)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("failed to connect to the database: " + err.Error())
	}

	sqlDB, err := DB.DB()
	if err != nil {
		panic("failed to connect to the database: " + err.Error())
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(5 * 60)

	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		fmt.Println("Migration started...")
		RunPostgresMigration()
	}
}

func RunPostgresMigration() {
	err := DB.AutoMigrate(
		&entity.User{},
	)

	if err != nil {
		fmt.Printf("Migration failed: %v\n", err)
	} else {
		fmt.Println("Migration completed successfully!")
	}
}
