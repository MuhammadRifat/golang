package auth

import (
	"url-shortner/src/common"
	"url-shortner/src/util"

	"golang.org/x/crypto/bcrypt"
)

type AuthServiceStruct struct {
	common.ServiceStruct[User]
}

var AuthService = AuthServiceStruct{
	ServiceStruct: common.ServiceStruct[User]{},
}

func (s *AuthServiceStruct) Login(request LoginRequest) (string, error) {
	user, err := s.FindOneRecordByQuery(map[string]interface{}{"email": request.Email})
	if err != nil {
		return "", util.UnauthorizedErr("invalid email or password")
	}

	// Compare the provided password with the hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return "", util.UnauthorizedErr("invalid email or password")
	}

	token, err := util.GenerateJWT(int(user.ID))
	if err != nil {
		return "", util.InternalServerErr("error generating jwt")
	}

	return token, nil
}

func (s *AuthServiceStruct) Register(request RegisterRequest) (User, error) {
	queryMap := map[string]interface{}{"email": request.Email}
	isExist, _ := s.FindOneRecordByQuery(queryMap)
	if isExist.ID != 0 {
		return User{}, util.BadRequestErr("email already exist")
	}

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
