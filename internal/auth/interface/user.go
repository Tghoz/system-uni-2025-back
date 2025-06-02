package interfaz

import (
	"context"
	model "system/internal/models"
)

type Auth_Repo interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetAllUsers(ctx context.Context) ([]*model.User, error)
	GetUserById(ctx context.Context, userID string) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, userID string) error
}
