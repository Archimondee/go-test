package services

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-test/database"
	"go-test/models/entity"
	"go-test/models/request"
	"strconv"
)

type NotificationServiceImpl struct {
	ctx          context.Context
	userService  UserService
	topicService TopicService
}

func NewNotificationServiceImpl(ctx context.Context, userService UserService, topicService TopicService) NotificationService {
	return &NotificationServiceImpl{ctx, userService, topicService}
}

func (ns *NotificationServiceImpl) PostTokenNotification(fcm *request.TokenCreateRequest) (*entity.FcmTokens, error) {
	fcm.Token = fcm.Token
	fcm.Os = fcm.Os
	fcm.UserId = fcm.UserId
	fcm.AppIdentifier = fcm.AppIdentifier

	_, err := ns.userService.FindUserById(fcm.UserId)

	if err != nil {
		return nil, err
	}

	result := database.DB.Create(&fcm)
	if result.Error != nil {
		return nil, result.Error
	}

	token, err := ns.FindTokenNotification(fcm.UserId, fcm.AppIdentifier)

	if err != nil {
		return nil, result.Error
	}

	return token, nil
}

func (ns *NotificationServiceImpl) FindTokenNotification(userId string, appIdentifier string) (*entity.FcmTokens, error) {
	var fcm *entity.FcmTokens

	result := database.DB.First(&fcm, "user_id = ? AND app_identifier = ?", userId, appIdentifier)
	if result.Error != nil {
		return nil, result.Error
	}

	return fcm, nil
}

func (ns *NotificationServiceImpl) FindToken(token string) (*entity.FcmTokens, error) {
	var fcm *entity.FcmTokens

	result := database.DB.First(&fcm, "token = ?", token)
	if result.Error != nil {
		return nil, result.Error
	}

	return fcm, nil
}

func (ns *NotificationServiceImpl) PostNotificationMessage(message *request.NotificationRequest, ctx *fiber.Ctx) (*entity.Notification, error) {
	//message.Body = message.Body
	//message.Data = message.Data
	//message.Title = message.Title
	//message.TopicId = message.TopicId
	//message.Image = message.Image
	//message.Token = message.Token

	res, err := ns.FindToken(message.Token)

	if err != nil {
		return nil, err
	}
	userId, err := strconv.ParseUint(res.UserId, 10, 64)

	var filenameString string
	filename := ctx.Locals("filename")
	if filename != nil {
		filenameString = fmt.Sprintf("%v", filename)
	} else {
		filenameString = ""
	}

	if message.TopicId != 0 {
		var topicData *entity.Topic

		result := database.DB.First(&topicData, message.TopicId)

		if result.Error != nil {
			return nil, result.Error
		}
	}

	var data = &entity.Notification{
		Title:         message.Title,
		AppIdentifier: res.AppIdentifier,
		Body:          message.Body,
		UserId:        uint(userId),
		Data:          message.Data,
		Image:         filenameString,
		TopicId:       message.TopicId,
	}

	result := database.DB.Create(&data)
	if result.Error != nil {

		return nil, result.Error
	}

	return data, nil
}
