package model

type AmbassadorRegistration struct {
	ID         uint64 `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	BorrowerID uint64 `json:"borrower_id" validate:"required" gorm:"NOT NULL"`
	Status     string `json:"status" gorm:"type:ENUM('pending','accepted','rejected') DEFAULT 'pending'"`
	BaseModel
}

type AmbassadorRegistrations struct {
	AmbassadorRegistrations []AmbassadorRegistration `json:"ambassador_registrations"`
}
