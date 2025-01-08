package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthControllerStruct struct{}

var AuthController = AuthControllerStruct{}

func (c *AuthControllerStruct) LoginHandler(ctx *gin.Context) {
	var loginRequest LoginRequest
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the auth service to handle login
	token, err := AuthService.Login(loginRequest)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (c *AuthControllerStruct) RegisterHandler(ctx *gin.Context) {
	var registerRequest RegisterRequest
	if err := ctx.ShouldBindJSON(&registerRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the auth service to handle registration
	user, err := AuthService.Register(registerRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}
