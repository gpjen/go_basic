package repositories

import (
	"github.com/gpjen/simapro/app/models"
	"github.com/gpjen/simapro/database"
	"gorm.io/gorm"
)

type RoleMenuRepository interface {
	FindAllByRoleID(roleID uint) []models.RoleMenu
	Update(roleMenus []*models.RoleMenu) error
}

type menumenuRepository struct {
	db *gorm.DB
}

func NewRoleMenuRepository() RoleMenuRepository {
	return &menumenuRepository{db: database.DB}
}

func (r *menumenuRepository) FindAllByRoleID(roleID uint) []models.RoleMenu {
	var roleMenus []models.RoleMenu
	r.db.Where("role_id = ?", roleID).Find(&roleMenus)
	return roleMenus
}

func (r *menumenuRepository) Update(roleMenus []*models.RoleMenu) error {
	return r.db.Save(roleMenus).Error
}
