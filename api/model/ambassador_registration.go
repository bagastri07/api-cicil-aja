package model

import "time"

type AmbassadorRegistration struct {
	ID         uint64 `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	BorrowerID uint64 `json:"borrower_id" validate:"required" gorm:"NOT NULL"`
	Status     string `json:"status" gorm:"type:ENUM('pending','accepted','rejected') DEFAULT 'pending'"`
	BaseModel
}

type AmbassadorRegistrations struct {
	AmbassadorRegistrations []AmbassadorRegistration `json:"ambassador_registrations"`
}

type GetAllAmbassadorRegistrations struct {
	ID            uint64    `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	BorrowerID    uint64    `json:"borrower_id" validate:"required" gorm:"NOT NULL"`
	Status        string    `json:"status" gorm:"type:ENUM('pending','accepted','rejected') DEFAULT 'pending'"`
	Name          string    `json:"name" gorm:"not null"`
	Birthday      time.Time `json:"birthday" gorm:"not null"`
	Email         string    `json:"email" gorm:"not null"`
	University    string    `json:"university" gorm:"not null"`
	StudyProgram  string    `json:"study_program" gorm:"not null"`
	StudentNumber string    `json:"student_number" gorm:"not null"`
	PhoneNumber   string    `json:"phone_number" gorm:"not null"`
}
