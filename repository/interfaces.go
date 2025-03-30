package repository

type Repository[T any] interface {
	Create(entity *T) error
	Update(entity *T) error
	Delete(id uint) error
	Exists(condition map[string]interface{}) (bool, error)
	FindOne(condition map[string]interface{}, relations ...string) (*T, error)
	FindAll(condition map[string]interface{}, relations ...string) ([]T, error)
}
