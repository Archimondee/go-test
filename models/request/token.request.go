package request

type TokenCreateRequest struct {
	Os            string `json:"os" validate:"required"`
	UserId        string `json:"user_id" validate:"required"`
	Token         string `json:"token" validate:"required"`
	AppIdentifier string `json:"app_identifier" validate:"required"`
}

func (TokenCreateRequest) TableName() string {
	return "fcm_tokens"
}
