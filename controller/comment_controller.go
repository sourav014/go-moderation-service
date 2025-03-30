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

type CommentController struct {
	commentService service.CommentService
}

func NewCommentController(commentService service.CommentService) *CommentController {
	return &CommentController{
		commentService: commentService,
	}
}

func (c *CommentController) GetComment(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}

	currentUser, ok := ctx.MustGet("currentUser").(*models.User)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	getCommentResponse, err := c.commentService.GetComment(uint(id), currentUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"comment": getCommentResponse})
}

func (c *CommentController) CreateComment(ctx *gin.Context) {
	var createCommentRequest dto.CreateCommentRequest
	if err := ctx.ShouldBindJSON(&createCommentRequest); err != nil {
		helpers.HandleValidationError(ctx, createCommentRequest, err)
		return
	}

	currentUser, ok := ctx.MustGet("currentUser").(*models.User)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	createCommentResponse, err := c.commentService.CreateComment(createCommentRequest, currentUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Comment created successfully", "comment": createCommentResponse})
}

func (c *CommentController) UpdateComment(ctx *gin.Context) {
	var updateCommentRequest dto.UpdateCommentRequest
	if err := ctx.ShouldBindJSON(&updateCommentRequest); err != nil {
		helpers.HandleValidationError(ctx, updateCommentRequest, err)
		return
	}

	currentUser, ok := ctx.MustGet("currentUser").(*models.User)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	idParam := ctx.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comments ID"})
		return
	}

	updateCommentResponse, err := c.commentService.UpdateComment(updateCommentRequest, uint(id), currentUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Comment updated successfully", "comment": updateCommentResponse})
}

func (c *CommentController) DeleteComment(ctx *gin.Context) {
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

	err = c.commentService.DeleteComment(uint(id), currentUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}
