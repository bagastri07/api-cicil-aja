package model

import "time"

type LoanBill struct {
	ID              uint64     `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	BorrowerID      uint64     `json:"borrower_id" validate:"required" gorm:"NOT NULL"`
	LoanTicketID    uint64     `json:"load_ticket" validate:"required" gorm:"NOT NULL"`
	PaymentDeadline time.Time  `jsonL:"payment_deadline" validate:"required" gorm:"NOT NULL"`
	PaidAt          *time.Time `json:"paid_at" validate:"required"`
	BaseModel
}
