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
		{TypeID: rt.ID, Code: "Bs", Name: "Bolivares"},
	}

	return batchInsert(db, values, 100)
}

func batchInsert(db *gorm.DB, values []model.ReferenceValue, _ int) error {
	return db.Transaction(func(tx *gorm.DB) error {
		for _, value := range values {
			// Verificar si el valor ya existe usando TypeID y Code como identificador único
			result := tx.FirstOrCreate(&value, model.ReferenceValue{
				TypeID: value.TypeID,
				Code:   value.Code,
			})
			if result.Error != nil {
				return result.Error
			}
		}
		return nil
	})
}
