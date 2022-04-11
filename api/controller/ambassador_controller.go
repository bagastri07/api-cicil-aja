package controller

import (
	"net/http"

	"github.com/bagastri07/api-cicil-aja/api/model"
	"github.com/bagastri07/api-cicil-aja/api/repository"
	"github.com/bagastri07/api-cicil-aja/api/token"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type AmbassadorController struct {
	ambassadorRepository *repository.AmbassadorRepository
}

func NewAmbassadorController() *AmbassadorController {
	return &AmbassadorController{
		ambassadorRepository: repository.NewAmbassadorReposotory(),
	}
}

func (ctl *AmbassadorController) HandleRegisterAsAmbassador(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*token.JwtCustomClaims)

	result, err := ctl.ambassadorRepository.RegisterAsAmbassador(claims.ID)

	if err != nil {
		return err
	}

	resp := model.MessageDataResponse{
		Message: "register success",
		Data:    result,
	}

	return c.JSON(http.StatusOK, resp)
}

// ============ Admin Controller ==============

func (ctl *AmbassadorController) HandleUpdateRegistrationStatusForAdmin(c echo.Context) error {
	registartionID := c.Param("registrationID")

	payload := new(model.UpdateRegistrationStatus)

	if err := c.Bind(payload); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if err := c.Validate(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := ctl.ambassadorRepository.UpdateAmbassadorRegistrationStatusForAdmin(registartionID, payload); err != nil {
		return err
	}

	resp := model.MessageResponse{
		Message: "updated with status " + payload.Status,
	}

	return c.JSON(http.StatusOK, resp)
}

func (ctl *AmbassadorController) HandleGetAllAmbassadorRegistrationsForAdmin(c echo.Context) error {
	status := c.QueryParam("status")

	result, err := ctl.ambassadorRepository.GetAllAmbassadorRegistrationForAdmin(status)

	if err != nil {
		return err
	}

	resp := model.DataResponse{
		Data: result,
	}

	return c.JSON(http.StatusOK, resp)
}

func (ctl *AmbassadorController) HandleGetAcceptedAmbassadorForAdmin(c echo.Context) error {
	result, err := ctl.ambassadorRepository.GetAllAcceptedAmbassadorForAdmin()

	if err != nil {
		return err
	}

	resp := &model.DataResponse{
		Data: result,
	}

	return c.JSON(http.StatusOK, resp)
}

func (ctl *AmbassadorController) HandleGetAllAmbassadorsAndNumberOfTickets(c echo.Context) error {
	result, err := ctl.ambassadorRepository.GetAllAmbassadorsWithTheNumberOfTicket()

	if err != nil {
		return err
	}

	resp := &model.DataResponse{
		Data: result,
	}

	return c.JSON(http.StatusOK, resp)
}
