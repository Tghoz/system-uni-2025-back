package repo

import (
	"context"
	model "system/internal/models"
)

type Transaction_interface interface {
	Create(ctx context.Context, account *model.Transaction) error
	GetAll(ctx context.Context) ([]*model.Transaction, error)
	GetByID(ctx context.Context, id string) (*model.Transaction, error)
}
