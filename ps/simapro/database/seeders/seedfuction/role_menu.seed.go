package seeds

import (
	"fmt"

	"github.com/gpjen/simapro/app/models"
	"gorm.io/gorm"
)

func RoleMenuSeeder(db *gorm.DB) {
	var adminRole models.Role
	var menus []models.Menu

	// Cek apakah role "Admin" ada
	if err := db.Where("name = ?", "Admin").First(&adminRole).Error; err != nil {
		fmt.Println("❌ Error getting admin role: ", err)
		return
	}

	// Ambil semua menu yang tersedia
	if err := db.Find(&menus).Error; err != nil {
		fmt.Println("❌ Error getting menus: ", err)
		return
	}

	// Jika tidak ada menu, hentikan proses
	if len(menus) == 0 {
		fmt.Println("❌ No menus found!")
		return
	}

	for _, menu := range menus {
		err := db.Where("role_id = ? AND menu_id = ?", adminRole.ID, menu.ID).FirstOrCreate(&models.RoleMenu{
			RoleID:    adminRole.ID,
			MenuID:    menu.ID,
			CanView:   true,
			CanShow:   true,
			CanCreate: true,
			CanUpdate: true,
			CanDelete: true,
			CreatedBy: createdBy,
		}).Error

		if err != nil {
			fmt.Println("❌ Error creating role_menu: ", err)
		}
	}

	fmt.Println("✅ RoleMenu seeder executed successfully!")
}
