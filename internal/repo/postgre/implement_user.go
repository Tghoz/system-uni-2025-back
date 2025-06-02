// system/internal/repo/postgre/user_repository.go
package postgre

import (
	"context"
	model "system/internal/models"

	"gorm.io/gorm"
)

// UserRepository implementa la interfaz Auth_Repo para el modelo User
type UserRepository struct {
	*GenericRepository[model.User]
}

// NewUserRepository crea un nuevo repositorio para usuarios
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		GenericRepository: NewGenericRepository[model.User](db),
	}
}

// CreateUser implementa el método de Auth_Repo
func (r *UserRepository) CreateUser(ctx context.Context, user *model.User) error {
	return r.Create(ctx, user)
}

// GetAllUsers implementa el método de Auth_Repo
func (r *UserRepository) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	return r.GetAll(ctx)
}

// GetUserById implementa el método de Auth_Repo
func (r *UserRepository) GetUserById(ctx context.Context, userID string) (*model.User, error) {
	return r.GetByID(ctx, userID)
}

// GetUserByEmail implementa el método de Auth_Repo
func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	return r.GetByField(ctx, "email", email)
}

// UpdateUser implementa el método de Auth_Repo
func (r *UserRepository) UpdateUser(ctx context.Context, user *model.User) error {
	return r.Update(ctx, user)
}

// DeleteUser implementa el método de Auth_Repo
func (r *UserRepository) DeleteUser(ctx context.Context, userID string) error {
	return r.Delete(ctx, userID)
}
