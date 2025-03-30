package repository

import (
	"github.com/sourav014/go-moderation-service/models"
	"gorm.io/gorm"
)

type CommentRepository struct {
	repository Repository[models.Comment]
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{
		repository: NewRepositoryImpl[models.Comment](db),
	}
}

func (c *CommentRepository) Create(comment *models.Comment) error {
	return c.repository.Create(comment)
}

func (c *CommentRepository) Update(comment *models.Comment) error {
	return c.repository.Update(comment)
}

func (c *CommentRepository) Delete(id uint) error {
	return c.repository.Delete(id)
}

func (c *CommentRepository) FindById(id uint) (*models.Comment, error) {
	condition := map[string]interface{}{"id": id}
	comment, err := c.repository.FindOne(condition, "User", "Post")

	if err != nil {
		return nil, err
	}

	return comment, nil
}
