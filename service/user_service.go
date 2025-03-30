package service

import (
	"github.com/sourav014/go-moderation-service/dto"
)

type UserService interface {
	SignupUser(signupUserRequest dto.SignupUserRequest) (dto.SignupUserResponse, error)
	LoginUser(loginUserRequest dto.LoginUserRequest) (dto.LoginUserResponse, error)
}
