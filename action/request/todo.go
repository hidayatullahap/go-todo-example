package request

type TodoCreateRequest struct {
	Message    string  `json:"message"`
	Note       *string `json:"note"`
	CustomDate *string `json:"custom_date"`
}
