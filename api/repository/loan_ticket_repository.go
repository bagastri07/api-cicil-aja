package repository

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/bagastri07/api-cicil-aja/api/model"
	"github.com/bagastri07/api-cicil-aja/database"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type LoanTicketRepository struct {
	dbClient *gorm.DB
}

func NewLoanTicketRepository() *LoanTicketRepository {
	dbClient := database.GetDBConnection()
	return &LoanTicketRepository{
		dbClient: dbClient,
	}
}

func (r *LoanTicketRepository) MakeNewLoanTicket(borrowerID uint64, payload *model.MakeLoanTicketPayload) (*model.LoanTicket, error) {
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
		BorrowerID:         borrowerID,
		LoanTotal:          payload.LoanAmount + (payload.LoanAmount * float64(payload.InterestRate)),
		Status:             "pending",
	}

	r.dbClient.Create(&loanTicket)

	return &loanTicket, nil
}

func (r *LoanTicketRepository) GetAllLoanTickets(borrowerID uint64, statuses string) (*model.LoanTickets, error) {
	loanTickets := new(model.LoanTickets)

	query := r.dbClient

	statusesSlice := strings.Split(statuses, ",")
	if statuses != "" {
		for index, status := range statusesSlice {
			if index == 0 {
				query = query.Where("status = ?", status)
			} else {
				query = query.Or("status = ?", status)
			}
		}
	}

	if err := query.Order("created_at desc").Find(&loanTickets.LoanTickets, "borrower_id", borrowerID).Error; err != nil {
		return nil, err
	}

	return loanTickets, nil
}

func (r *LoanTicketRepository) GetLoanTicketById(borrowerID uint64, loanTicketID string) (*model.LoanTicket, error) {
	loatTicket := new(model.LoanTicket)

	if err := r.dbClient.Where("borrower_id", borrowerID).Preload("LoanBills").First(loatTicket, loanTicketID).Error; err != nil {
		return nil, err
	}

	return loatTicket, nil
}

func (r *LoanTicketRepository) DeleteLoanTicketById(borrowerID uint64, loanTicketID string) (*model.LoanTicket, error) {
	loatTicket := new(model.LoanTicket)

	if err := r.dbClient.Where("borrower_id", borrowerID).First(loatTicket, loanTicketID).Error; err != nil {
		return nil, err
	}

	r.dbClient.Delete(loatTicket)

	return loatTicket, nil
}

// ============ Ambassador Repository ==============

func (r *LoanTicketRepository) ReviewLoanTikcetByAmbassador(ambassadorID uint64, loanTicketID string) (*model.LoanTicket, error) {
	loatTicket := new(model.LoanTicket)

	if err := r.dbClient.Where("ambassador_id = ?", ambassadorID).Where("ID = ?", loanTicketID).First(&loatTicket).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if loatTicket.ReviewedByAmbassadorAt != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, &model.MessageResponse{
			Message: "this loan ticket is already reviewed",
		})
	}

	now := time.Now()
	loatTicket.ReviewedByAmbassadorAt = &now

	if err := r.dbClient.Save(&loatTicket).Error; err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return loatTicket, nil
}

func (r *LoanTicketRepository) GetAllLoanTicketsForAmbassador(ambassadorID uint64, statuses string) (*model.LoanTickets, error) {
	loanTickets := new(model.LoanTickets)

	query := r.dbClient

	statusesSlice := strings.Split(statuses, ",")
	if statuses != "" {
		for index, status := range statusesSlice {
			if index == 0 {
				query = query.Where("status = ?", status)
			} else {
				query = query.Or("status = ?", status)
			}
		}
	}

	if err := query.Order("created_at desc").Find(&loanTickets.LoanTickets, "ambassador_id = ?", ambassadorID).Error; err != nil {
		return nil, err
	}

	return loanTickets, nil
}

// ============ Admin Repository ==============

func (r *LoanTicketRepository) GetAllLoanTicketsForAdmin(statuses string) (*model.LoanTickets, error) {
	loanTickets := new(model.LoanTickets)

	query := r.dbClient

	statusesSlice := strings.Split(statuses, ",")

	if statuses != "" {
		for index, status := range statusesSlice {
			if index == 0 {
				query = query.Where("status = ?", status)
			} else {
				query = query.Or("status = ?", status)
			}
		}
	}

	if err := query.Order("created_at desc").Find(&loanTickets.LoanTickets).Error; err != nil {
		return nil, err
	}

	return loanTickets, nil
}

func (r *LoanTicketRepository) GetLoanTicketByIdForAdmin(loanTicketID string) (*model.LoanTicket, error) {
	loatTicket := new(model.LoanTicket)

	result := r.dbClient.Preload("LoanBills").First(loatTicket, loanTicketID)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return loatTicket, nil
}

func (r *LoanTicketRepository) UpdateStatusLoanTicketByIDForAdmin(loanTicketID, status string) (*model.LoanTicket, error) {
	loanTicket, err := r.GetLoanTicketByIdForAdmin(loanTicketID)

	if err != nil {
		return nil, err
	}

	if (loanTicket.Status != "pending" && loanTicket.Status != "amba-ready") || loanTicket.Status == status {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, &model.MessageResponse{
			Message: "this loan ticket is already " + loanTicket.Status,
		})
	}

	if status == "amba-ready" {
		ambaSelectedID, err := r.AssignAmbassadorToTicket(loanTicket.BorrowerID)

		if err != nil {
			return nil, err
		}

		loanTicket.AmbassadorID = ambaSelectedID
	} else if status == "accepted" {
		if loanTicket.Status != "amba-ready" {
			return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, &model.MessageResponse{
				Message: fmt.Sprintf("this loan ticket is %s, can not be accepted", loanTicket.Status),
			})
		}

		if loanTicket.ReviewedByAmbassadorAt == nil {
			return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, &model.MessageResponse{
				Message: "this loan ticket need to be reviewed by ambassador",
			})
		}

		acceptedTime := time.Now()
		loanTicket.AcceptedAt = &acceptedTime

		loanTicketRepo := NewLoanBillRepository()
		if err := loanTicketRepo.MakeAllBillsByLoanTicketIDForAdmin(loanTicket); err != nil {
			return nil, err
		}

	}

	loanTicket.Status = status

	if err := r.dbClient.Save(&loanTicket).Error; err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return loanTicket, nil
}

func (r *LoanTicketRepository) AssignAmbassadorToTicket(borrowerID uint64) (*uint64, error) {
	ambassadorRepo := NewAmbassadorReposotory()

	result, err := ambassadorRepo.GetAllAmbassadorsWithTheNumberOfTicket()

	ambaTickets := result.AmbassadarAndTickets

	ambaSelected := ambaTickets[0]

	if borrowerID == ambaSelected.ID {
		ambaSelected = ambaTickets[1]
	}

	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return &ambaSelected.ID, nil
}
