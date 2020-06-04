package request

type TodoCreateRequest struct {
	Message    string  `json:"message"`
	Note       *string `json:"note"`
	CustomDate *string `json:"custom_date"`
}

type TagCreateRequest struct {
	Name string `json:"name" validate:"required"`
}
