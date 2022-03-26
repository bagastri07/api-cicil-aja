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
}
