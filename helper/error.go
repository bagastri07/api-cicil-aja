package helper

import (
	"github.com/labstack/echo/v4"
)

func ErrorParsing(httpStatus int, err error) error {
	return echo.NewHTTPError(httpStatus, err.Error())
}
