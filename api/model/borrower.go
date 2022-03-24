package model

import "time"

type Borrower struct {
	ID           uint64    	`json:"id" validate:"required"`
	Name         string 	`json:"name" validate:"required"`
	Birthday     time.Time  `json:"birthday" validate:"required,datetime"`
	Email        string 	`json:"email" validate:"required,email"`
	Password     string 	`json:"password" validate:"required"`
	University   string 	`json:"university" validate:"required"`
	StudyProgram string 	`json:"study_program" validate:"required"`
	StudyID      string 	`json:"study_id" validate:"required"`
	PhoneCode    string 	`json:"phone_code" validate:"required,min=2,max=3,numeric"`
	PhoneNumber  string 	`json:"phone_number" validate:"required,min=5,max=12,numeric"`
}

type Borrowers struct {
	Borrowers []Borrower `json:"borrowers"`
}