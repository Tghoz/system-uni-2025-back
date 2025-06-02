package model



type Saving struct {
	ID          string `json:"id"`
	Name		string `json:"name"`
	Amount      int64  `json:"amount"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
	DeletedAt   int64  `json:"deleted_at"`

	AccountID uint     `gorm:"not null"`             // Clave foránea
	Account   *Account `gorm:"foreignKey:AccountID"` // Relación con la cuenta
}