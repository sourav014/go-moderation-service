package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sourav014/go-moderation-service/dto"
	helpers "github.com/sourav014/go-moderation-service/helper"
	"github.com/sourav014/go-moderation-service/models"
	"github.com/sourav014/go-moderation-service/service"
)

type PostController struct {
	postService service.PostService
}

func NewPostController(postService service.PostService) *PostController {
	return &PostController{
		postService: postService,
	}
}

func (p *PostController) GetPost(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	currentUser, ok := ctx.MustGet("currentUser").(*models.User)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	getPostResponse, err := p.postService.GetPost(uint(id), currentUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, getPostResponse)
}

func (p *PostController) CreatePost(ctx *gin.Context) {
	var createPostRequest dto.CreatePostRequest
	if err := ctx.ShouldBindJSON(&createPostRequest); err != nil {
		helpers.HandleValidationError(ctx, createPostRequest, err)
		return
	}

	currentUser, ok := ctx.MustGet("currentUser").(*models.User)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	createPostResponse, err := p.postService.CreatePost(createPostRequest, currentUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Post created successfully", "post": createPostResponse})
}

func (p *PostController) UpdatePost(ctx *gin.Context) {
	var updatePostRequest dto.UpdatePostRequest
	if err := ctx.ShouldBindJSON(&updatePostRequest); err != nil {
		helpers.HandleValidationError(ctx, updatePostRequest, err)
		return
	}

	idParam := ctx.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	currentUser, ok := ctx.MustGet("currentUser").(*models.User)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	updatePostResponse, err := p.postService.UpdatePost(updatePostRequest, uint(id), currentUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Post updated successfully", "post": updatePostResponse})
}

func (p *PostController) DeletePost(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	currentUser, ok := ctx.MustGet("currentUser").(*models.User)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	err = p.postService.DeletePost(uint(id), currentUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
