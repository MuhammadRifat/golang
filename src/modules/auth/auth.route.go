package auth

import (
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", LoginHandler)
		authGroup.POST("/register", RegisterHandler)
		// Add more routes as needed
	}
}
