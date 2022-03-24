package repository

import (
	"github.com/bagastri07/api-cicil-aja/api/model"
	"github.com/bagastri07/api-cicil-aja/database"
	"gorm.io/gorm"
)

type BorrowerRepository struct {
	dbClient *gorm.DB
}

func NewBorrowerRepository() *BorrowerRepository {
	dbClient := database.GetDBConnection()
	return &BorrowerRepository{
		dbClient: dbClient,
	}
}

func (r *BorrowerRepository) GetBorrowerByID(borrowerID uint64) (*model.Borrower, error) {
	var borrower model.Borrower

	err := r.dbClient.Where("id = ?", borrowerID).First(&borrower).Error

	if err != nil {
		return nil, err
	}

	return &borrower, nil
}