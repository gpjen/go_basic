package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `json:"name" gorm:"not null"`
	Email string `json:"email" gorm:"not null"`
}
