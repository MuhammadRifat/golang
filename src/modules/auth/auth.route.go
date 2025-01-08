package auth

import (
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", AuthController.LoginHandler)
		authGroup.POST("/register", AuthController.RegisterHandler)
	}
}
