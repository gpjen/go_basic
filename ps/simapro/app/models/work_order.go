package models

import "time"

type WorkOrder struct {
	ID          uint   `gorm:"primaryKey"`
	OrderNumber string `gorm:"size:100;not null"`
	CreatedBy   string `gorm:"not null"`
	UpdatedBy   string
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
