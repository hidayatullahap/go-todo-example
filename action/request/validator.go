package request

import "github.com/go-playground/validator"

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func NewCustomValidator() *CustomValidator {
	return &CustomValidator{Validator: validator.New()}
}
