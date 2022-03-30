package model

type Admin struct {
	ID       uint64 `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Username string `json:"username" validate:"required" gorm:"not null"`
	Pasword  string `json:"password" validate:"required" gorm:"not null"`
	BaseModel
}
