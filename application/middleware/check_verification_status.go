package middleware

import (
	"net/http"

	"github.com/bagastri07/api-cicil-aja/api/model"
	"github.com/bagastri07/api-cicil-aja/api/repository"
	"github.com/bagastri07/api-cicil-aja/api/token"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CheckVerificationStatus() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userToken := c.Get("user").(*jwt.Token)
			claims := userToken.Claims.(*token.JwtCustomClaims)

			borrowerRepository := repository.NewBorrowerRepository()

			borrower, err := borrowerRepository.FindForrowerByEmail(claims.Email)

			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err)
			}

			if borrower.VerifiedAt == nil {
				return echo.NewHTTPError(http.StatusBadRequest, &model.MessageResponse{
					Message: "borrower is unverified",
				})
			}
			return next(c)
		}
	}
}
