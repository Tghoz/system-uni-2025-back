package model

import (
	"system/pkg/helpers"
)

type Transaction struct {
	helpers.UUID
	Type      string  `gorm:"size:30;not null" json:"type"`
	CreatedAt string  `gorm:"not null"  json:"created_at"`
	Amount    float64 `gorm:"not null;default:0"`
	Currency  string  `gorm:"not null" json:"currency"`
	From      string  `gorm:"not null" json:"from"`
	To        string  `gorm:"not null" json:"to"`
	Status    string  `gorm:"size:30;not null" json:"status"`
	Reference string  `gorm:"not null" json:"reference"`

	AccountID uint     `gorm:"not null"`             // Clave foránea
	Account   *Account `gorm:"foreignKey:AccountID"` // Relación con la cuenta

	// Relación uno a muchos con Planning
	Plans []Planning `gorm:"foreignKey:TransactionID;constraint:OnDelete:CASCADE" json:"plans,omitempty"`
}
