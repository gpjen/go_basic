package services

import (
	"github.com/gpjen/simapro/app/models"
	"github.com/gpjen/simapro/app/repositories"
)

type RoleMenuService interface {
	FindAllByRoleID(roleID uint) []models.RoleMenu
}

type roleMenuService struct {
	roleMenuRepository repositories.RoleMenuRepository
}

func NewRoleMenuService(roleMenuRepository repositories.RoleMenuRepository) RoleMenuService {
	return &roleMenuService{roleMenuRepository}
}

func (r *roleMenuService) FindAllByRoleID(roleID uint) []models.RoleMenu {
	return r.roleMenuRepository.FindAllByRoleID(roleID)
}
