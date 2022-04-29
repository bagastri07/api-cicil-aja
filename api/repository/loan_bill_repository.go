package repository

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/bagastri07/api-cicil-aja/api/model"
	"github.com/bagastri07/api-cicil-aja/constant"
	"github.com/bagastri07/api-cicil-aja/database"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type LoanBillRepository struct {
	dbClient *gorm.DB
}

func NewLoanBillRepository() *LoanBillRepository {
	return &LoanBillRepository{
		dbClient: database.GetDBConnection(),
	}
}

func (r *LoanBillRepository) GetAllLoanBill(borrowerID uint64, ticketID, status string) (*model.LoanBills, error) {
	loanBills := new(model.LoanBills)

	query := r.dbClient
	if status == "paid" || status == "unpaid" {
		query = query.Where("status = ?", status)
	}

	if ticketID != "" {
		query = query.Where("loan_ticket_id = ?", ticketID)
	}

	if err := query.Find(&loanBills.LoanBills, "borrower_id = ?", borrowerID).Error; err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return loanBills, nil
}

func (r *LoanBillRepository) PayLoanBillByID(borrowerID uint64, LoanBillID string) (*model.LoanBill, error) {
	loanBill := new(model.LoanBill)

	if err := r.dbClient.Where("borrower_id", borrowerID).First(&loanBill, LoanBillID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if loanBill.PaidAt != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, &model.MessageResponse{
			Message: "this loan bill already paid",
		})
	}

	now := time.Now()
	loanBill.PaidAt = &now
	loanBill.Status = "paid"

	if err := r.dbClient.Save(&loanBill).Error; err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// search AmbassadorID by loan Bill
	loanTicket := new(model.LoanTicket)
	if err := r.dbClient.Preload("LoanBills").First(loanTicket, loanBill.LoanTicketID).Error; err != nil {
		return nil, err
	}

	loanBillRepo := NewComissionTransactionRepository()
	comissionAmount := constant.COMISSION_RATE * loanBill.BillAmount
	if err := loanBillRepo.CreateComissionTrasaction(*loanTicket.AmbassadorID, comissionAmount, "in"); err != nil {
		return nil, err
	}

	return loanBill, nil
}

// ============ Admin Repository ==============

func (r *LoanBillRepository) MakeAllBillsByLoanTicketIDForAdmin(loanTicket *model.LoanTicket) error {
	// Make Bills
	loanTenure, _ := strconv.Atoi(loanTicket.LoanTenureInMonths)
	billAmount := loanTicket.LoanTotal / float64(loanTenure)
	billTime := time.Now()

	var loanBills []model.LoanBill

	for i := 0; i < loanTenure; i++ {
		billTime = billTime.AddDate(0, 1, 0)
		loanBill := &model.LoanBill{
			BorrowerID:      loanTicket.BorrowerID,
			LoanTicketID:    loanTicket.ID,
			PaymentDeadline: billTime,
			BillAmount:      billAmount,
			Status:          "unpaid",
		}
		loanBills = append(loanBills, *loanBill)
	}

	r.dbClient.Model(loanTicket).Association("LoanBills").Append(loanBills)

	return nil
}
