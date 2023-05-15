package entity

import (
	"gorm.io/gorm"
	"time"
)

type Topic struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	TopicName string         `json:"topic_name" gorm:"uniqueIndex:compositeindex;index;not null"`
	CreatedAt time.Time      `json:"created_at" gorm:"default:CURRENT_TIMESTAMP()"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP()"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
