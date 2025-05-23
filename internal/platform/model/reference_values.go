package model

type ReferenceValue struct {
	ID     uint          `gorm:"primaryKey"`
	TypeID uint          `gorm:"not null"`
	Type   ReferenceType `gorm:"foreignKey:TypeID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Code   string        `gorm:"size:30;not null"` // Valor único: 'checking', 'USD'
	Name   string        `gorm:"size:100;not null"`
}
