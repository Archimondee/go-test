package request

type NotificationRequest struct {
	TopicId uint   `json:"topicId" validate:"required"`
	Body    string `json:"body" validate:"required,min=3"`
	Data    string `json:"data" validate:"required"`
	Title   string `json:"title" validate:"required"`
	Token   string `json:"token" validate:"required"`
	Image   string `json:"image"`
}

type NotificationUpdate struct {
	IsRead uint `json:"is_read"`
}

func (NotificationRequest) TableName() string {
	return "notifications"
}
