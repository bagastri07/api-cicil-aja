package middleware

import (
	"os"

	"github.com/bagastri07/api-cicil-aja/api/token"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func VerifyToken() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &token.JwtCustomClaims{},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	})
}
