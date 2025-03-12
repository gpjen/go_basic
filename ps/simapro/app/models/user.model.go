package models

import (
	"time"

	"github.com/gpjen/simapro/app/dto"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"size:100;not null"`
	Email     string    `gorm:"uniqueIndex;size:100;not null"`
	Password  string    `gorm:"size:255;not null"`
	CreatedBy string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	Roles []Role `gorm:"many2many:user_roles"`
}

func (m *User) ToUserDTO() *dto.UserDTO {
	return &dto.UserDTO{
		ID:    m.ID,
		Name:  m.Name,
		Email: m.Email,
	}
}

func (m *User) ToUserDetailDTO() *dto.UserDetailDTO {

	res := &dto.UserDetailDTO{
		ID:        m.ID,
		Name:      m.Name,
		Email:     m.Email,
		CreatedBy: m.CreatedBy,
		CreatedAt: m.CreatedAt.Format("02-01-2006 15:04:05"),
		UpdatedAt: m.UpdatedAt.Format("02-01-2006 15:04:05"),
		Roles:     []dto.RoleDTO{},
	}

	for _, role := range m.Roles {
		roledto := role.ToRoleDTO()
		res.Roles = append(res.Roles, *roledto)
	}

	return res

}
