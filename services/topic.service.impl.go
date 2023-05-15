package services

import (
	"context"
	"fmt"
	"go-test/database"
	"go-test/models/entity"
	"go-test/models/request"
)

type TopicServiceImpl struct {
	ctx context.Context
}

func NewTopicServiceImpl(ctx context.Context) TopicService {
	return &TopicServiceImpl{ctx}
}

func (nt *TopicServiceImpl) FindTopicById(topicId uint) (*entity.Topic, error) {
	var topicData *entity.Topic
	fmt.Println("error in here")
	result := database.DB.First(&topicData, topicId)

	if result.Error != nil {
		return nil, result.Error
	}

	return topicData, nil
}

func (nt *TopicServiceImpl) FindTopicByName(topicName string) (*entity.Topic, error) {
	var topic *entity.Topic
	result := database.DB.First(&topic, "topic_name = ?", topicName)

	if result.Error != nil {
		return nil, result.Error
	}

	return topic, nil
}

func (nt *TopicServiceImpl) PostTopic(topic *request.TopicCreateRequest) (*entity.Topic, error) {
	topic.TopicName = topic.TopicName

	result := database.DB.Create(&topic)

	if result.Error != nil {
		return nil, result.Error
	}

	topicResult, err := nt.FindTopicByName(topic.TopicName)

	if err != nil {
		return nil, err
	}

	return topicResult, nil

	//return nil, nil
}
