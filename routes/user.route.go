package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-test/controllers"
	"go-test/middleware"
	"go-test/services"
)

type UserRouteController struct {
	userController         controllers.UserController
	userService            services.UserService
	notificationController controllers.NotificationController
}

func NewUserRouteController(userController controllers.UserController, userService services.UserService, notificationController controllers.NotificationController) UserRouteController {
	return UserRouteController{userController, userService, notificationController}
}

func (rc *UserRouteController) UserRoute(rg *fiber.App) {
	router := rg.Group("/user")
	router.Get("/me", middleware.DeserializeUser(rc.userService), rc.userController.GetMe)
	router.Get("/notifications", middleware.DeserializeUser(rc.userService), rc.notificationController.GetNotification)
	router.Get("/notifications/:id", middleware.DeserializeUser(rc.userService), rc.notificationController.GetNotificationById)
	router.Put("/notifications/:id", middleware.DeserializeUser(rc.userService), rc.notificationController.UpdateReadNotification)
	router.Get("/counter-notifications", middleware.DeserializeUser(rc.userService), rc.notificationController.GetCounter)
}
