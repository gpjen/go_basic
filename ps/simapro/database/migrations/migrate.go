package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gpjen/simapro/app/models"
	"github.com/gpjen/simapro/config"
	"github.com/gpjen/simapro/database"
	"github.com/joho/godotenv"
)

var modelList = []interface{}{
	&models.User{},
	&models.Role{},
	&models.Menu{},
	&models.RoleMenu{},
}

var modelPivotList = []string{
	"user_roles",
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	config.LoadConfig()

	database.ConnectDB()
	db := database.DB

	action := flag.String("action", "", "Migration action: up or down")
	flag.Parse()

	switch *action {
	case "up":
		fmt.Println("🕒 Running migrations...")

		err := db.AutoMigrate(modelList...)

		if err != nil {
			log.Fatalf("❌ Error performing migration: %v", err)
		}

		fmt.Println("✅ Migrations completed successfully!")

	case "down":
		fmt.Println("🕒 Dropping all tables...")

		err := db.Migrator().DropTable(modelList...)
		if err != nil {
			log.Fatalf("❌ Error dropping tables: %v", err)
		}

		for _, pivotTable := range modelPivotList {
			err := db.Migrator().DropTable(pivotTable)
			if err != nil {
				log.Fatalf("❌ Error dropping pivot table %s: %v", pivotTable, err)
			}
		}

		fmt.Println("✅ All tables dropped successfully!")

	default:
		fmt.Println("Invalid action. Use -action=up to migrate or -action=down to drop tables")
		os.Exit(1)
	}
}
