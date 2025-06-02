// system/internal/repo/container.go
package repo

import (
	models "system/internal/models"
	"system/internal/repo/postgre"

	account "system/internal/accounts/interface"

	"gorm.io/gorm"
)

type RepositoryContainer struct {
	User    *postgre.UserRepository // Usa el repositorio específico para User
	Account account.Account_inteface
}

func NewRepositoryContainer(db *gorm.DB) *RepositoryContainer {
	return &RepositoryContainer{
		User:    postgre.NewUserRepository(db),
		Account: postgre.NewGenericRepository[models.Account](db),
	}
}
