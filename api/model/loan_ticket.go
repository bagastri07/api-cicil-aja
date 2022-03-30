package model

import "time"

type LoanTicket struct {
	ID                     uint64     `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	BorrowerID             uint64     `json:"borrower_id" validate:"required" gorm:"NOT NULL"`
	LoanAmount             uint64     `json:"loan_amount" validate:"required,gt=10000" gorm:"NOT NULL"`
	LoanTenureInMonths     string     `json:"loan_tenure_in_months" gorm:"type:ENUM('3','6','12') NOT NULL"`
	InterestRate           float32    `json:"interest_rate" validate:"required" gorm:"NOT NULL"`
	LoanType               string     `json:"loanType" validate:"required" gorm:"type:ENUM('college tuition', 'online shopping') NOT NULL"`
	ItemUrl                string     `json:"item_url" validate:"required,datauri" gorm:"NOT NULL"`
	ReviewedByAmbassadorAt *time.Time `json:"reviewed_by_ambassador_at"`
	AcceptedAt             *time.Time `json:"accepted_at"`
	LoanBills              []LoanBill `json:"loand_bills" gorm:"foreignKey:loan_ticket_id"`
	AmbassadorID           uint64     `json:"ambassador_id" gorm:"NOT NULL"`
	BaseModel
}
