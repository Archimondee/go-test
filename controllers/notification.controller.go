package controllers

import (
	firebase "firebase.google.com/go"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-test/database"
	"go-test/models/entity"
	"go-test/models/request"
	"go-test/models/response"
	"go-test/services"
	"go-test/utils"
	"math"
	"strconv"
)

type NotificationController struct {
	userService         services.UserService
	notificationService services.NotificationService
	Firebase            *firebase.App
}

func NewTokenController(userService services.UserService, notificationService services.NotificationService, Firebase *firebase.App) NotificationController {
	return NotificationController{userService, notificationService, Firebase}

}

func (nc *NotificationController) PostTokenFcm(ctx *fiber.Ctx) error {
	token := new(request.TokenCreateRequest)

	if err := ctx.BodyParser(token); err != nil {
		return err
	}

	errors := utils.ValidateStruct(token)
	if errors != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Errors",
			"errors":  errors,
			"status":  fiber.StatusInternalServerError,
		})
	}

	result, err := nc.notificationService.PostTokenNotification(token)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error",
			"status":  fiber.StatusInternalServerError,
			"error":   err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    result,
	})
}

func (nc *NotificationController) PostNotification(ctx *fiber.Ctx) error {
	notification := new(request.NotificationRequest)

	if err := ctx.BodyParser(notification); err != nil {
		return err
	}

	errors := utils.ValidateStruct(notification)
	if errors != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Errors",
			"errors":  errors,
			"status":  fiber.StatusInternalServerError,
		})
	}

	result, err := nc.notificationService.PostNotificationMessage(notification, ctx)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error",
			"status":  fiber.StatusInternalServerError,
			"error":   err.Error(),
		})
	}
	errorToken := utils.SendToToken(nc.Firebase, notification.Token)
	if errorToken != nil {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Success",
			"error":   errorToken,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    result,
	})
}

func (nc *NotificationController) GetNotification(ctx *fiber.Ctx) error {
	currentUser := ctx.Locals("currentUser")

	user, err := currentUser.(*entity.User)
	fmt.Println(user.ID)
	var notification []*response.NotificationResponse

	if !err {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Errors",
			"errors":  "Error parse",
			"status":  fiber.StatusInternalServerError,
		})
	}

	sql := fmt.Sprintf(`SELECT * FROM notifications WHERE user_id = %d`, user.ID)

	//if identifier := ctx.Query("%s ")

	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	perPage := 10
	var total int64
	var meta response.Meta

	database.DB.Preload("User").Model(&entity.Notification{}).Count(&total).Where("user_id = ?", user.ID)
	sql = fmt.Sprintf("%s LIMIT %d OFFSET %d", sql, perPage, (page-1)*perPage)

	result := database.DB.Raw(sql).Scan(&notification)

	meta.Page = page
	meta.TotalData = int(total)
	meta.TotalPages = int(math.Ceil(float64(total / int64(perPage))))

	if result.Error != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Errors",
			"errors":  result.Error,
			"status":  fiber.StatusInternalServerError,
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "Success",
		"data":    notification,
		"status":  fiber.StatusOK,
		"meta":    meta,
	})
}

func (nc *NotificationController) GetNotificationById(ctx *fiber.Ctx) error {
	notificationId := ctx.Params("id")
	var notification *response.NotificationResponse

	currentUser := ctx.Locals("currentUser")

	user, err := currentUser.(*entity.User)
	if !err {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Errors",
			"errors":  "Error parse",
			"status":  fiber.StatusInternalServerError,
		})
	}

	result := database.DB.Preload("User").First(&notification, "user_id = ? AND id = ?", user.ID, notificationId)

	if result.Error != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Errors",
			"errors":  result.Error,
			"status":  fiber.StatusInternalServerError,
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "Errors",
		"data":    notification,
		"status":  fiber.StatusOK,
	})
}

func (nc *NotificationController) UpdateReadNotification(ctx *fiber.Ctx) error {
	isRead := new(request.NotificationUpdate)
	id := ctx.Params("id")
	var notification *entity.Notification
	if err := ctx.BodyParser(isRead); err != nil {
		return err
	}

	currentUser := ctx.Locals("currentUser")

	user, err := currentUser.(*entity.User)
	if !err {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Errors",
			"errors":  "Error parse",
			"status":  fiber.StatusInternalServerError,
		})
	}

	errors := utils.ValidateStruct(isRead)
	if errors != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Errors",
			"errors":  errors,
			"status":  fiber.StatusInternalServerError,
		})
	}

	fmt.Println("data", isRead.IsRead, id)

	result := database.DB.First(&notification, "user_id = ? AND id = ?", user.ID, id)

	if result.Error != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Errors nih",
			"errors":  result.Error.Error(),
			"status":  fiber.StatusInternalServerError,
		})
	}
	notification.IsRead = isRead.IsRead

	resData := database.DB.Save(&notification)

	if resData.Error != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Errors",
			"errors":  resData.Error,
			"status":  fiber.StatusInternalServerError,
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "Success",
		"status":  fiber.StatusOK,
	})
}

func (nc *NotificationController) GetCounter(ctx *fiber.Ctx) error {
	currentUser := ctx.Locals("currentUser")

	user, err := currentUser.(*entity.User)

	if !err {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Errors",
			"errors":  "Error parse",
			"status":  fiber.StatusInternalServerError,
		})
	}

	var total int64
	result := database.DB.Preload("User").Model(&entity.Notification{}).Count(&total).Where("user_id = ? AND is_read = 0", user.ID)

	if result.Error != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Errors",
			"errors":  result.Error,
			"status":  fiber.StatusInternalServerError,
		})
	}
	return ctx.Status(200).JSON(fiber.Map{
		"message": "Success",
		"counter": total,
		"status":  fiber.StatusOK,
	})
}
