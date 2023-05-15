package services

import (
	"context"
	"go-test/config"
	"go-test/database"
	"go-test/models/entity"
	"go-test/models/request"
	"go-test/utils"
	"strings"
)

type AuthServiceImpl struct {
	ctx         context.Context
	userService UserService
}

func NewAuthService(ctx context.Context, userService UserService) AuthService {
	return &AuthServiceImpl{ctx, userService}
}

func (uc *AuthServiceImpl) SignUpUser(user *request.UserCreateRequest) (*entity.User, error) {
	user.Email = strings.ToLower(user.Email)
	user.Name = user.Name
	user.Phone = user.Phone
	user.Address = user.Address
	hashedPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashedPassword

	result := database.DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	var newUser *entity.User
	res := database.DB.First(&newUser, "email = ?", user.Email)
	if res.Error != nil {
		return nil, result.Error
	}

	return newUser, nil
}

func (uc *AuthServiceImpl) SignInUser(user *request.UserSigninRequest) (*string, error) {

	userSignIn, err := uc.userService.FindUserByEmail(user.Email)

	if err != nil {
		return nil, err
	}

	errPassword := utils.VerifyPassword(userSignIn.Password, user.Password)
	if errPassword != nil {
		return nil, errPassword
	}

	config, _ := config.LoadConfig(".")
	access_token, err := utils.CreateToken(userSignIn.ID, config.AccessTokenPrivateKey)

	return &access_token, nil
}
