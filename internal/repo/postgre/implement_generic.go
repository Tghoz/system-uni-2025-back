// system/internal/repo/postgre/generic_ repository.go
package postgre

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

// GenericRepository maneja operaciones CRUD básicas para cualquier modelo
type GenericRepository[T any] struct {
	db *gorm.DB
}

func NewGenericRepository[T any](db *gorm.DB) *GenericRepository[T] {
	return &GenericRepository[T]{db: db}
}

// Create crea una nueva entidad
func (r *GenericRepository[T]) Create(ctx context.Context, entity *T) error {
	return r.db.WithContext(ctx).Create(entity).Error
}

// GetAll obtiene todas las entidades
func (r *GenericRepository[T]) GetAll(ctx context.Context) ([]*T, error) {
	var entities []*T
	result := r.db.WithContext(ctx).Find(&entities)
	return entities, result.Error
}

// GetByID obtiene una entidad por su ID
func (r *GenericRepository[T]) GetByID(ctx context.Context, id string) (*T, error) {
	var entity T
	result := r.db.WithContext(ctx).Where("id = ?", id).First(&entity)
	if result.Error != nil {
		return nil, result.Error // ¡Cambio importante aquí!
	}
	return &entity, result.Error
}

// GetByField obtiene una entidad por un campo específico
func (r *GenericRepository[T]) GetByField(ctx context.Context, field string, value interface{}) (*T, error) {
	var entity T
	query := field + " = ?"
	result := r.db.WithContext(ctx).Where(query, value).First(&entity)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &entity, result.Error
}

// Update actualiza una entidad
func (r *GenericRepository[T]) Update(ctx context.Context, entity *T) error {
	return r.db.WithContext(ctx).Save(entity).Error
}

// Delete elimina una entidad por su ID
func (r *GenericRepository[T]) Delete(ctx context.Context, id string) error {
	var entity T
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&entity).Error
}

