package services

import (
	"github.com/gpjen/simapro/app/dto"
	"github.com/gpjen/simapro/app/models"
	"github.com/gpjen/simapro/app/repositories"
)

type MenuService interface {
	GetMenuHierarchy(param *dto.ParamQueryMenuDTO) ([]*models.Menu, error)
	MenuUpdates(dto dto.UpdateMenusDto) error
}

type menuService struct {
	menuRepository repositories.MenuRepository
}

func NewMenuService(menuRepository repositories.MenuRepository) MenuService {
	return &menuService{
		menuRepository: menuRepository,
	}
}

func (m *menuService) GetMenuHierarchy(param *dto.ParamQueryMenuDTO) ([]*models.Menu, error) {
	return m.menuRepository.GetMenuHierarchy(param)
}

func (m *menuService) MenuUpdates(input dto.UpdateMenusDto) error {
	return m.menuRepository.Updates(input)
}
