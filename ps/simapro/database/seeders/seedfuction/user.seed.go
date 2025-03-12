package seeds

import (
	"fmt"
	"log"

	"github.com/gpjen/simapro/app/models"
	"github.com/gpjen/simapro/utils"
	"gorm.io/gorm"
)

func UserSeeder(db *gorm.DB) {
	defaultPass, _ := utils.HashPassword("123321")

	users := []models.User{
		{
			Name:      "Admin",
			Email:     "admin@mail.com",
			Password:  defaultPass,
			CreatedBy: createdBy,
			Roles:     []models.Role{{Name: "Admin", Description: "Admin role", Active: true, CreatedBy: createdBy}},
		},
		{
			Name:      "Production Manager",
			Email:     "productionmanager@mail.com",
			Password:  defaultPass,
			CreatedBy: createdBy,
			Roles:     []models.Role{{Name: "Production Manager", Description: "Production Manager role", Active: true, CreatedBy: createdBy}},
		},
		{
			Name:      "Operator",
			Email:     "user@mail.com",
			Password:  defaultPass,
			CreatedBy: createdBy,
			Roles:     []models.Role{{Name: "Operator", Description: "Operator role", Active: true, CreatedBy: createdBy}},
		},
	}

	for _, user := range users {
		result := db.Where("email = ?", user.Email).FirstOrCreate(&user)
		if result.Error != nil {
			log.Printf("❌ Error creating user %s: %v", user.Email, result.Error)
		}
	}

	fmt.Println("✅ User seeder executed successfully!")
}
