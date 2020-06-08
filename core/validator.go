package core

import (
	"fmt"
	"strings"

	stdErr "errors"

	"github.com/go-playground/validator"
	"github.com/hidayatullahap/go-todo-example/core/errors"
	"github.com/labstack/echo/v4"
)

func BindValidate(c echo.Context, i interface{}) error {
	err := c.Bind(i)
	if err != nil {
		return errors.NewErrorBind(err)
	}

	err = c.Validate(i)
	if err != nil {
		return err
	}

	return nil
}

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	var errMessages string
	err := cv.Validator.Struct(i)
	if err != nil {
		var tmpErrMessages []string
		for _, err := range err.(validator.ValidationErrors) {
			errMessage := fmt.Sprintf(`%s is %s %s %s`, strings.ToLower(err.Field()), err.Tag(), err.Type(), err.Param())
			tmpErrMessages = append(tmpErrMessages, errMessage)
		}

		errMessages = strings.Join(tmpErrMessages, ";")
	}

	if errMessages != "" {
		err = stdErr.New(errMessages)
	}

	return err
}

func NewCustomValidator() *CustomValidator {
	return &CustomValidator{Validator: validator.New()}
}
