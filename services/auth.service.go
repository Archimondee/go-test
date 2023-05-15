package services

import (
	"go-test/models/entity"
	"go-test/models/request"
)

type AuthService interface {
	SignUpUser(input *request.UserCreateRequest) (*entity.User, error)
	SignInUser(input *request.UserSigninRequest) (*string, error)
}
