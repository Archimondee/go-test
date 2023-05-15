package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-test/services"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return UserController{userService}
}

func (uc *UserController) GetMe(ctx *fiber.Ctx) error {
	currentUser := ctx.Locals("currentUser")

	return ctx.Status(200).JSON(fiber.Map{"status": fiber.StatusOK, "message": "Success", "data": currentUser})
}
