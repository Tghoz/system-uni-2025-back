package repo

import (
	"context"
	model "system/internal/models"
)

type Planning_inteface interface {
	Create(ctx context.Context, account *model.Planning) error
	GetAll(ctx context.Context) ([]*model.Planning, error)
	GetByID(ctx context.Context, id string) (*model.Planning, error)
	Update(ctx context.Context, account *model.Planning) error
	Delete(ctx context.Context, id string) error
	// Agrega otros métodos según necesites
}
