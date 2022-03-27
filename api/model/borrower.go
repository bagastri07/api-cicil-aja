package model

import "time"

type Borrower struct {
	ID            uint64     `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Name          string     `json:"name" validate:"required" gorm:"not null"`
	Birthday      time.Time  `json:"birthday" validate:"required" gorm:"not null"`
	Email         string     `json:"email" validate:"required,email" gorm:"not null"`
	Password      string     `json:"password" validate:"required" gorm:"not null"`
	University    string     `json:"university" validate:"required" gorm:"not null"`
	StudyProgram  string     `json:"study_program" validate:"required" gorm:"not null"`
	StudentNumber string     `json:"student_number" validate:"required" gorm:"not null"`
	PhoneNumber   string     `json:"phone_number" validate:"required,min=10,max=14,e164" gorm:"not null"`
	VerifiedAt    *time.Time `json:"verified_at" gorm:"type:DATETIME DEFAULT NULL"`
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
