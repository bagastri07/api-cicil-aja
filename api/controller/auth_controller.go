package controller

import (
	"net/http"

	"github.com/bagastri07/api-cicil-aja/api/model"
	"github.com/bagastri07/api-cicil-aja/api/repository"
	"github.com/bagastri07/api-cicil-aja/api/token"
	"github.com/bagastri07/api-cicil-aja/util"
	googleOauth2 "github.com/bagastri07/api-cicil-aja/util/oauth2"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	borrowerRepository *repository.BorrowerRepository
}

func NewAuthController() *AuthController {
	return &AuthController{
		borrowerRepository: repository.NewBorrowerRepository(),
	}
}

func (ctl *AuthController) BorrowerLogin(c echo.Context) error {
	loginCredential := new(model.LoginRequest)

	c.Bind(loginCredential)

	borrower, err := ctl.borrowerRepository.FindForrowerByEmail(loginCredential.Email)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if loginCredential.Email != borrower.Email || !util.CheckPasswordHash(loginCredential.Password, borrower.Password) {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	token, err := token.GenerateToken(borrower.ID, borrower.Email, "borrower")

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, &model.DataResponse{
		Data: token,
	})
}

func (ctl *AuthController) BorrowerLoginWithGoogle(c echo.Context) error {
	gsrv := googleOauth2.GetNewOuathService()
	return gsrv.GoogleLogin(c)
}

func (ctl *AuthController) BorrowerLoginGoogleCallback(c echo.Context) error {
	gsrv := googleOauth2.GetNewOuathService()
	return gsrv.GoogleCallback(c)
}
