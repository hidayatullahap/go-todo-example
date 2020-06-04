package errors

import (
	"errors"
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
)

const (
	Expected Bind = iota
	Got
	Field
	Offset
	Error
)

type Bind int

func (b Bind) String() string {
	return [...]string{"expected", "got", "field", "offset", "error"}[b]
}

const (
	ErrTypeUndefined BindErrorType = iota
	ErrTypeSyntax
	ErrTypeMarshal
)

type BindErrorType int

func (b BindErrorType) GetType(s string) BindErrorType {
	var errType BindErrorType

	if strings.Contains(s, "syntax error") {
		errType = ErrTypeSyntax
	} else if strings.Contains(s, "unmarshal type error") {
		errType = ErrTypeMarshal
	} else {
		errType = ErrTypeUndefined
	}

	return errType
}

func (b BindErrorType) IsProperType() bool {
	if b == ErrTypeUndefined {
		return false
	}

	return true
}

type BindError struct {
	Expected     string
	Got          string
	Field        string
	Offset       string
	ErrorMessage string
	Type         BindErrorType
}

func (b BindError) ToString() string {
	var message string

	switch b.Type {
	case ErrTypeMarshal:
		message = fmt.Sprintf("request body %s expected %s got %s", b.Field, b.Expected, b.Got)
	case ErrTypeSyntax:
		message = fmt.Sprintf("bad request json body format: %s", b.ErrorMessage)
	default:
		message = "bind: error type undefined"
	}

	return message
}

func NewErrorBind(err error) error {
	if httpErr, ok := err.(*echo.HTTPError); ok {
		bindErr, errConvert := ConvertHttpErrorToBindError(httpErr)
		if errConvert == nil {
			err = errors.New(bindErr.ToString())
		}
	}

	return err
}

func ConvertHttpErrorToBindError(httpErr *echo.HTTPError) (*BindError, error) {
	var bindErr BindError
	var httpErrMessage string

	if v, ok := httpErr.Message.(string); ok {
		httpErrMessage = strings.ToLower(v)
	} else {
		return nil, httpErr
	}

	bindErr.Type = bindErr.Type.GetType(httpErrMessage)

	if !bindErr.Type.IsProperType() {
		return nil, errors.New("cannot convert to BindError")
	}

	s := strings.Split(httpErrMessage, ":")
	for _, messages := range s {
		message := strings.Split(messages, ",")

		for _, items := range message {
			items = strings.TrimSpace(items)
			item := strings.Split(items, "=")
			if len(item) > 1 {
				key := item[0]
				value := item[1]

				switch key {
				case Expected.String():
					bindErr.Expected = value
				case Got.String():
					bindErr.Got = value
				case Field.String():
					bindErr.Field = value
				case Offset.String():
					bindErr.Offset = value
				case Error.String():
					bindErr.ErrorMessage = value
				}
			}
		}
	}

	return &bindErr, nil
}
