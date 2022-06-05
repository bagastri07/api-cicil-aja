package controller

import (
	"net/http"
	"strconv"

	"github.com/bagastri07/api-cicil-aja/api/model"
	"github.com/bagastri07/api-cicil-aja/api/repository"
	"github.com/bagastri07/api-cicil-aja/api/token"
	"github.com/bagastri07/api-cicil-aja/constant"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type LoanTicketController struct {
	e                    *echo.Echo
	loanTicketRepository repository.LoanTicketRepository
}

func NewLoanTicketController(e *echo.Echo, repo repository.LoanTicketRepository) *LoanTicketController {
	return &LoanTicketController{
		e:                    e,
		loanTicketRepository: repo,
	}
}

func (ctl *LoanTicketController) HandleMakeLoanTicket(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*token.JwtCustomClaims)

	payload := new(model.MakeLoanTicketPayload)

	if err := c.Bind(payload); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if err := c.Validate(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	payload.InterestRate = ctl.assignInterestRate(payload.LoanTenureInMonths)

	result, err := ctl.loanTicketRepository.MakeNewLoanTicket(claims.ID, payload)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resp := &model.MessageDataResponse{
		Message: "loan ticket created",
		Data:    result,
	}
	return c.JSON(http.StatusOK, resp)
}

func (ctl *LoanTicketController) assignInterestRate(loanTenure string) float32 {
	if loanTenure == "3" {
		return constant.InterestThreeMonths
	} else if loanTenure == "6" {
		return constant.InterestSixMonths
	} else {
		return constant.InterestTwelveMonths
	}
}

func (ctl *LoanTicketController) HandleGetAllTicket(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*token.JwtCustomClaims)

	status := c.QueryParam("status")

	result, err := ctl.loanTicketRepository.GetAllLoanTickets(claims.ID, status)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resp := &model.DataResponse{
		Data: result,
	}

	return c.JSON(http.StatusOK, resp)
}

func (ctl *LoanTicketController) HandleGetLoanTicketByID(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*token.JwtCustomClaims)

	ticketID := c.Param("loanTicketID")

	result, err := ctl.loanTicketRepository.GetLoanTicketById(claims.ID, ticketID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resp := &model.DataResponse{
		Data: result,
	}
	return c.JSON(http.StatusOK, resp)
}

func (ctl *LoanTicketController) HandleDeleteLoanTicketByID(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*token.JwtCustomClaims)

	ticketID := c.Param("loanTicketID")

	result, err := ctl.loanTicketRepository.DeleteLoanTicketById(claims.ID, ticketID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resp := &model.MessageDataResponse{
		Message: "deleted",
		Data:    result,
	}
	return c.JSON(http.StatusOK, resp)
}

func (ctl *LoanTicketController) HandleCalculateEstimateLoanTicket(c echo.Context) error {
	payload := new(model.CalculateEstimateLoanTicketPayload)

	if err := c.Bind(payload); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if err := c.Validate(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if payload.LoanTenureInMonths == "3" {
		payload.InterestRate = constant.InterestThreeMonths
	} else if payload.LoanTenureInMonths == "6" {
		payload.InterestRate = constant.InterestSixMonths
	} else if payload.LoanTenureInMonths == "12" {
		payload.InterestRate = constant.InterestTwelveMonths
	}

	loanTenure, err := strconv.Atoi(payload.LoanTenureInMonths)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	interest := payload.LoanAmount * payload.InterestRate
	loanTotal := payload.LoanAmount + interest
	monthlyBill := loanTotal / float64(loanTenure)
	result := &model.CalculateEstimateLoanTicketResponse{
		Interest:    interest,
		LoanTotal:   loanTotal,
		MonthlyBill: monthlyBill,
	}

	resp := &model.DataResponse{
		Data: result,
	}

	return c.JSON(http.StatusOK, resp)
}

// ============ Ambassador Controller ==============
func (ctl *LoanTicketController) HandleReviewLoanTicketByAmbassador(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*token.JwtCustomClaims)

	ticketID := c.Param("loanTicketID")

	result, err := ctl.loanTicketRepository.ReviewLoanTikcetByAmbassador(claims.ID, ticketID)

	if err != nil {
		return err
	}

	resp := &model.MessageDataResponse{
		Message: "reviewed",
		Data:    result,
	}

	return c.JSON(http.StatusOK, resp)
}

func (ctl *LoanTicketController) HandleGetAllLoanTicketForAmbassador(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*token.JwtCustomClaims)

	status := c.QueryParam("status")

	result, err := ctl.loanTicketRepository.GetAllLoanTicketsForAmbassador(claims.ID, status)

	if err != nil {
		return err
	}

	resp := &model.DataResponse{
		Data: result,
	}

	return c.JSON(http.StatusOK, resp)
}

func (ctl *LoanTicketController) HandleGetLoanTicketByIDForAmbassador(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*token.JwtCustomClaims)

	ticketID := c.Param("loanTicketID")

	result, err := ctl.loanTicketRepository.GetLoanLoanTicketByIdForAmbassador(claims.ID, ticketID)

	if err != nil {
		return err
	}

	resp := &model.DataResponse{
		Data: result,
	}
	return c.JSON(http.StatusOK, resp)
}

// ============ Admin Controller ==============

func (ctl *LoanTicketController) HandleGetAllTicketForAdmin(c echo.Context) error {
	status := c.QueryParam("status")

	result, err := ctl.loanTicketRepository.GetAllLoanTicketsForAdmin(status)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resp := &model.DataResponse{
		Data: result,
	}

	return c.JSON(http.StatusOK, resp)
}

func (ctl *LoanTicketController) HandleGetLoanTicketByIDForAdmin(c echo.Context) error {
	ticketID := c.Param("loanTicketID")

	result, err := ctl.loanTicketRepository.GetLoanTicketByIdForAdmin(ticketID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resp := &model.DataResponse{
		Data: result,
	}
	return c.JSON(http.StatusOK, resp)
}

func (ctl *LoanTicketController) HandleUpdateStatusLoanTicketByIDForAdmin(c echo.Context) error {
	ticketID := c.Param("loanTicketID")

	payload := new(model.UpdateLoanTicketStatus)

	if err := c.Bind(payload); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if err := c.Validate(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	result, err := ctl.loanTicketRepository.UpdateStatusLoanTicketByIDForAdmin(ticketID, payload.Status)

	if err != nil {
		return err
	}

	resp := &model.DataResponse{
		Data: result,
	}
	return c.JSON(http.StatusOK, resp)
}
