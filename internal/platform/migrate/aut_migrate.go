package migrate

import (
	"fmt"
	models "system/internal/models"
	"system/internal/platform/model"

	"gorm.io/gorm"
)

type Migrator struct {
	db *gorm.DB
}

func NewMigrator(db *gorm.DB) *Migrator {
	return &Migrator{db: db}
}

func (m *Migrator) Run() error {
	// Orden estratégico de migración
	migrationSequence := []func() error{
		m.migrateReferenceTables,
		m.migrateCoreTables,
		m.seedEssentialData,
	}
	for _, step := range migrationSequence {
		if err := step(); err != nil {
			return fmt.Errorf("migration failed: %v", err)
		}
	}
	return nil
}

func (m *Migrator) migrateReferenceTables() error {
	return m.db.AutoMigrate(
		&model.ReferenceType{},
		&model.ReferenceValue{},
	)
}

func (m *Migrator) seedEssentialData() error {
	essentialData := []struct {
		name     string
		seedFunc func(*gorm.DB) error
	}{
		{"reference_types", seedReferenceTypes},
		{"account_types", seedAccountTypes},
		{"currencies", seedCurrencies},
	}
	for _, data := range essentialData {
		if err := data.seedFunc(m.db); err != nil {
			return fmt.Errorf("seeding %s failed: %v", data.name, err)
		}
	}
	return nil
}

func (m *Migrator) migrateCoreTables() error {
	return m.db.AutoMigrate(
		&models.Account{},
		&models.User{},
		&models.Transaction{},
		&models.Saving{},
		&models.Planning{},
		// Agregar otros modelos principales aquí
	)
}
