package model

import (
	"time"
)

type VerificationToken struct {
	ID        uint64    `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Email     string    `json:"email" gorm:"not null"`
	Token     string    `json:"token" gorm:"not null"`
	ExpiredAt time.Time `json:"expired_at" gorm:"not null"`
}
