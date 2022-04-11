package model

import "time"

type LoanTicket struct {
	ID                     uint64     `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	BorrowerID             uint64     `json:"borrower_id" validate:"required" gorm:"NOT NULL"`
	LoanAmount             float64    `json:"loan_amount" validate:"required,gt=10000" gorm:"NOT NULL"`
	LoanTotal              float64    `json:"loan_total" gorm:"NOT NULL"`
	LoanTenureInMonths     string     `json:"loan_tenure_in_months" gorm:"type:ENUM('3','6','12') NOT NULL"`
	InterestRate           float32    `json:"interest_rate" validate:"required" gorm:"NOT NULL"`
	LoanType               string     `json:"loanType" validate:"required" gorm:"type:ENUM('college-bill', 'shopping') NOT NULL"`
	ItemUrl                string     `json:"item_url" validate:"required,datauri" gorm:"NOT NULL"`
	ReviewedByAmbassadorAt *time.Time `json:"reviewed_by_ambassador_at"`
	AcceptedAt             *time.Time `json:"accepted_at"`
	Status                 string     `json:"status" gorm:"type:ENUM('pending', 'accepted') DEFAULT 'pending'"`
	LoanBills              []LoanBill `json:"loan_bills" gorm:"foreignKey:loan_ticket_id"`
	AmbassadorID           *uint64    `json:"ambassador_id" gorm:"default:null"`
	BaseModel
}

type LoanTickets struct {
	LoanTickets []LoanTicket `json:"loan_tickets"`
}
