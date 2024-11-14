package util

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("test-secret-key")

// CustomClaims will hold additional claims you want to add to the JWT
type CustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateJWT(username string) (string, error) {
	// Set the expiration time for the token
	expirationTime := time.Now().Add(24 * time.Hour) // 1 day expiration

	// Define the claims
	claims := &CustomClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "URL Shortner",
		},
	}

	// Create a token using the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
