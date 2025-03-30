package main

import (
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/sourav014/go-moderation-service/config"
	"github.com/sourav014/go-moderation-service/controller"
	"github.com/sourav014/go-moderation-service/db"
	"github.com/sourav014/go-moderation-service/jwttoken"
	"github.com/sourav014/go-moderation-service/middleware"
	"github.com/sourav014/go-moderation-service/models"
	"github.com/sourav014/go-moderation-service/repository"
	"github.com/sourav014/go-moderation-service/router"
	"github.com/sourav014/go-moderation-service/service"
)

func main() {
	config := config.NewConfig()

	db, err := db.NewDatabase(*config)
	if err != nil {
		log.Fatalf("error while initilizaing db %v", err)
	}

	db.Db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})

	userRepository := repository.NewUserRepository(db.GetDB())
	postRepository := repository.NewPostRepository(db.GetDB())
	commentRepository := repository.NewCommentRepository(db.GetDB())

	validate := validator.New()

	jwtMaker := jwttoken.NewJWTMaker(os.Getenv("SECRET_KEY"))
	userService := service.NewUserServiceImpl(*userRepository, validate, jwtMaker)
	userController := controller.NewUserController(userService)

	postService := service.NewPostServiceImpl(*postRepository)
	postController := controller.NewPostController(postService)

	commentService := service.NewCommentServiceImpl(*commentRepository)
	commentController := controller.NewCommentController(commentService)

	authMiddleware := middleware.NewAuthMiddlewareImpl(*userRepository, jwtMaker)

	router := router.SetupRouter(*userController, *postController, *commentController, authMiddleware)

	router.Run()
}
