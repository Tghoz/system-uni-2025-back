// system/internal/repo/account_repository.go
package repo

import (
	"context"
	model "system/internal/models"
)

type Account_inteface interface {
	Create(ctx context.Context, account *model.Account) error
	GetAll(ctx context.Context) ([]*model.Account, error)
	GetByID(ctx context.Context, id string) (*model.Account, error)
	Update(ctx context.Context, account *model.Account) error
	Delete(ctx context.Context, id string) error
	// Agrega otros métodos según necesites
}
