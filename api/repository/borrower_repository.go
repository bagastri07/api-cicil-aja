package repository

import (
	"errors"

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
	borrower := new(model.Borrower)

	err := r.dbClient.Where("id = ?", borrowerID).First(&borrower).Error

	if err != nil {
		return nil, err
	}

	return borrower, nil
}

func (r *BorrowerRepository) CreateBorrower(borrower *model.Borrower) error {
	if r.dbClient.Where("email = ?", borrower.Email).Find(&borrower).RowsAffected > 0 {
		return errors.New("borrower with this email is already exist")
	}

	if err := r.dbClient.Create(&borrower).Error; err != nil {
		return err
	}
	return nil
}
