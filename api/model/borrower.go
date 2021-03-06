package model

import (
	"time"
)

type Borrower struct {
	ID                            uint64                          `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Name                          string                          `json:"name" validate:"required" gorm:"not null"`
	Birthday                      time.Time                       `json:"birthday" validate:"required" gorm:"not null"`
	Email                         string                          `json:"email" validate:"required,email" gorm:"not null"`
	Password                      string                          `json:"password" validate:"required" gorm:"not null"`
	University                    string                          `json:"university" validate:"required" gorm:"not null"`
	StudyProgram                  string                          `json:"study_program" validate:"required" gorm:"not null"`
	StudentNumber                 string                          `json:"student_number" validate:"required" gorm:"not null"`
	PhoneNumber                   string                          `json:"phone_number" validate:"required,min=10,max=14,e164" gorm:"not null"`
	Document                      *BorrowerDocument               `json:"borrower_document"`
	VerifiedAt                    *time.Time                      `json:"verified_at" gorm:"type:DATETIME DEFAULT NULL"`
	AmbassadorRegistration        *AmbassadorRegistration         `json:"ambassador_registration,omitempty" gorm:"foreignKey:borrower_id"`
	IsAmbassador                  bool                            `json:"is_ambassador" gorm:"type:tinyint(1) DEFAULT 0 NOT NULL"`
	AcceptedAsAmbassadorAt        *time.Time                      `json:"accepted_as_ambassador_at" gorm:"default:null"`
	AmbassadorLoanTickets         []LoanTicket                    `json:"ambassador_loan_tickets,omitempty" gorm:"foreignKey:ambassador_id"`
	AmbassadorComissionTrasaction []AmbassadorComissionTrasaction `json:"ambassador_comission_transaction,omitempty" gorm:"foreignKey:ambassador_id"`
	BankAccountInformation        *BankAccountInformation         `json:"bank_information" gorm:"foreignKey:borrower_id"`
	LoanTickets                   *[]LoanTicket                   `json:"loan_tickets,omitempty" gorm:"foreignKey:borrower_id"`
	LoanBills                     *[]LoanBill                     `json:"loan_bills,omitempty" gorm:"foreignKey:borrower_id"`
	BaseModel
}

type BorrowerDocument struct {
	ID         uint64 `json:"id"`
	BorrowerID uint64 `json:"borrower_id"`
	KTMUrl     string `json:"ktm_url"`
	KTPUrl     string `json:"ktp_url"`
	BaseModel
}

type Borrowers struct {
	Borrowers []Borrower `json:"borrowers"`
}

type AmbassadorWithTheNumberOfTicket struct {
	ID             uint64
	AcceptedAt     *time.Time
	NumberOfTicket int64
}

type AmbassadorWithTheNumberOfTickets struct {
	AmbassadarAndTickets []AmbassadorWithTheNumberOfTicket
}
