package model

import (
	"system/pkg/helpers"
	"time"
)

type User struct {
	helpers.UUID
	Name      string `gorm:"size:255;not null" json:"user_name"`
	Email     string `gorm:"unique;size:255;not null;index" json:"email"`
	Password  string `gorm:"size:255;not null" json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
