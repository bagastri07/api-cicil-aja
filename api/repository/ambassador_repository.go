package repository

import (
	"errors"
	"net/http"
	"strings"

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

func (r *AmbassadorRepository) GetAllAmbassadorRegistrationForAdmin(statuses string) (*model.AmbassadorRegistrations, error) {
	registrations := new(model.AmbassadorRegistrations)

	query := r.dbClient

	statusesSlice := strings.Split(statuses, ",")

	for index, status := range statusesSlice {
		if index == 0 {
			query = query.Where("status = ?", status)
		} else {
			query = query.Or("status = ?", status)
		}
	}

	if err := query.Find(&registrations.AmbassadorRegistrations).Error; err != nil {
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
