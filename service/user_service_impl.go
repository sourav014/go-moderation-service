package service

import (
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/sourav014/go-moderation-service/dto"
	"github.com/sourav014/go-moderation-service/jwttoken"
	"github.com/sourav014/go-moderation-service/models"
	"github.com/sourav014/go-moderation-service/repository"
	"github.com/sourav014/go-moderation-service/utils"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
	JWTMaker       *jwttoken.JWTMaker
}

func NewUserServiceImpl(userRepository repository.UserRepository, validate *validator.Validate, jwtMaker *jwttoken.JWTMaker) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
		JWTMaker:       jwtMaker,
	}
}

func (u *UserServiceImpl) SignupUser(signupUserRequest dto.SignupUserRequest) (dto.SignupUserResponse, error) {
	var signupUserResponse dto.SignupUserResponse
	err := u.Validate.Struct(signupUserRequest)
	if err != nil {
		return signupUserResponse, err
	}

	user, err := u.UserRepository.FindByEmail(signupUserRequest.Email)
	if err != nil {
		return signupUserResponse, err
	}

	if user != nil {
		return signupUserResponse, errors.New("user already exists")
	}

	hashedPassword, err := utils.GenerateHashString(signupUserRequest.Password)
	if err != nil {
		return signupUserResponse, err
	}

	newUser := models.User{
		Name:     signupUserRequest.Name,
		Email:    signupUserRequest.Email,
		Password: hashedPassword,
	}
	u.UserRepository.Create(&newUser)

	signupUserResponse = dto.SignupUserResponse{
		Name:  newUser.Name,
		Email: newUser.Email,
	}

	return signupUserResponse, nil
}

func (u *UserServiceImpl) LoginUser(loginUserRequest dto.LoginUserRequest) (dto.LoginUserResponse, error) {
	var loginUserResponse dto.LoginUserResponse
	err := u.Validate.Struct(loginUserRequest)
	if err != nil {
		return loginUserResponse, err
	}

	user, err := u.UserRepository.FindByEmail(loginUserRequest.Email)
	if err != nil {
		return loginUserResponse, err
	}

	err = utils.CompareHashString(user.Password, loginUserRequest.Password)
	if err != nil {
		return loginUserResponse, errors.New("invalid credentials")
	}

	accessToken, _, err := u.JWTMaker.CreateToken(user.ID, time.Minute*15)
	if err != nil {
		return loginUserResponse, errors.New("failed to generate token")
	}

	refreshToken, _, err := u.JWTMaker.CreateToken(user.ID, time.Hour*24)
	if err != nil {
		return loginUserResponse, errors.New("failed to generate token")
	}

	loginUserResponse = dto.LoginUserResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User: dto.SignupUserResponse{
			Name:  user.Name,
			Email: user.Email,
		},
	}

	return loginUserResponse, nil
}
