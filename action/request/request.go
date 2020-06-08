package request

// Example for using other validator rule
// `json:"value" validate:"min=1,max=10,required"`
type TodoCreateRequest struct {
	Message string   `json:"message" validate:"required"`
	Note    *string  `json:"note"`
	Tags    *[]int32 `json:"tags"`
}

type TagCreateRequest struct {
	Name string `json:"name" validate:"required"`
}
