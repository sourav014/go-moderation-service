package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sourav014/go-moderation-service/controller"
	"github.com/sourav014/go-moderation-service/middleware"
)

func SetupRouter(userController controller.UserController, postController controller.PostController, commentController controller.CommentController, middleware middleware.AuthMiddleware) *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		userRoutes := v1.Group("/users")
		{
			userRoutes.POST("/signup", userController.SignupUser)
			userRoutes.POST("/login", userController.LoginUser)
		}

		authorized := v1.Group("/")
		authorized.Use(middleware.CheckUserAuthentication)
		{
			postRoutes := authorized.Group("/posts")
			{
				postRoutes.GET("/:id", postController.GetPost)
				postRoutes.POST("/", postController.CreatePost)
				postRoutes.PUT("/:id", postController.UpdatePost)
				postRoutes.DELETE("/:id", postController.DeletePost)
			}

			commentRoutes := authorized.Group("/comments")
			{
				commentRoutes.GET("/:id", commentController.GetComment)
				commentRoutes.POST("/", commentController.CreateComment)
				commentRoutes.PUT("/:id", commentController.UpdateComment)
				commentRoutes.DELETE("/:id", commentController.DeleteComment)
			}
		}
	}

	return router
}
