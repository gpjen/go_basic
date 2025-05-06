package repositories

import (
	"errors"

	"github.com/gpjen/simapro/app/dto"
	"github.com/gpjen/simapro/app/models"
	"github.com/gpjen/simapro/database"
	"gorm.io/gorm"
)

type RoleRepository interface {
	IsDuplicateName(name string) error
	CreateRole(role *models.Role) error
	GetRoles(page, perPage int) ([]*models.Role, int64, error)
	GetRole(roleId uint) (*models.Role, error)
	GetRoleByIds(ids []uint) ([]*models.Role, error)
	UpdateRole(role *models.Role) error
	DeleteRole(role *models.Role) error

	GetRoleMenuPermision(roleId uint, param *dto.ParamQueryMenuPermissionDTO) (*models.Role, error)
}

type roleRepository struct{}

func NewRoleRepository() RoleRepository {
	return &roleRepository{}
}

func (r *roleRepository) IsDuplicateName(name string) error {
	var role models.Role
	if err := database.DB.Where("name = ?", name).First(&role).Error; err != nil {
		return nil
	}
	return errors.New("name already exists")
}

func (r *roleRepository) CreateRole(role *models.Role) error {
	return database.DB.Create(role).Error
}

func (r *roleRepository) GetRoles(page, perPage int) ([]*models.Role, int64, error) {
	var roles []*models.Role
	var totalItems int64

	database.DB.Model(&models.Role{}).Count(&totalItems)

	if err := database.DB.Offset((page - 1) * perPage).Limit(perPage).Find(&roles).Error; err != nil {
		return nil, 0, err
	}
	return roles, totalItems, nil
}

func (r *roleRepository) GetRoleByIds(ids []uint) ([]*models.Role, error) {
	var roles []*models.Role

	if len(ids) == 0 {
		return []*models.Role{}, nil
	}

	if err := database.DB.Where("id IN ?", ids).Find(&roles).Error; err != nil {
		return nil, err
	}

	return roles, nil
}

func (r *roleRepository) GetRole(roleId uint) (*models.Role, error) {
	var role models.Role
	if err := database.DB.First(&role, roleId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("role not found")
		}
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) UpdateRole(role *models.Role) error {
	return database.DB.Save(role).Error
}

func (r *roleRepository) DeleteRole(role *models.Role) error {
	return database.DB.Delete(role).Error
}

func (r *roleRepository) GetRoleMenuPermision(roleId uint, param *dto.ParamQueryMenuPermissionDTO) (*models.Role, error) {
	var role models.Role

	if err := database.DB.First(&role, roleId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("role not found")
		}
		return nil, err
	}

	// get parent menu
	query := database.DB.Table("role_menus").Joins("JOIN menus ON role_menus.menu_id = menus.id").Where("role_id = ?", roleId)

	if param.CanView {
		query = query.Where("can_view = ?", true)
	}

	if param.MenuActive {
		query = query.Where("menus.active = ?", true)
	}

	if err := query.Preload("Menu").Find(&role.RoleMenus).Error; err != nil {
		return nil, err
	}

	return &role, nil
}
