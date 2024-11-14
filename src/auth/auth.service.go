package auth

import (
	"errors"
	"url-shortner/util"

	"golang.org/x/crypto/bcrypt"
)

type AuthServiceStruct struct{}

var AuthService = AuthServiceStruct{}

func (s *AuthServiceStruct) Login(request LoginRequest) (string, error) {
	var user User
	// Access the global DB instance directly
	err := util.DB.Where("email = ?", request.Email).First(&user).Error
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// Compare the provided password with the hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return "", errors.New("invalid email or password")
	}

	token, err := util.GenerateJWT(request.Email)
	if err != nil {
		return "", errors.New("error generating jwt")
	}

	return token, nil
}

func (s *AuthServiceStruct) Register(request RegisterRequest) (User, error) {
	// Hash the password before saving it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}

	user := User{
		Name:     request.Name,
		Email:    request.Email,
		Password: string(hashedPassword),
	}

	// Access the global DB instance directly
	if err := util.DB.Create(&user).Error; err != nil {
		return User{}, err
	}

	return user, nil
}
