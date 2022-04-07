package model

import "time"

type LoginRequest struct {
	Email    string
	Password string
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

type UpdateBorrowerBankAccount struct {
	BankName         string `json:"bank_name" validate:"required" gorm:"NOT NULL"`
	AccountNumber    string `json:"account_number" validate:"required" gorm:"NOT NULL"`
	AccountRecipient string `json:"account_recipient" validate:"required"`
}
