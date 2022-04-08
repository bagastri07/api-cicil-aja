package validator

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ApiError struct {
	Field string
	Msg   string
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ApiError, len(ve))
			for i, fe := range ve {
				out[i] = ApiError{fe.Field(), messageForTag(fe.Param(), (fe.Tag()))}
			}
			return echo.NewHTTPError(http.StatusBadRequest, gin.H{"errors": out})
		}
	}
	return nil
}

func messageForTag(param string, tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "min":
		return fmt.Sprintf("This field should have Min %s digit character", param)
	case "max":
		return fmt.Sprintf("This field should have Max %s digit character", param)
	case "e164":
		return "This field should have correct number format (e164):"
	case "oneof":
		return fmt.Sprintf("This field should one of [%s]", param)
	case "url":
		return "this field should be valid url, example: https://example.com"
	}
	return ""
}

func Init(e *echo.Echo) {
	validate := validator.New()

	e.Validator = &CustomValidator{validator: validate}
}
