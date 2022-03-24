package controller

import (
	"net/http"
	"strconv"

	"github.com/bagastri07/api-cicil-aja/api/model"
	"github.com/bagastri07/api-cicil-aja/api/repository"
	"github.com/bagastri07/api-cicil-aja/helper"
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

func (ctl *BorrowerController) HandleGetBorrowerByID(c echo.Context) error {
	var resp model.DataResponse

	borrowerID, err := strconv.ParseInt(c.Param("borrowerID"), 10, 64)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"messages": err.Error(),
		})
	}

	borrower, err := ctl.borrowerRepository.GetBorrowerByID(uint64(borrowerID))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"messages": err.Error(),
		})
	}

	resp.Data = borrower

	return c.JSON(http.StatusOK, resp)
}

func (ctl *BorrowerController) HandleCreateNewBorrower(c echo.Context) error {
	borrower := new(model.Borrower)

	if err := c.Bind(borrower); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := c.Validate(borrower); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	hashedPasword, err := helper.HashPassword(borrower.Password)

	if err != nil {
		return helper.ErrorParsing(http.StatusBadRequest, err)
	}

	borrower.Password = hashedPasword

	if err := ctl.borrowerRepository.CreateBorrower(borrower); err != nil {
		return helper.ErrorParsing(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, borrower)
}
