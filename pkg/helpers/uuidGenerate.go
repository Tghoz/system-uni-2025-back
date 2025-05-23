package helpers

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UUID struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
}

func (b *UUID) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == uuid.Nil {
		b.ID = uuid.New()
	}
	return
}
