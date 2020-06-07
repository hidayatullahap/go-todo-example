package core

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	var errMessages string
	err := cv.Validator.Struct(i)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errMessage := fmt.Sprintf(`request body %s %s %s %s; `, strings.ToLower(err.Field()), err.Tag(), err.Type(), err.Param())
			errMessages += errMessage
		}
	}

	if errMessages != "" {
		err = errors.New(errMessages)
	}

	return err
}

func NewCustomValidator() *CustomValidator {
	return &CustomValidator{Validator: validator.New()}
}
