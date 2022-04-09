package model

import "time"

type LoanBill struct {
	ID              uint64     `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	BorrowerID      uint64     `json:"borrower_id" validate:"required" gorm:"NOT NULL"`
	LoanTicketID    uint64     `json:"load_ticket_id" validate:"required" gorm:"NOT NULL"`
	BillAmount      float64    `json:"bill_amount" validate:"required" gorm:"NOT NULL"`
	PaymentDeadline time.Time  `jsonL:"payment_deadline" validate:"required" gorm:"NOT NULL"`
	PaidAt          *time.Time `json:"paid_at" validate:"required"`
	Status          string     `json:"status" gorm:"type:ENUM('paid', 'unpaid') DEFAULT 'unpaid'"`
	BaseModel
}

type LoanBills struct {
	LoanBills []LoanBill `json:"loan_bills"`
}
