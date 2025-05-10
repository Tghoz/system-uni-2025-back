package models

import (
	"gorm.io/gorm"
)

// Defines todo table for database communications
type User struct {
	gorm.Model
	Name     string `gorm:"size:255"`
	Email    string `gorm:"unique;size:255"`
	Password string `gorm:"size:255"`
}
