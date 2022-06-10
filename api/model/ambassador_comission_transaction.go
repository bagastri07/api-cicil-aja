package model

type AmbassadorComissionTrasaction struct {
	ID           uint64  `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Amount       float64 `json:"amount"`
	Type         string  `json:"type" gorm:"type:ENUM('in', 'out') NOT NULL"`
	AmbassadorID uint64  `json:"ambassador_id"`
	BaseModel
}

type AmbassadorBalanceDetail struct {
	In      float64 `json:"in"`
	Out     float64 `json:"out"`
	Balance float64 `json:"balance"`
}

type WitdhdrawPayload struct {
	Ammount float64 `json:"ammount" validate:"required,number"`
}
