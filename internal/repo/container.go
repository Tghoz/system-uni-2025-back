// system/internal/repo/container.go
package repo

import (
	models "system/internal/models"
	"system/internal/repo/postgre"

	account "system/internal/accounts/interface"
	planning "system/internal/planning/interface"
	transaction "system/internal/transaction/interface"

	"gorm.io/gorm"
)

type RepositoryContainer struct {
	User        *postgre.UserRepository // Usa el repositorio específico para User
	Account     account.Account_inteface
	Transaction transaction.Transaction_interface
	Planning    planning.Planning_inteface
}

func NewRepositoryContainer(db *gorm.DB) *RepositoryContainer {
	return &RepositoryContainer{
		User:        postgre.NewUserRepository(db),
		Account:     postgre.NewGenericRepository[models.Account](db),
		Transaction: postgre.NewGenericRepository[models.Transaction](db),
		Planning:    postgre.NewGenericRepository[models.Planning](db),
	}
}
