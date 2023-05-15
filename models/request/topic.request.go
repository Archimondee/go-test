package request

type TopicCreateRequest struct {
	TopicName string `json:"topic_name" gorm:"uniqueIndex:compositeindex;index;not null"`
}

func (TopicCreateRequest) TableName() string {
	return "topics"
}
