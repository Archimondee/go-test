package services

import "go-test/models/entity"

type UserService interface {
	FindUserByEmail(string) (*entity.User, error)
	FindUserById(string) (*entity.User, error)
}
