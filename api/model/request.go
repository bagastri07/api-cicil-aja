package model

import "time"

type LoginRequest struct {
	Email    string
	Password string
}

type UpdateBorrowerPayload struct {
	Name          string    `json:"name" validate:"required"`
	Birthday      time.Time `json:"birthday" validate:"required"`
	University    string    `json:"university" validate:"required"`
	StudyProgram  string    `json:"study_program" validate:"required"`
	StudentNumber string    `json:"student_number" validate:"required"`
	PhoneNumber   string    `json:"phone_number" validate:"required,min=10,max=14,e164"`
	BaseModel
}

type UpdateBorrowerBankAccountPayload struct {
	BankName         string `json:"bank_name" validate:"required" gorm:"NOT NULL"`
	AccountNumber    string `json:"account_number" validate:"required" gorm:"NOT NULL"`
	AccountRecipient string `json:"account_recipient" validate:"required"`
}

type ChangePasswordPayload struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}

type MakeLoanTicketPayload struct {
	LoanAmount         float64 `json:"loan_amount" validate:"required"`
	LoanType           string  `json:"loan_type" validate:"required,oneof=college-bill shopping"`
	LoanTenureInMonths string  `json:"loan_tenure_in_months" validate:"required,oneof=3 6 12"`
	ItemUrl            string  `json:"item_url" validate:"required,url"`
	InterestRate       float32 `json:"interest_rate"`
}
