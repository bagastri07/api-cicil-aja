package main

import (
	"github.com/bagastri07/api-cicil-aja/api/model"
	"github.com/bagastri07/api-cicil-aja/database"
	"github.com/joho/godotenv"
)

func main() {
	// load env
	godotenv.Load(".env")
	db := database.GetDBConnection()

	db.AutoMigrate(&model.Borrower{})
	db.AutoMigrate(&model.BorrowerDocument{})
	db.AutoMigrate(&model.VerificationToken{})
	db.AutoMigrate(&model.LoanTicket{})
	db.AutoMigrate(&model.LoanBill{})
	db.AutoMigrate(&model.AmbassadorRegistration{})
	db.AutoMigrate(&model.BankAccountInformation{})
	db.AutoMigrate(&model.Admin{})
	db.AutoMigrate(&model.AmbassadorComissionTrasaction{})

}
