package response

type Meta struct {
	TotalData  int `json:"total_data"`
	TotalPages int `json:"total_pages"`
	Page       int `json:"page"`
}
