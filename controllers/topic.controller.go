package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-test/models/request"
	"go-test/services"
	"go-test/utils"
)

type TopicController struct {
	topicService services.TopicService
}

func NewTopicController(topicService services.TopicService) TopicController {
	return TopicController{topicService}
}

func (nc *TopicController) PostTopic(ctx *fiber.Ctx) error {
	topic := new(request.TopicCreateRequest)

	if err := ctx.BodyParser(topic); err != nil {

		return err
	}

	errors := utils.ValidateStruct(topic)

	if errors != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Errors",
			"errors":  errors,
			"status":  fiber.StatusInternalServerError,
		})
	}

	result, errPost := nc.topicService.PostTopic(topic)

	if errPost != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error",
			"status":  fiber.StatusInternalServerError,
			"error":   errPost.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"status":  fiber.StatusOK,
		"data":    result,
	})
}
