package repository

import (
	"github.com/sourav014/go-moderation-service/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	repository Repository[models.User]
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		repository: NewRepositoryImpl[models.User](db),
	}
}

func (u *UserRepository) Create(user *models.User) error {
	return u.repository.Create(user)
}

func (u *UserRepository) FindByEmail(email string) (*models.User, error) {
	condition := map[string]interface{}{"email": email}

	user, err := u.repository.FindOne(condition)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) FindById(id uint) (*models.User, error) {
	condition := map[string]interface{}{"id": id}

	user, err := u.repository.FindOne(condition)
	if err != nil {
		return nil, err
	}

	return user, nil
}
