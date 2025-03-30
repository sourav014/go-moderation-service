package repository

import (
	"github.com/sourav014/go-moderation-service/models"
	"gorm.io/gorm"
)

type PostRepository struct {
	repository Repository[models.Post]
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{
		repository: NewRepositoryImpl[models.Post](db),
	}
}

func (p *PostRepository) Create(post *models.Post) error {
	return p.repository.Create(post)
}

func (p *PostRepository) Update(post *models.Post) error {
	return p.repository.Update(post)
}

func (p *PostRepository) Delete(id uint) error {
	return p.repository.Delete(id)
}

func (p *PostRepository) FindById(id uint) (*models.Post, error) {
	condition := map[string]interface{}{"id": id}
	post, err := p.repository.FindOne(condition, "User", "Comments")

	if err != nil {
		return nil, err
	}

	return post, nil
}
