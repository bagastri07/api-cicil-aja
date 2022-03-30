package model

type BankAccountInformation struct {
	ID                            uint64                          `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	BankName                      string                          `json:"bank_name" gorm:"NOT NULL"`
	AcountNumber                  string                          `json:"account_number" gorm:"NOT NULL"`
	AccountRecipient              string                          `json:"account_recipient"`
	BorrowerID                    uint64                          `json:"borrower_id" gorm:"NOT NULL"`
	Balance                       uint64                          `json:"balance"`
	AmbassadorComissionTrasaction []AmbassadorComissionTrasaction `json:"ambassador_comission_transaction_"`
	BaseModel
}
