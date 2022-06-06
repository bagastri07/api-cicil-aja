package controller

import (
	"net/http"

	"github.com/bagastri07/api-cicil-aja/api/model"
	"github.com/bagastri07/api-cicil-aja/api/repository"
	"github.com/bagastri07/api-cicil-aja/api/token"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type ComissionController struct {
	comissionRepository *repository.ComissionRepository
}

func NewCommissionTransactionController() *ComissionController {
	return &ComissionController{
		comissionRepository: repository.NewComissionTransactionRepository(),
	}
}

func (ctl *ComissionController) HandleGetBalanceDetailAmbassador(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*token.JwtCustomClaims)

	result, err := ctl.comissionRepository.GetComissionBalanceAmbassador(claims.ID)

	if err != nil {
		return err
	}

	resp := &model.DataResponse{
		Data: result,
	}

	return c.JSON(http.StatusOK, resp)
}

func (ctl *ComissionController) HandleGetAllComissionHistory(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*token.JwtCustomClaims)

	result, err := ctl.comissionRepository.GetAllComissionHistory(claims.ID)

	if err != nil {
		return err
	}

	resp := &model.DataResponse{
		Data: result,
	}

	return c.JSON(http.StatusOK, resp)
}
