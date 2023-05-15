package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-test/config"
	"go-test/services"
	"go-test/utils"
	"net/http"
	"strings"
)

func DeserializeUser(userService services.UserService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var access_token string
		cookie := ctx.Cookies("access_token")

		authorizationHeader := ctx.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			access_token = fields[1]
		} else if cookie != "" {
			access_token = cookie
		}

		if access_token == "" {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"status":  "fail",
				"message": "You are not logged in",
			})
		}

		config, _ := config.LoadConfig(".")
		sub, err := utils.ValidateToken(access_token, config.AccessTokenPublicKey)

		if err != nil {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"status":  "fail",
				"message": err.Error(),
			})
		}

		user, err := userService.FindUserById(fmt.Sprint(sub))

		if err != nil {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"status":  "fail",
				"message": "The user belonging to this token no longer exists",
			})
		}

		ctx.Locals("currentUser", user)
		return ctx.Next()
	}
}
