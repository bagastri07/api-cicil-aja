package model

type AmbassadorComissionTrasaction struct {
	ID                       uint64 `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Ammount                  uint64 `json:"ammount"`
	BankAccountInformationID uint64 `json:"bank_account_information_id" gorm:"foreignKey:bank_account_information_id"`
	BaseModel
}
