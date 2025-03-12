package seeds

import (
	"fmt"
	"log"

	"github.com/gpjen/simapro/app/models"
	"gorm.io/gorm"
)

var createdBy = "seeder"

func MenuSeeder(db *gorm.DB) {
	parentMenus := []models.Menu{
		{Name: "Dashboard", Route: "/dashboard", Icon: "", Order: 1},
		{Name: "Product", Route: "/product", Icon: "", Order: 2},
		{Name: "Work Order", Route: "/work-order", Icon: "", Order: 3},
		{Name: "Master Data", Route: "", Icon: "", Order: 4},
		{Name: "Laporan", Route: "", Icon: "", Order: 5},
	}

	for i := range parentMenus {
		menu := parentMenus[i]

		result := db.Where("name = ?", menu.Name).FirstOrCreate(&menu)
		if result.Error != nil {
			log.Printf("❌ Error creating menu %s: %v", menu.Name, result.Error)
		}

		parentMenus[i] = menu
	}

	masterDataID := parentMenus[3].ID

	childMenus := []models.Menu{
		{Name: "User", Route: "/users", IdParent: &masterDataID, Order: 1},
		{Name: "Role", Route: "/roles", IdParent: &masterDataID, Order: 2},
		{Name: "Permission", Route: "/permissions", IdParent: &masterDataID, Order: 3},
	}

	for _, menu := range childMenus {
		result := db.Where("name = ?", menu.Name).FirstOrCreate(&menu)
		if result.Error != nil {
			log.Printf("❌ Error creating user %s: %v", menu.Name, result.Error)
		}
	}

	fmt.Println("✅ Menus seeder executed successfully!")
}
