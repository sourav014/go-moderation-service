package service

import (
	"github.com/sourav014/go-moderation-service/dto"
	"github.com/sourav014/go-moderation-service/models"
)

type PostService interface {
	GetPost(id uint, user *models.User) (dto.GetPostResponse, error)
	CreatePost(createPostRequest dto.CreatePostRequest, user *models.User) (dto.CreatePostResponse, error)
	UpdatePost(updatePostRequest dto.UpdatePostRequest, postId uint, user *models.User) (dto.UpdatePostResponse, error)
	DeletePost(id uint, user *models.User) error
}
