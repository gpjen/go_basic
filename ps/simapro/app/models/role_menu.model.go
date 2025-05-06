package models

import (
	"time"
)

type RoleMenu struct {
	RoleID uint `gorm:"primaryKey;not null;index"`
	MenuID uint `gorm:"primaryKey;not null;index"`

	CanView   bool `gorm:"default:false;not null"`
	CanShow   bool `gorm:"default:false;not null"`
	CanCreate bool `gorm:"default:false;not null"`
	CanUpdate bool `gorm:"default:false;not null"`
	CanDelete bool `gorm:"default:false;not null"`

	CreatedBy string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	Role Role `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Menu Menu `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// func (m *RoleMenu) ToRoleMenuDTO() *dto.MenuPermissionDTO {

// 	res := &dto.MenuPermissionDTO{}

// 	return res
// }
