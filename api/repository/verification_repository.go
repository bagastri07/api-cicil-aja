package repository

import (
	"errors"
	"time"

	"github.com/bagastri07/api-cicil-aja/api/model"
	"github.com/bagastri07/api-cicil-aja/database"
	"gorm.io/gorm"
)

type VerificationRepository struct {
	dbClient *gorm.DB
}

func NewVerificationRepository() *VerificationRepository {
	dbClient := database.GetDBConnection()
	return &VerificationRepository{
		dbClient: dbClient,
	}
}

func (r *VerificationRepository) CreateNewVerification(email, verificationToken string) error {
	verification := &model.VerificationToken{
		Email:     email,
		Token:     verificationToken,
		ExpiredAt: time.Now().Add(time.Second * 30),
	}

	if err := r.dbClient.Create(&verification).Error; err != nil {
		return err
	}

	return nil
}

func (r *VerificationRepository) VerifyBorrower(email, token string) error {
	verification := new(model.VerificationToken)

	if err := r.dbClient.Where("email = ?", email).Last(&verification).Error; err != nil {
		return err
	}

	if verification.Used {
		return errors.New("verification link is used ")
	}

	if !verification.ExpiredAt.After(time.Now()) {
		return errors.New("verification link expired")
	}

	if verification.Email == email && verification.Token == token {
		verification.Used = true
		r.dbClient.Save(&verification)

		borrower := new(model.Borrower)
		if err := r.dbClient.Where("email = ?", email).First(&borrower).Error; err != nil {
			return err
		}

		now := time.Now()
		if err := r.dbClient.Model(&borrower).Updates(&model.Borrower{
			VerifiedAt: &now,
		}).Error; err != nil {
			return err
		}
	} else {
		return errors.New("token or email not match")
	}

	return nil

}
