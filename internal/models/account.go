package model

import (
	"fmt"

	"system/internal/platform/model"
	"system/pkg/helpers"
	"time"

	"gorm.io/gorm"
)

type Account struct {
	helpers.UUID
	AccountType  string  `gorm:"size:30;not null" json:"account"`
	CurrencyType string  `gorm:"size:30;not null" json:"currency"`
	Amount       float64 `gorm:"not null;default:0" json:"amount"`
	CreatedAt    time.Time

	// Relación con el usuario
	UserID string `gorm:"type:uuid;not null" json:"user_id"`
	User   *User  `gorm:"foreignKey:UserID"`

	// Relaciones
	Transactions []*Transaction `gorm:"foreignKey:AccountID"`
	Saving       []*Saving      `gorm:"foreignKey:AccountID"` // Relación con el ahorro

}

func (a *Account) ValidateAccountType(db *gorm.DB) error {
	var exists bool
	db.Model(&model.ReferenceValue{}).
		Joins("JOIN reference_types ON reference_values.type_id = reference_types.id").
		Where("reference_types.type = ? AND reference_values.code = ?", "account_type", a.AccountType).
		Select("count(*) > 0").
		Find(&exists)

	if !exists {
		return fmt.Errorf("invalid account type: %s", a.AccountType)
	}
	return nil
}

// Hook de GORM para validar antes de crear/actualizar
func (a *Account) BeforeSave(tx *gorm.DB) error {
	return a.ValidateAccountType(tx)
}
