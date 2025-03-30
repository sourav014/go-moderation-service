package service

import (
	"errors"
	"time"

	"github.com/sourav014/go-moderation-service/dto"
	"github.com/sourav014/go-moderation-service/models"
	"github.com/sourav014/go-moderation-service/repository"
)

type PostServiceImpl struct {
	PostRepository repository.PostRepository
}

func NewPostServiceImpl(postRepository repository.PostRepository) PostService {
	return &PostServiceImpl{
		PostRepository: postRepository,
	}
}

func mapCommentsToResponse(comments []models.Comment) []dto.GetCommentResponse {
	var commentResponses []dto.GetCommentResponse
	for _, comment := range comments {
		commentResponses = append(commentResponses, dto.GetCommentResponse{
			ID:                   comment.ID,
			Content:              comment.Content,
			UserID:               comment.UserID,
			UserName:             comment.User.Name,
			PostID:               comment.PostID,
			FlaggedForModeration: comment.FlaggedForModeration,
			ModerationStatus:     comment.ModerationStatus,
			CreatedAt:            comment.CreatedAt,
			UpdatedAt:            comment.UpdatedAt,
		})
	}
	return commentResponses
}

func (p *PostServiceImpl) GetPost(id uint, user *models.User) (dto.GetPostResponse, error) {
	post, err := p.PostRepository.FindById(id)
	if err != nil {
		return dto.GetPostResponse{}, err
	}

	if post == nil {
		return dto.GetPostResponse{}, errors.New("post not found")
	}

	return dto.GetPostResponse{
		ID:        post.ID,
		Content:   post.Content,
		UserID:    post.UserID,
		UserName:  post.User.Name,
		Comments:  mapCommentsToResponse(post.Comments),
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}, nil
}

func (p *PostServiceImpl) CreatePost(createPostRequest dto.CreatePostRequest, user *models.User) (dto.CreatePostResponse, error) {
	post := models.Post{
		Content: createPostRequest.Content,
		UserID:  user.ID,
	}

	if err := p.PostRepository.Create(&post); err != nil {
		return dto.CreatePostResponse{}, err
	}

	response := dto.CreatePostResponse{
		ID:        post.ID,
		Content:   post.Content,
		UserID:    post.UserID,
		CreatedAt: post.CreatedAt,
	}

	return response, nil
}

func (p *PostServiceImpl) UpdatePost(updatePostRequest dto.UpdatePostRequest, postId uint, user *models.User) (dto.UpdatePostResponse, error) {
	post, err := p.PostRepository.FindById(postId)
	if err != nil {
		return dto.UpdatePostResponse{}, errors.New("post not found")
	}

	if post == nil {
		return dto.UpdatePostResponse{}, errors.New("post not found")
	}

	if post.UserID != user.ID {
		return dto.UpdatePostResponse{}, errors.New("unauthorized to update this post")
	}

	post.Content = updatePostRequest.Content
	post.UpdatedAt = time.Now()

	if err := p.PostRepository.Update(post); err != nil {
		return dto.UpdatePostResponse{}, err
	}

	response := dto.UpdatePostResponse{
		ID:        post.ID,
		Content:   post.Content,
		UserID:    post.UserID,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}

	return response, nil
}

func (p *PostServiceImpl) DeletePost(id uint, user *models.User) error {
	post, err := p.PostRepository.FindById(id)
	if err != nil {
		return err
	}

	if post == nil {
		return errors.New("post not found")
	}

	if post.UserID != user.ID {
		return errors.New("unauthorized: you can only delete your own posts")
	}

	return p.PostRepository.Delete(id)
}
