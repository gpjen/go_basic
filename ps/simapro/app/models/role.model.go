package models

import (
	"time"

	"github.com/gpjen/simapro/app/dto"
)

type Role struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"size:100;not null;unique"`
	Description string    `gorm:"size:255;not null"`
	Active      bool      `gorm:"default:true"`
	CreatedBy   string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`

	RoleMenus []RoleMenu `gorm:"foreignKey:RoleID"`
}

func (m *Role) ToRoleDTO() *dto.RoleDTO {
	return &dto.RoleDTO{
		ID:          m.ID,
		Name:        m.Name,
		Description: m.Description,
	}
}

func (m *Role) ToRoleDetailDTO() *dto.RoleDetailDTO {
	res := &dto.RoleDetailDTO{
		ID:          m.ID,
		Name:        m.Name,
		Description: m.Description,
		CreatedBy:   m.CreatedBy,
		CreatedAt:   m.CreatedAt.Format("02-01-2006 15:04:05"),
		UpdatedAt:   m.UpdatedAt.Format("02-01-2006 15:04:05"),
	}

	return res
}

func (m *Role) ToRoleMenuPermision() *dto.MenuPermissionDTO {

	res := &dto.MenuPermissionDTO{
		RoleID:      m.ID,
		Name:        m.Name,
		Description: m.Description,
		Menus:       []dto.MenuDTO{},
	}

	for _, roleMenu := range m.RoleMenus {
		menu := dto.MenuDTO{
			MenuID:     roleMenu.MenuID,
			Name:       roleMenu.Menu.Name,
			Route:      roleMenu.Menu.Route,
			Icon:       roleMenu.Menu.Icon,
			CanView:    roleMenu.CanView,
			CanShow:    roleMenu.CanShow,
			CanCreate:  roleMenu.CanCreate,
			CanUpdate:  roleMenu.CanUpdate,
			CanDelete:  roleMenu.CanDelete,
			Parent:     roleMenu.Menu.IdParent,
			Order:      roleMenu.Menu.Order,
			MenuActive: roleMenu.Menu.Active,
		}

		res.Menus = append(res.Menus, menu)
	}

	return res
}
