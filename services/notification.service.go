package services

import (
	"github.com/gofiber/fiber/v2"
	"go-test/models/entity"
	"go-test/models/request"
)

type NotificationService interface {
	PostTokenNotification(request *request.TokenCreateRequest) (*entity.FcmTokens, error)
	FindTokenNotification(string, string) (*entity.FcmTokens, error)
	FindToken(string) (*entity.FcmTokens, error)
	PostNotificationMessage(request *request.NotificationRequest, ctx *fiber.Ctx) (*entity.Notification, error)
}
