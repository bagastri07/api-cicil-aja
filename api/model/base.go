package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	CreatedAt *time.Time     `json:"created_at" db:"created_at" gorm:"type:DATETIME DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time     `json:"updated_at" db:"updated_at" gorm:"type:DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" db:"deleted_at"`
}
