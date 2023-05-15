package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-test/controllers"
)

type TopicRouteController struct {
	topicController controllers.TopicController
}

func NewTopicRouteontroller(topicController controllers.TopicController) TopicRouteController {
	return TopicRouteController{topicController}
}

func (nc *TopicRouteController) TopicRoute(rg *fiber.App) {
	route := rg.Group("topic")
	route.Post("/add", nc.topicController.PostTopic)
}
