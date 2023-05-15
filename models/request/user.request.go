package request

type UserCreateRequest struct {
	Name     string `json:"name" query:"name" validate:"required,min=3,max=32"`
	Email    string `json:"email" query:"email" validate:"required,email"`
	Password string `json:"password" query:"password" validate:"required,min=6"`
	Address  string `json:"address" query:"address"`
	Phone    string `json:"phone" query:"phone"`
}

type UserSigninRequest struct {
	Email    string `json:"email" query:"email" validate:"required,email"`
	Password string `json:"password" query:"password" validate:"required,min=6"`
}

func (UserCreateRequest) TableName() string {
	return "users"
}
