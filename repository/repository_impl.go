package repository

import (
	"errors"

	"gorm.io/gorm"
)

type RepositoryImpl[T any] struct {
	Db *gorm.DB
}

func NewRepositoryImpl[T any](Db *gorm.DB) Repository[T] {
	return &RepositoryImpl[T]{Db: Db}
}

func (r *RepositoryImpl[T]) Create(entity *T) error {
	return r.Db.Create(&entity).Error
}

func (r *RepositoryImpl[T]) Update(entity *T) error {
	return r.Db.Save(entity).Error
}

func (r *RepositoryImpl[T]) Delete(id uint) error {
	var entity T
	result := r.Db.Delete(&entity, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("record not found")
	}
	return nil
}

func (r *RepositoryImpl[T]) Exists(condition map[string]interface{}) (bool, error) {
	var count int64
	result := r.Db.Model(new(T)).Where(condition).Count(&count)

	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}

func (r *RepositoryImpl[T]) FindOne(condition map[string]interface{}, relations ...string) (*T, error) {
	var entity T
	query := r.Db

	for _, relation := range relations {
		query = query.Preload(relation)
	}

	result := query.Where(condition).First(&entity)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &entity, nil
}

func (r *RepositoryImpl[T]) FindAll(filters map[string]interface{}, relations ...string) ([]T, error) {
	var entites []T
	query := r.Db

	for key, value := range filters {
		query = query.Where(key+" = ?", value)
	}

	for _, relation := range relations {
		query = query.Preload(relation)
	}

	result := query.Find(&entites)

	if result.Error != nil {
		return nil, result.Error
	}
	return entites, nil
}
