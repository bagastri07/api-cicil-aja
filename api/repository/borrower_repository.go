package repository

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/bagastri07/api-cicil-aja/api/model"
	"github.com/bagastri07/api-cicil-aja/database"
	"github.com/labstack/echo/v4"
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

func (r *BorrowerRepository) GetBorrowerByEmail(borrowerEmail string) (*model.Borrower, error) {
	borrower := new(model.Borrower)

	err := r.dbClient.Where("email = ?", borrowerEmail).First(&borrower).Error

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

func (r *BorrowerRepository) UpdateBorrower(updatedBorrower *model.UpdateBorrowerPayload, borrowerEmail string) (*model.Borrower, error) {
	borrower := new(model.Borrower)

	if err := r.dbClient.Where("email = ?", borrowerEmail).First(&borrower).Error; err != nil {
		return nil, err
	}

	err := r.dbClient.Model(&borrower).Updates(&model.Borrower{
		Name:          updatedBorrower.Name,
		Birthday:      updatedBorrower.Birthday,
		University:    updatedBorrower.University,
		StudyProgram:  updatedBorrower.StudyProgram,
		StudentNumber: updatedBorrower.StudentNumber,
		PhoneNumber:   updatedBorrower.PhoneNumber,
	}).Error

	if err != nil {
		return nil, err
	}

	return borrower, nil
}

func (r *BorrowerRepository) UpdateBorrowerBankAccount(payload *model.UpdateBorrowerBankAccountPayload, borrowerId uint64) (*model.Borrower, error) {
	borrower := new(model.Borrower)

	err := r.dbClient.Preload("Document").Preload("BankAccountInformation").First(borrower, borrowerId).Error

	if err != nil {
		return nil, err
	}

	if borrower.BankAccountInformation == nil {
		// Create
		r.dbClient.Model(borrower).Association("BankAccountInformation").
			Append(&model.BankAccountInformation{
				BankName:         payload.BankName,
				AccountNumber:    payload.AccountNumber,
				AccountRecipient: payload.AccountRecipient,
			})
	} else {
		// Update
		borrower.BankAccountInformation.AccountNumber = payload.AccountNumber
		borrower.BankAccountInformation.BankName = payload.BankName
		borrower.BankAccountInformation.AccountRecipient = payload.AccountRecipient
	}

	r.dbClient.Save(borrower.BankAccountInformation)

	return borrower, nil
}

func (r *BorrowerRepository) FindForrowerByEmail(borrowerEmail string) (*model.Borrower, error) {
	var borrower model.Borrower

	res := r.dbClient.Preload("Document").Find(&borrower, "email = ?", borrowerEmail)

	if res.Error != nil {
		return nil, res.Error
	}

	fmt.Println(borrower)

	return &borrower, nil
}

func (r *BorrowerRepository) FindBorrowerByID(borrowerId uint64) (*model.Borrower, error) {
	borrower := new(model.Borrower)

	err := r.dbClient.Preload("Document").Preload("BankAccountInformation").First(borrower, borrowerId).Error

	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return borrower, nil
}

func (r *BorrowerRepository) UploadKtmImage(filePath string, borrowerId uint64) (*model.Borrower, error) {
	borrower := new(model.Borrower)

	err := r.dbClient.Preload("Document").Preload("BankAccountInformation").First(borrower, borrowerId).Error

	if err != nil {
		return nil, err
	}

	if borrower.Document == nil {
		r.dbClient.Model(borrower).Association("Document").Append(&model.BorrowerDocument{
			KTMUrl: filePath,
		})
	} else {
		oldPath := borrower.Document.KTMUrl
		os.Remove(filepath.Join("public", oldPath))

		borrower.Document.KTMUrl = filePath
	}
	r.dbClient.Save(borrower.Document)

	return borrower, nil
}

func (r *BorrowerRepository) ChangePassword(borrowerID uint64, newPassword string) error {
	borrower := new(model.Borrower)

	if err := r.dbClient.First(&borrower, borrowerID).Error; err != nil {
		return err
	}

	borrower.Password = newPassword

	r.dbClient.Save(&borrower)

	return nil
}
