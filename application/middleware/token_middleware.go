package middleware

import (
	"net/http"
	"os"

	"github.com/bagastri07/api-cicil-aja/api/model"
	"github.com/bagastri07/api-cicil-aja/api/repository"
	"github.com/bagastri07/api-cicil-aja/api/token"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func VerifyToken() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &token.JwtCustomClaims{},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	})
}

func IsAdmin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userToken := c.Get("user").(*jwt.Token)
			claims := userToken.Claims.(*token.JwtCustomClaims)

			if !claims.Admin {
				return echo.NewHTTPError(http.StatusForbidden, &model.MessageResponse{
					Message: "you are not admin",
				})
			}
			return next(c)
		}
	}
}

func IsAmbassador() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userToken := c.Get("user").(*jwt.Token)
			claims := userToken.Claims.(*token.JwtCustomClaims)

			borrowerRepo := repository.NewBorrowerRepository()

			result, err := borrowerRepo.FindBorrowerByID(claims.ID)

			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}

			if !result.IsAmbassador {
				return echo.NewHTTPError(http.StatusForbidden, &model.MessageResponse{
					Message: "you are not ambassador",
				})
			}
			return next(c)
		}
	}
}
