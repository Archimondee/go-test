package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-test/controllers"
	"go-test/middleware"
	"go-test/services"
	"go-test/utils"
)

type NotificationRouteController struct {
	notificationController controllers.NotificationController
	userService            services.UserService
}

func NewNotificationRouteontroller(notificationController controllers.NotificationController, userService services.UserService) NotificationRouteController {
	return NotificationRouteController{notificationController, userService}
}

func (nc *NotificationRouteController) NotificationRoute(rg *fiber.App) {
	rg.Post("/fcm-tokens", middleware.DeserializeUser(nc.userService), nc.notificationController.PostTokenFcm)
	rg.Post("/fcm-notifications", utils.HandleFileUploadToBucket, nc.notificationController.PostNotification)
}
