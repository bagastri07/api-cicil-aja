package repository

import (
	"net/http"

	"github.com/bagastri07/api-cicil-aja/api/model"
	"github.com/bagastri07/api-cicil-aja/database"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ComissionRepository struct {
	dbClient *gorm.DB
}

func NewComissionTransactionRepository() *ComissionRepository {
	dbClient := database.GetDBConnection()
	return &ComissionRepository{
		dbClient: dbClient,
	}
}

func (r *ComissionRepository) CreateComissionTrasaction(AmbassadorID uint64, amount float64, transactionType string) error {
	comissionTransaction := &model.AmbassadorComissionTrasaction{
		Amount:       amount,
		AmbassadorID: AmbassadorID,
		Type:         transactionType,
	}

	if err := r.dbClient.Create(comissionTransaction).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (r *ComissionRepository) GetComissionBalanceAmbassador(ambassadorID uint64) (*model.AmbassadorBalanceDetail, error) {
	balanceDetail := new(model.AmbassadorBalanceDetail)
	comissionTransaction := new(model.AmbassadorComissionTrasaction)

	if err := r.dbClient.Model(comissionTransaction).
		Select("COALESCE(SUM(amount), 0) as 'in'").
		Where("ambassador_id = ?", ambassadorID).
		Where("type = ?", "in").
		Scan(&balanceDetail.In).Error; err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := r.dbClient.Model(comissionTransaction).
		Select("COALESCE(SUM(amount), 0) as 'out'").
		Where("ambassador_id = ?", ambassadorID).
		Where("type = ?", "out").
		Scan(&balanceDetail.Out).Error; err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	balanceDetail.Balance = balanceDetail.In - balanceDetail.Out

	return balanceDetail, nil
}
