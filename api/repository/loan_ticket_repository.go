package repository

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/bagastri07/api-cicil-aja/api/model"
	"github.com/bagastri07/api-cicil-aja/database"
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
	}

	r.dbClient.Create(&loanTicket)

	return &loanTicket, nil
}

func (r *LoanTicketrRepository) GetAllLoanTickets(borrowerID uint64, status string) (*model.LoanTickets, error) {
	loanTickets := new(model.LoanTickets)

	query := sq.Select("*").From("loan_tickets")

	if status == "pending" {
		query = query.Where(sq.Eq{"accepted_at": nil})
	} else if status == "accepted" {
		query = query.Where(sq.NotEq{"accepted_at": nil})
	}

	sql, _, _ := query.
		Where(fmt.Sprintf("borrower_id = %d", borrowerID)).
		OrderBy("created_at DESC").
		ToSql()

	if err := r.dbClient.Raw(sql).Scan(&loanTickets.LoanTickets).Error; err != nil {
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

	query := sq.Select("*").From("loan_tickets")

	if status == "pending" {
		query = query.Where(sq.Eq{"accepted_at": nil})
	} else if status == "accepted" {
		query = query.Where(sq.NotEq{"accepted_at": nil})
	}

	sql, _, _ := query.
		OrderBy("created_at DESC").
		ToSql()

	if err := r.dbClient.Raw(sql).Scan(&loanTickets.LoanTickets).Error; err != nil {
		return nil, err
	}

	return loanTickets, nil
}

func (r *LoanTicketrRepository) GetLoanTicketByIdForAdmin(loanTicketID string) (*model.LoanTicket, error) {
	loatTicket := new(model.LoanTicket)

	if err := r.dbClient.First(loatTicket, loanTicketID).Error; err != nil {
		return nil, err
	}

	return loatTicket, nil
}
