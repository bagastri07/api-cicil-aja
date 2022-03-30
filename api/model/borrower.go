package model

import (
	"time"
)

type Borrower struct {
	ID                     uint64                  `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Name                   string                  `json:"name" validate:"required" gorm:"not null"`
	Birthday               time.Time               `json:"birthday" validate:"required" gorm:"not null"`
	Email                  string                  `json:"email" validate:"required,email" gorm:"not null"`
	Password               string                  `json:"password" validate:"required" gorm:"not null"`
	University             string                  `json:"university" validate:"required" gorm:"not null"`
	StudyProgram           string                  `json:"study_program" validate:"required" gorm:"not null"`
	StudentNumber          string                  `json:"student_number" validate:"required" gorm:"not null"`
	PhoneNumber            string                  `json:"phone_number" validate:"required,min=10,max=14,e164" gorm:"not null"`
	AmbassadorRegistration *AmbassadorRegistration `json:"ambassador_registration" gorm:"foreignKey:borrower_id"`
	ApprovedAsAmbassadorAt *time.Time              `json:"approved_as_ambassador_at"`
	AmbassadorLoanTickets  []*LoanTicket           `json:"ambassador_loan_tickets" gorm:"foreignKey:ambassador_id"`
	BorrowerDocument       BorrowerDocument        `json:"borrower_document" gorm:"foreignKey:borrower_id"`
	BankAccountInformation BankAccountInformation  `json:"bank_information" gorm:"foreignKey:borrower_id"`
	LoanTickets            []LoanTicket            `json:"loan_tickets" gorm:"foreignKey:borrower_id"`
	LoanBills              []LoanBill              `json:"loan_bills" gorm:"foreignKey:borrower_id"`
	VerifiedAt             *time.Time              `json:"verified_at" gorm:"type:DATETIME DEFAULT NULL"`
	BaseModel
}

type BorrowerDocument struct {
	ID         uint64 `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	BorrowerID uint64 `json:"borrower_id" gorm:"NOT NULL"`
	KTPUrl     string `json:"ktp_url" gorm:"NOT NULL"`
	KTMUrl     string `json:"ktm_url" gorm:"NOT NULL"`
	BaseModel
}

type UpdateBorrower struct {
	Name          string    `json:"name" validate:"required"`
	Birthday      time.Time `json:"birthday" validate:"required"`
	University    string    `json:"university" validate:"required"`
	StudyProgram  string    `json:"study_program" validate:"required"`
	StudentNumber string    `json:"student_number" validate:"required"`
	PhoneNumber   string    `json:"phone_number" validate:"required,min=10,max=14,e164"`
	BaseModel
}

type Borrowers struct {
	Borrowers []Borrower `json:"borrowers"`
}
