package main

import (
	"os"

	"github.com/bagastri07/api-cicil-aja/application/router"
	"github.com/common-nighthawk/go-figure"
	"github.com/joho/godotenv"
)

func main() {
	figure.NewColorFigure("CicilAja API", "", "purple", true).Print()

	// load env
	godotenv.Load(".env")

	// Init router
	e := router.Init()

	// APP_PORT
	appPort := os.Getenv("APP_PORT")

	if appPort == "" {
		appPort = "7500"
	}

	// Start App
	e.Logger.Fatal(e.Start(":" + appPort))
}
