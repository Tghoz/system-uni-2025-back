package migrate

import (
	"system/internal/platform/model"

	"gorm.io/gorm"
)

func seedReferenceTypes(db *gorm.DB) error {
	types := []model.ReferenceType{
		{Type: "account_type", Description: "Tipos de cuenta bancaria"},
		{Type: "currency", Description: "Divisas internacionales"},
	}

	for _, t := range types {
		result := db.FirstOrCreate(&t, model.ReferenceType{Type: t.Type})
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

func seedAccountTypes(db *gorm.DB) error {
	var rt model.ReferenceType
	if err := db.Where("type = ?", "account_type").First(&rt).Error; err != nil {
		return err
	}

	values := []model.ReferenceValue{
		{TypeID: rt.ID, Code: "checking", Name: "Cuenta Corriente"},
		{TypeID: rt.ID, Code: "savings", Name: "Cuenta de Ahorros"},
	}

	return batchInsert(db, values, 100)
}

func seedCurrencies(db *gorm.DB) error {
	var rt model.ReferenceType
	if err := db.Where("type = ?", "currency").First(&rt).Error; err != nil {
		return err
	}

	values := []model.ReferenceValue{
		{TypeID: rt.ID, Code: "USD", Name: "Dólar"},
		{TypeID: rt.ID, Code: "EUR", Name: "Euro"},
	}

	return batchInsert(db, values, 100)
}

func batchInsert(db *gorm.DB, values []model.ReferenceValue, batchSize int) error {
	return db.Transaction(func(tx *gorm.DB) error {
		for i := 0; i < len(values); i += batchSize {
			end := i + batchSize
			if end > len(values) {
				end = len(values)
			}

			chunk := values[i:end]
			if err := tx.Create(&chunk).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
