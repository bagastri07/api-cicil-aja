package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

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
		return echo.NewHTTPError(http.StatusBadRequest, err)
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

func (ctl *BorrowerController) HandleUpdateBorrowerBankAccount(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*token.JwtCustomClaims)
	fmt.Print(claims)

	payload := new(model.UpdateBorrowerBankAccount)

	if err := c.Bind(payload); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if err := c.Validate(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	result, err := ctl.borrowerRepository.UpdateBorrowerBankAccount(payload, claims.ID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resp := &model.DataResponse{
		Data: result,
	}

	return c.JSON(http.StatusOK, resp)
}

func (ctl *BorrowerController) HandleGetCurrentBorrower(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*token.JwtCustomClaims)

	result, err := ctl.borrowerRepository.FindBorrowerByID(claims.ID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusAccepted, result)
}

func (ctl *BorrowerController) HandleUploadBorrowerDocument(c echo.Context) error {
	// Read file
	file, err := c.FormFile("img_ktm")

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	src, err := file.Open()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	defer src.Close()

	// Destination
	filename := util.GenerateRandomString(30, util.STR_ALPHANUMERIC) + filepath.Ext(file.Filename)
	filePath := filepath.Join("img", filepath.Base(filename))
	dst, err := os.Create(filepath.Join("public", filePath))
	if err != nil {
		return err
	}

	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*token.JwtCustomClaims)

	result, err := ctl.borrowerRepository.UploadKtmImage(filePath, claims.ID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resp := &model.DataResponse{
		Data: result,
	}

	return c.JSON(http.StatusOK, resp)
}
