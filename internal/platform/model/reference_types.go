package model

type ReferenceType struct {
	ID          uint   `gorm:"primaryKey"`
	Type        string `gorm:"uniqueIndex;size:50;not null"` // Ej: 'account_type', 'currency'
	Description string
}
