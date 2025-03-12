package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code        string `gorm:"size:255;not null;unique"`
	Name        string `gorm:"size:100;not null"`
	Description string `gorm:"type:text"`
	CreatedBy   string `gorm:"not null"`
}
