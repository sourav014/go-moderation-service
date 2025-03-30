package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sourav014/go-moderation-service/dto"
	helpers "github.com/sourav014/go-moderation-service/helper"
	"github.com/sourav014/go-moderation-service/service"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (u *UserController) SignupUser(ctx *gin.Context) {
	var signupUserRequest dto.SignupUserRequest
	if err := ctx.ShouldBindJSON(&signupUserRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	signupUserResponse, err := u.userService.SignupUser(signupUserRequest)
	if err != nil {
		helpers.HandleValidationError(ctx, signupUserRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User signed up successfully", "user": signupUserResponse})
}

func (u *UserController) LoginUser(ctx *gin.Context) {
	var loginUserRequest dto.LoginUserRequest
	if err := ctx.ShouldBindJSON(&loginUserRequest); err != nil {
		helpers.HandleValidationError(ctx, loginUserRequest, err)
		return
	}

	loginUserResponse, err := u.userService.LoginUser(loginUserRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "data": loginUserResponse})
}
