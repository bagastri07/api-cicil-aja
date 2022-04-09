package repository

import (
	"errors"
	"net/http"
	"time"

	"github.com/bagastri07/api-cicil-aja/api/model"
	"github.com/bagastri07/api-cicil-aja/database"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type LoanTicketrRepository struct {
	dbClient *gorm.DB
}

func NewLoanTicketRepository() *LoanTicketrRepository {
	dbClient := database.GetDBConnection()
	return &LoanTicketrRepository{
		dbClient: dbClient,
	}
}

func (r *LoanTicketrRepository) MakeNewLoanTicket(borrowerID uint64, payload *model.MakeLoanTicketPayload) (*model.LoanTicket, error) {
	borrower := new(model.Borrower)

	if err := r.dbClient.Preload("LoanTickets").First(&borrower, borrowerID).Error; err != nil {
		return nil, err
	}

	loanTicket := model.LoanTicket{
		LoanAmount:         payload.LoanAmount,
		LoanTenureInMonths: payload.LoanTenureInMonths,
		LoanType:           payload.LoanType,
		InterestRate:       payload.InterestRate,
		ItemUrl:            payload.ItemUrl,
		AmbassadorID:       1,
		BorrowerID:         borrowerID,
		LoanTotal:          payload.LoanAmount + (payload.LoanAmount * float64(payload.InterestRate)),
		Status:             "pending",
	}

	r.dbClient.Create(&loanTicket)

	return &loanTicket, nil
}

func (r *LoanTicketrRepository) GetAllLoanTickets(borrowerID uint64, status string) (*model.LoanTickets, error) {
	loanTickets := new(model.LoanTickets)

	query := r.dbClient

	if status == "pending" || status == "accepted" {
		query = query.Where("status = ?", status)
	}

	if err := query.Order("created_at desc").Find(&loanTickets.LoanTickets, "borrower_id", borrowerID).Error; err != nil {
		return nil, err
	}

	return loanTickets, nil
}

func (r *LoanTicketrRepository) GetLoanTicketById(borrowerID uint64, loanTicketID string) (*model.LoanTicket, error) {
	loatTicket := new(model.LoanTicket)

	if err := r.dbClient.Where("borrower_id", borrowerID).First(loatTicket, loanTicketID).Error; err != nil {
		return nil, err
	}

	return loatTicket, nil
}

func (r *LoanTicketrRepository) DeleteLoanTicketById(borrowerID uint64, loanTicketID string) (*model.LoanTicket, error) {
	loatTicket := new(model.LoanTicket)

	if err := r.dbClient.Where("borrower_id", borrowerID).First(loatTicket, loanTicketID).Error; err != nil {
		return nil, err
	}

	r.dbClient.Delete(loatTicket)

	return loatTicket, nil
}

// ============ Admin Repository ==============

func (r *LoanTicketrRepository) GetAllLoanTicketsForAdmin(status string) (*model.LoanTickets, error) {
	loanTickets := new(model.LoanTickets)

	query := r.dbClient

	if status == "pending" || status == "accepted" {
		query = query.Where("status = ?", status)
	}

	if err := query.Order("created_at desc").Find(&loanTickets.LoanTickets).Error; err != nil {
		return nil, err
	}

	return loanTickets, nil
}

func (r *LoanTicketrRepository) GetLoanTicketByIdForAdmin(loanTicketID string) (*model.LoanTicket, error) {
	loatTicket := new(model.LoanTicket)

	result := r.dbClient.First(loatTicket, loanTicketID)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return loatTicket, nil
}

func (r *LoanTicketrRepository) AcceptLoanTicketByIDForAdmin(loanTicketID string) (*model.LoanTicket, error) {
	loanTicket, err := r.GetLoanTicketByIdForAdmin(loanTicketID)

	if err != nil {
		return nil, err
	}

	if loanTicket.AcceptedAt != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, &model.MessageResponse{
			Message: "this loan ticket is already accepted",
		})
	}

	acceptedTime := time.Now()
	loanTicket.AcceptedAt = &acceptedTime
	loanTicket.Status = "accepted"

	loanTicketRepo := NewLoanBillRepository()
	if err := loanTicketRepo.MakeAllBillsByLoanTicketIDForAdmin(loanTicket); err != nil {
		return nil, err
	}

	if err := r.dbClient.Save(&loanTicket).Error; err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return loanTicket, nil
}
