package model

import "gorm.io/gorm"

type Planning struct {
	gorm.Model
	Name        string  `gorm:"size:100;not null" json:"name"`
	Description string  `gorm:"size:255" json:"description"`
	Amount      float64 `gorm:"not null;default:0" json:"amount"`
	Service     string  `gorm:"size:50;not null" json:"service"`
	Value       float64 `gorm:"not null;default:0" json:"value"`
	Month       string  `gorm:"not null" json:"date"`


	// Nueva clave foránea (tipo string para UUID)
    TransactionID string      `gorm:"size:36;not null;index" json:"transaction_id"` // index para mejor rendimiento
    Transaction   Transaction `gorm:"foreignKey:TransactionID" json:"-"`
}
