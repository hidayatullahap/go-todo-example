package request

type TodoCreateRequest struct {
	Message    string   `json:"message" validate:"required"`
	Note       *string  `json:"note"`
	CustomDate *string  `json:"custom_date"`
	Tags       *[]int32 `json:"tags"`
}

type TagCreateRequest struct {
	Name string `json:"name" validate:"required"`
}
