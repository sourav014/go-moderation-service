package dto

import "time"

type SignupUserRequest struct {
	Name     string `binding:"required,min=3,max=100" json:"name"`
	Email    string `binding:"required,email" json:"email"`
	Password string `binding:"required,min=8" json:"password"`
}

type SignupUserResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type LoginUserRequest struct {
	Email    string `binding:"required,email" json:"email"`
	Password string `binding:"required" json:"password"`
}

type LoginUserResponse struct {
	AccessToken  string             `json:"access_token"`
	RefreshToken string             `json:"refresh_token"`
	User         SignupUserResponse `json:"user"`
}

type CreatePostRequest struct {
	Content string `json:"content" binding:"required"`
}

type CreatePostResponse struct {
	ID        uint      `json:"id"`
	Content   string    `json:"content"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdatePostRequest struct {
	Content string `json:"content" binding:"required"`
}

type UpdatePostResponse struct {
	ID        uint      `json:"id"`
	Content   string    `json:"content"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateCommentRequest struct {
	PostID  uint   `json:"post_id" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type CreateCommentResponse struct {
	ID        uint      `json:"id"`
	Content   string    `json:"content"`
	UserID    uint      `json:"user_id"`
	PostID    uint      `json:"post_id"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateCommentRequest struct {
	Content string `json:"content" binding:"required"`
}

type UpdateCommentResponse struct {
	ID        uint      `json:"id"`
	Content   string    `json:"content"`
	UserID    uint      `json:"user_id"`
	PostID    uint      `json:"post_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetPostResponse struct {
	ID        uint                 `json:"id"`
	Content   string               `json:"content"`
	UserID    uint                 `json:"user_id"`
	UserName  string               `json:"user_name"`
	Comments  []GetCommentResponse `json:"comments"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
}

type GetCommentResponse struct {
	ID                   uint      `json:"id"`
	Content              string    `json:"content"`
	UserID               uint      `json:"user_id"`
	UserName             string    `json:"user_name"`
	PostID               uint      `json:"post_id"`
	FlaggedForModeration bool      `json:"flagged_for_moderation"`
	ModerationStatus     string    `json:"moderation_status"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}
