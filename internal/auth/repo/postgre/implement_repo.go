package postgre

import (
	"context"
	"system/internal/models"
	"system/internal/auth/repo" // Importar la interfaz

	"gorm.io/gorm"
)

var _ repo.Auth_Repo = (*AuthPostgresRepo)(nil)

type AuthPostgresRepo struct {
	db *gorm.DB // Conexión a la base de datos (inyectada)
}

func NewUserRepository(db *gorm.DB) repo.Auth_Repo { // Tipo de retorno: interfaz
	return &AuthPostgresRepo{db: db}
}

func (r *AuthPostgresRepo) CreateUser(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *AuthPostgresRepo) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	result := r.db.WithContext(ctx).Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *AuthPostgresRepo) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	var users []*model.User

	
	result := r.db.WithContext(ctx).
		Order("created_at DESC"). 
		Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
