package model

type AmbassadorComissionTrasaction struct {
	ID           uint64  `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Amount       float64 `json:"amount"`
	Type         string  `json:"type" gorm:"type:ENUM('in', 'out') NOT NULL"`
	AmbassadorID uint64  `json:"ambassador_id"`
	BaseModel
}

type AmbassadorBalanceDetail struct {
	Debit   float64 `json:"debit"`
	Credit  float64 `json:"credit"`
	Balance float64 `json:"balance"`
}
