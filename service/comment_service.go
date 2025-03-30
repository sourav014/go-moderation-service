package service

import (
	"github.com/sourav014/go-moderation-service/dto"
	"github.com/sourav014/go-moderation-service/models"
)

type CommentService interface {
	GetComment(id uint, user *models.User) (dto.GetCommentResponse, error)
	CreateComment(createCommentRequest dto.CreateCommentRequest, user *models.User) (dto.CreateCommentResponse, error)
	UpdateComment(updateCommentRequest dto.UpdateCommentRequest, commendId uint, user *models.User) (dto.UpdateCommentResponse, error)
	DeleteComment(id uint, user *models.User) error
}
