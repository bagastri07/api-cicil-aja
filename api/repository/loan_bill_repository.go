package repository

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/bagastri07/api-cicil-aja/api/model"
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
	fmt.Println(status == "paid")

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
