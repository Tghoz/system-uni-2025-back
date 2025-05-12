package postgre

import (
	"context"
	"system/internal/auth/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB // Conexión a la base de datos (inyectada)
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	result := r.db.WithContext(ctx).Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

//! crrear mas funciones para el repo
