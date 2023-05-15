package entity

import "time"

type FcmTokens struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	CreatedAt     time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP()"`
	Os            string    `json:"os"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP()"`
	UserId        string    `json:"user_id"`
	Token         string    `json:"token"`
	AppIdentifier string    `json:"app_identifier"`
}
