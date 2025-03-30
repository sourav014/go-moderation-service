package service

import (
	"errors"
	"log"

	"github.com/sourav014/go-moderation-service/constants"
	"github.com/sourav014/go-moderation-service/dto"
	"github.com/sourav014/go-moderation-service/models"
	"github.com/sourav014/go-moderation-service/repository"
	"github.com/sourav014/go-moderation-service/service/moderation"
	"github.com/sourav014/go-moderation-service/service/notification"
)

type CommentServiceImpl struct {
	CommentRepository repository.CommentRepository
}

func NewCommentServiceImpl(commentRepository repository.CommentRepository) CommentService {
	return &CommentServiceImpl{
		CommentRepository: commentRepository,
	}
}

func (c *CommentServiceImpl) GetComment(id uint, user *models.User) (dto.GetCommentResponse, error) {
	comment, err := c.CommentRepository.FindById(id)
	if err != nil {
		return dto.GetCommentResponse{}, err
	}

	if comment == nil {
		return dto.GetCommentResponse{}, errors.New("comment not found")
	}

	return dto.GetCommentResponse{
		ID:                   comment.ID,
		Content:              comment.Content,
		UserID:               comment.UserID,
		UserName:             comment.User.Name,
		PostID:               comment.PostID,
		FlaggedForModeration: comment.FlaggedForModeration,
		ModerationStatus:     comment.ModerationStatus,
		CreatedAt:            comment.CreatedAt,
		UpdatedAt:            comment.UpdatedAt,
	}, nil
}

func (c *CommentServiceImpl) CreateComment(createCommentRequest dto.CreateCommentRequest, user *models.User) (dto.CreateCommentResponse, error) {
	comment := models.Comment{
		Content: createCommentRequest.Content,
		UserID:  user.ID,
		PostID:  createCommentRequest.PostID,
	}

	moderationResults, err := moderation.GetContentModerationDetails(comment.Content)
	if err != nil {
		return dto.CreateCommentResponse{}, err
	}

	comment.FlaggedForModeration = false
	comment.ModerationStatus = "approved"

	for category, confidence := range moderationResults {
		if threshold, exists := constants.ModerationThresholds[category]; exists && confidence >= threshold {
			comment.FlaggedForModeration = true
			comment.ModerationStatus = "flagged"
			break
		}
	}

	err = c.CommentRepository.Create(&comment)
	if err != nil {
		return dto.CreateCommentResponse{}, err
	}

	if comment.FlaggedForModeration {
		go func() {
			err := notification.SendEmail(user.Email, "Comment Moderation Alert", "Your comment has been flagged for review due to policy violations.")
			if err != nil {
				log.Printf("Failed to send moderation email to %s: %v", user.Email, err)
			}
		}()
	}

	return dto.CreateCommentResponse{
		ID:        comment.ID,
		Content:   comment.Content,
		PostID:    comment.PostID,
		UserID:    comment.UserID,
		CreatedAt: comment.CreatedAt,
	}, nil
}

func (c *CommentServiceImpl) UpdateComment(updateCommentRequest dto.UpdateCommentRequest, commentId uint, user *models.User) (dto.UpdateCommentResponse, error) {
	comment, err := c.CommentRepository.FindById(commentId)
	if err != nil {
		return dto.UpdateCommentResponse{}, err
	}

	if comment == nil {
		return dto.UpdateCommentResponse{}, errors.New("comment not found")
	}

	if comment.UserID != user.ID {
		return dto.UpdateCommentResponse{}, errors.New("unauthorized: you can only update your own comments")
	}

	comment.Content = updateCommentRequest.Content

	moderationResults, err := moderation.GetContentModerationDetails(comment.Content)
	if err != nil {
		return dto.UpdateCommentResponse{}, err
	}

	comment.FlaggedForModeration = false
	comment.ModerationStatus = "approved"

	for category, confidence := range moderationResults {
		if threshold, exists := constants.ModerationThresholds[category]; exists && confidence >= threshold {
			comment.FlaggedForModeration = true
			comment.ModerationStatus = "flagged"
			break
		}
	}

	err = c.CommentRepository.Update(comment)
	if err != nil {
		return dto.UpdateCommentResponse{}, err
	}

	if comment.FlaggedForModeration {
		go func() {
			err := notification.SendEmail(user.Email, "Comment Moderation Alert", "Your updated comment has been flagged for review due to policy violations.")
			if err != nil {
				log.Printf("Failed to send moderation email to %s: %v", user.Email, err)
			}
		}()
	}

	return dto.UpdateCommentResponse{
		ID:        comment.ID,
		Content:   comment.Content,
		PostID:    comment.PostID,
		UserID:    comment.UserID,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
	}, nil
}

func (c *CommentServiceImpl) DeleteComment(id uint, user *models.User) error {
	comment, err := c.CommentRepository.FindById(id)
	if err != nil {
		return err
	}

	if comment == nil {
		return errors.New("comment not found")
	}

	if comment.UserID != user.ID {
		return errors.New("unauthorized: you can only delete your own comments")
	}

	return c.CommentRepository.Delete(id)
}
