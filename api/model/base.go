package model

import "time"

type BaseModel struct {
	CreatedAt *time.Time `json:"created_at" db:"created_at" gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at" gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}

type Login struct {
	Email    string
	Password string
}
