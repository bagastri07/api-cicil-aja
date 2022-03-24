package model

import "time"

type Borrower struct {
	ID           uint64    `json:"id"`
	Name         string    `json:"name" validate:"required"`
	Birthday     time.Time `json:"birthday" validate:"required"`
	Email        string    `json:"email" validate:"required,email"`
	Password     string    `json:"password" validate:"required"`
	University   string    `json:"university" validate:"required"`
	StudyProgram string    `json:"study_program" validate:"required"`
	StudentID    string    `json:"student_id" validate:"required"`
	PhoneNumber  string    `json:"phone_number" validate:"required,min=10,max=14,e164"`
	BaseModel
}

type Borrowers struct {
	Borrowers []Borrower `json:"borrowers"`
}
