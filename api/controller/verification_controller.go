package controller

import (
	"fmt"
	"net/http"
	"os"

	"github.com/bagastri07/api-cicil-aja/api/model"
	"github.com/bagastri07/api-cicil-aja/api/repository"
	"github.com/bagastri07/api-cicil-aja/api/token"
	"github.com/bagastri07/api-cicil-aja/util"
	"github.com/bagastri07/api-cicil-aja/util/gomail"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type VerificationController struct {
	verificationRepository *repository.VerificationRepository
}

func NewVerificationController() *VerificationController {
	return &VerificationController{
		verificationRepository: repository.NewVerificationRepository(),
	}
}

func (ctl *VerificationController) HandleSendEmailVerification(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*token.JwtCustomClaims)

	verificationToken := util.GenerateRandomString(20, util.STR_ALPHANUMERIC)

	err := ctl.verificationRepository.CreateNewVerification(claims.Email, verificationToken)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data := struct {
		Email string
		Url   string
	}{
		Email: claims.Email,
		Url:   fmt.Sprintf("%s/verifications/verify-borrower?verificationToken=%s&email=%s", os.Getenv("BASE_EMAIL_VERIF_URL"), verificationToken, claims.Email),
	}

	// send email
	gmSvc := gomail.GetEmailService()
	status, err := gmSvc.SendEmailVerification(claims.Email, data)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if !status {
		return echo.NewHTTPError(http.StatusBadGateway, err)
	}

	resp := &model.MessageResponse{
		Message: "email verification send",
	}

	return c.JSON(http.StatusOK, resp)
}

func (ctl *VerificationController) HandleVerifyBorrower(c echo.Context) error {

	verificationToken := c.QueryParam("verificationToken")
	emailBorrower := c.QueryParam("email")

	if err := ctl.verificationRepository.VerifyBorrower(emailBorrower, verificationToken); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.Redirect(http.StatusMovedPermanently, os.Getenv("REDIRECT_VERIFY_URL"))
}
