package services

import (
	"context"
	"go-test/database"
	"go-test/models/entity"
)

type UserServiceImpl struct {
	ctx context.Context
}

func NewUserServiceImpl(ctx context.Context) UserService {
	return &UserServiceImpl{ctx}
}

func (us *UserServiceImpl) FindUserByEmail(email string) (*entity.User, error) {
	var user *entity.User
	result := database.DB.First(&user, "email = ?", email)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (us *UserServiceImpl) FindUserById(userId string) (*entity.User, error) {
	var user *entity.User
	result := database.DB.First(&user, "id = ?", userId)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
