package controller

import (
	"net/http"

	"github.com/bagastri07/api-cicil-aja/api/model"
	"github.com/bagastri07/api-cicil-aja/api/repository"
	"github.com/bagastri07/api-cicil-aja/api/token"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type LoanBillController struct {
	loanBillRepository *repository.LoanBillRepository
}

func NewLoanBillController() *LoanBillController {
	return &LoanBillController{
		loanBillRepository: repository.NewLoanBillRepository(),
	}
}

func (ctl *LoanBillController) GetAllLoanBill(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*token.JwtCustomClaims)

	status := c.QueryParam("status")
	ticketID := c.QueryParam("ticketID")
	result, err := ctl.loanBillRepository.GetAllLoanBill(claims.ID, ticketID, status)

	if err != nil {
		return err
	}

	resp := &model.DataResponse{
		Data: result,
	}

	return c.JSON(http.StatusOK, resp)
}
