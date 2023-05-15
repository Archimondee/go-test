package services

import (
	"go-test/models/entity"
	"go-test/models/request"
)

type TopicService interface {
	FindTopicById(uint) (*entity.Topic, error)
	FindTopicByName(string) (*entity.Topic, error)
	PostTopic(request *request.TopicCreateRequest) (*entity.Topic, error)
}
