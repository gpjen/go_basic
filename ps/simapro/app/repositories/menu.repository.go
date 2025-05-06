package repositories

import (
	"strings"

	"github.com/gpjen/simapro/app/dto"
	"github.com/gpjen/simapro/app/models"
	"github.com/gpjen/simapro/database"
	"github.com/gpjen/simapro/utils"
	"gorm.io/gorm"
)

type MenuRepository interface {
	GetMenuHierarchy(param *dto.ParamQueryMenuDTO) ([]*models.Menu, error)
	Updates(input dto.UpdateMenusDto) error

	buildMenuHierarchy(menu *models.Menu, param *dto.ParamQueryMenuDTO, currentDepth int, maxDepth int)
}

type menuRepository struct {
	db *gorm.DB
}

func NewMenuRepository() MenuRepository {
	return &menuRepository{db: database.DB}
}

func (r *menuRepository) GetMenuHierarchy(param *dto.ParamQueryMenuDTO) ([]*models.Menu, error) {

	var rootMenus []*models.Menu

	query := r.db

	if param.Active {
		query = query.Where("active = ?", true)
	}

	order := "ASC"
	if strings.ToLower(param.Order) == "desc" {
		order = "DESC"
	}

	err := query.Where("id_parent IS NULL").Order("\"active\" DESC, \"order\" " + order).Find(&rootMenus).Error
	if err != nil {
		return nil, err
	}

	for i := range rootMenus {
		r.buildMenuHierarchy(rootMenus[i], param, 1, utils.MaxHirarcy)
	}

	return rootMenus, nil
}

func (r *menuRepository) buildMenuHierarchy(menu *models.Menu, param *dto.ParamQueryMenuDTO, currentDepth int, maxDepth int) {

	if currentDepth >= maxDepth {
		return
	}

	query := r.db.Where("id_parent = ?", menu.ID)

	if param.Active {
		query = query.Where("active = ?", true)
	}

	order := "ASC"
	if strings.ToLower(param.Order) == "desc" {
		order = "DESC"
	}

	var children []models.Menu
	err := query.Order("\"active\" DESC, \"order\" " + order).Find(&children).Error
	if err != nil || len(children) == 0 {
		return
	}

	menu.Children = children

	for i := range menu.Children {
		r.buildMenuHierarchy(&menu.Children[i], param, currentDepth+1, maxDepth)
	}
}

func (r *menuRepository) Updates(input dto.UpdateMenusDto) error {
	tx := r.db.Begin()

	for _, data := range input.Data {
		updateFields := map[string]interface{}{
			"id_parent": data.IdParent,
			"icon":      data.Icon,
			"order":     data.Order,
			"active":    data.Active,
		}

		if err := tx.Model(&models.Menu{}).
			Where("id = ?", data.Id).
			Updates(updateFields).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
