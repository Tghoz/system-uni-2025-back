// system/internal/repo/account_repository.go
package repo

import (
	"context"
	model "system/internal/models"
)

type Account_inteface interface {
	Create(ctx context.Context, account *model.Account) error
	GetByID(ctx context.Context, id string) (*model.Account, error)
	// Agrega otros métodos según necesites
}
