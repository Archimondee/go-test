package entity

import (
	"gorm.io/gorm"
	"time"
)

type Notification struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	Body          string         `json:"body" validate:"required,min=3"`
	CreatedAt     time.Time      `json:"created_at" gorm:"default:CURRENT_TIMESTAMP()"`
	CreatedBy     string         `json:"created_by"`
	Data          string         `json:"data"`
	IsRead        uint           `json:"is_read" gorm:"default:0" `
	Title         string         `json:"title"`
	UpdatedAt     time.Time      `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP()"`
	DeletedAt     gorm.DeletedAt `json:"-"`
	AppIdentifier string         `json:"app_identifier"`
	UserId        uint           `json:"user_id"`
	TopicId       uint           `json:"topic_id"`
	Image         string         `json:"image"`
}
