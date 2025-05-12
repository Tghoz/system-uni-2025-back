package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name      string    `gorm:"size:255"`
	Email     string    `gorm:"unique;size:255"`
	Password  string    `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Hook de GORM: Se ejecuta ANTES de crear el registro en la DB.
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return
}
