package repository

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/bagastri07/api-cicil-aja/api/model"
	"github.com/bagastri07/api-cicil-aja/database"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AmbassadorRepository struct {
	dbClient *gorm.DB
}

func NewAmbassadorReposotory() *AmbassadorRepository {
	return &AmbassadorRepository{
		dbClient: database.GetDBConnection(),
	}
}

func (r *AmbassadorRepository) RegisterAsAmbassador(borrowerID uint64) (*model.AmbassadorRegistration, error) {
	registration := new(model.AmbassadorRegistration)

	err := r.dbClient.Last(&registration, "borrower_id = ?", borrowerID).Error

	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, echo.NewHTTPError(http.StatusInternalServerError)
		}
	}

	if errors.Is(err, gorm.ErrRecordNotFound) || registration.Status == "rejected" {
		newRegistration := &model.AmbassadorRegistration{
			Status:     "pending",
			BorrowerID: borrowerID,
		}

		if err := r.dbClient.Create(&newRegistration).Error; err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError)
		}
		return newRegistration, nil
	}

	if registration.Status == "pending" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, &model.MessageResponse{
			Message: "this borrower still have pending registration",
		})
	}

	return nil, echo.NewHTTPError(http.StatusBadRequest, &model.MessageResponse{
		Message: "this borrower is already accepted as ambassador",
	})
}

func (r *AmbassadorRepository) GetAllAmbassadorsWithTheNumberOfTicket() (*model.AmbassadorWithTheNumberOfTickets, error) {
	ambassadorLoanTicket := new(model.AmbassadorWithTheNumberOfTickets)

	if err := r.dbClient.Raw(`	SELECT  b.id, COUNT(lt.ambassador_id) AS number_of_ticket,
									b.accepted_as_ambassador_at AS accepted_at
								from cicil_aja.borrowers b 
								LEFT JOIN cicil_aja.loan_tickets lt 
								on b.id  = lt.ambassador_id AND
									lt.status <> 'rejected'
								WHERE b.is_ambassador = 1 AND lt.deleted_at IS NULL
								GROUP  BY  b.id
								ORDER BY number_of_ticket ASC, accepted_at ASC`).
		Scan(&ambassadorLoanTicket.AmbassadarAndTickets).Error; err != nil {
		return nil, err
	}

	return ambassadorLoanTicket, nil
}

// ============ Admin Repository ==============

func (r *AmbassadorRepository) UpdateAmbassadorRegistrationStatusForAdmin(registrationID string, payload *model.UpdateRegistrationStatus) error {
	registration := new(model.AmbassadorRegistration)

	if err := r.dbClient.Where("ID = ?", registrationID).First(&registration).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if registration.Status == "rejected" || registration.Status == "accepted" {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, &model.MessageResponse{
			Message: "cannot update, status is already accepted or rejected",
		})
	}

	if payload.Status == "accepted" {
		borrower := new(model.Borrower)
		if err := r.dbClient.First(&borrower, registration.BorrowerID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return echo.NewHTTPError(http.StatusNotFound, err.Error())
			}
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		now := time.Now()
		borrower.AcceptedAsAmbassadorAt = &now
		borrower.IsAmbassador = true

		if err := r.dbClient.Save(&borrower).Error; err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

	}

	registration.Status = payload.Status

	if err := r.dbClient.Save(&registration).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return nil
}

func (r *AmbassadorRepository) GetAllAmbassadorRegistrationForAdmin(statuses string) ([]model.GetAllAmbassadorRegistrations, error) {
	registrations := []model.GetAllAmbassadorRegistrations{}

	query := r.dbClient.Table("ambassador_registrations").
		Select(`ambassador_registrations.id, 
				ambassador_registrations.borrower_id,
				ambassador_registrations.status,
				borrowers.name,
				borrowers.birthday,
				borrowers.email,
				borrowers.university,
				borrowers.study_program,
				borrowers.phone_number,
				borrowers.student_number`)

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

	if err := query.
		Joins("LEFT JOIN borrowers ON borrowers.id = ambassador_registrations.borrower_id").
		Scan(&registrations).Error; err != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return registrations, nil
}

func (r *AmbassadorRepository) GetAllAcceptedAmbassadorForAdmin() (*model.Borrowers, error) {
	borrowers := new(model.Borrowers)

	if err := r.dbClient.Where("is_ambassador = ?", true).Find(&borrowers.Borrowers).Error; err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return borrowers, nil
}
