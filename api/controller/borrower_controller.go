package controller

import (
	"net/http"

	"github.com/bagastri07/api-cicil-aja/api/model"
	"github.com/bagastri07/api-cicil-aja/api/repository"
	"github.com/bagastri07/api-cicil-aja/api/token"
	"github.com/bagastri07/api-cicil-aja/util"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type BorrowerController struct {
	borrowerRepository *repository.BorrowerRepository
}

func NewBorrowerController() *BorrowerController {
	return &BorrowerController{
		borrowerRepository: repository.NewBorrowerRepository(),
	}
}

func (ctl *BorrowerController) HandleGetBorrowerByEmail(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*token.JwtCustomClaims)

	borrower, err := ctl.borrowerRepository.GetBorrowerByEmail(claims.Email)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &model.DataResponse{
		Data: borrower,
	})
}

func (ctl *BorrowerController) HandleCreateNewBorrower(c echo.Context) error {
	borrower := new(model.Borrower)

	if err := c.Bind(borrower); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if err := c.Validate(borrower); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	hashedPasword, err := util.HashPassword(borrower.Password)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	borrower.Password = hashedPasword

	if err := ctl.borrowerRepository.CreateBorrower(borrower); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	resp := &model.DataResponse{
		Data: borrower,
	}

	return c.JSON(http.StatusCreated, resp)
}

func (ctl *BorrowerController) HandleUpdateBorrower(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*token.JwtCustomClaims)

	updatedBorrower := new(model.UpdateBorrower)

	if err := c.Bind(updatedBorrower); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if err := c.Validate(updatedBorrower); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	result, err := ctl.borrowerRepository.UpdateBorrower(updatedBorrower, claims.Email)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	resp := &model.DataResponse{
		Data: result,
	}

	return c.JSON(http.StatusCreated, resp)
}
