package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sourav014/go-moderation-service/jwttoken"
	"github.com/sourav014/go-moderation-service/repository"
)

type AuthMiddlewareImpl struct {
	UserReposity repository.UserRepository
	JWTMaker     *jwttoken.JWTMaker
}

func NewAuthMiddlewareImpl(userRepository repository.UserRepository, jwtMaker *jwttoken.JWTMaker) AuthMiddleware {
	return &AuthMiddlewareImpl{
		UserReposity: userRepository,
		JWTMaker:     jwtMaker,
	}
}

func (auth *AuthMiddlewareImpl) CheckUserAuthentication(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")

	if authHeader == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		return
	}

	authToken := strings.Split(authHeader, " ")
	if len(authToken) != 2 || authToken[0] != "Bearer" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
		return
	}

	tokenString := authToken[1]

	userClaims, err := auth.JWTMaker.VerifyToken(tokenString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return
	}

	if time.Now().Unix() > userClaims.ExpiresAt.Unix() {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
		return
	}

	user, err := auth.UserReposity.FindById(userClaims.ID)
	if err != nil || user.ID == 0 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Set("currentUser", user)

	ctx.Next()
}
