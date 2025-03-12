package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gpjen/simapro/config"
	"github.com/gpjen/simapro/database"
	seeds "github.com/gpjen/simapro/database/seeders/seedfuction"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var availableSeeders = []struct {
	Name   string
	Seeder func(*gorm.DB)
}{
	{"userSeeder", seeds.UserSeeder},
	{"menuSeeder", seeds.MenuSeeder},
	{"roleMenuSeeder", seeds.RoleMenuSeeder},
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	config.LoadConfig()

	database.ConnectDB()
	db := database.DB

	seedName := flag.String("seed", "", "Specific seeder to run (leave empty to run all seeders)")
	flag.Parse()

	fmt.Println("Running seeders...")

	if *seedName == "" {
		for _, s := range availableSeeders {
			s.Seeder(db)
		}
	} else {
		found := false
		for _, s := range availableSeeders {
			if s.Name == *seedName {
				s.Seeder(db)
				fmt.Printf("✅ %s completed!\n", s.Name)
				found = true
				break
			}
		}

		if !found {
			fmt.Printf("⚠️ Seeder '%s' not found!\n", *seedName)
			fmt.Println("📝 Available seeders:")
			for _, s := range availableSeeders {
				fmt.Printf(" - %s\n", s.Name)
			}
		}
	}
}
