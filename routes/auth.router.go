package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-test/controllers"
)

type AuthRouteController struct {
	authController controllers.AuthController
}

func NewAuthRouteController(authController controllers.AuthController) AuthRouteController {
	return AuthRouteController{authController}
}

func (rc *AuthRouteController) AuthRoute(rg *fiber.App) {
	router := rg.Group("/auth")

	router.Post("/register", rc.authController.SignUpUser)
	router.Post("/login", rc.authController.SignInUser)
}
