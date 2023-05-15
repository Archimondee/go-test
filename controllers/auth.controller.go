package controllers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go-test/models/request"
	"go-test/services"
	"go-test/utils"
)

type AuthController struct {
	authService services.AuthService
	ctx         context.Context
}

func NewAuthController(authService services.AuthService, ctx context.Context) AuthController {
	return AuthController{authService, ctx}
}

func (ac *AuthController) SignUpUser(ctx *fiber.Ctx) error {
	user := new(request.UserCreateRequest)

	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	errors := utils.ValidateStruct(user)
	if errors != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Errors",
			"error":   errors,
			"status":  fiber.StatusInternalServerError,
		})
	}

	newUser, err := ac.authService.SignUpUser(user)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to store data",
			"status":  fiber.StatusInternalServerError,
			"error":   err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Success",
		"data":    newUser,
	})
}

func (ac *AuthController) SignInUser(ctx *fiber.Ctx) error {
	user := new(request.UserSigninRequest)

	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	errors := utils.ValidateStruct(user)
	if errors != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Errors",
			"error":   errors,
			"status":  fiber.StatusInternalServerError,
		})
	}

	token, err := ac.authService.SignInUser(user)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error",
			"status":  fiber.StatusInternalServerError,
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"token":   token,
	})
}
