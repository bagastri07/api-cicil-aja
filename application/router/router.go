package router

import (
	"net/http"

	"github.com/bagastri07/api-cicil-aja/api/controller"
	customMiddleware "github.com/bagastri07/api-cicil-aja/application/middleware"
	"github.com/bagastri07/api-cicil-aja/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	//Use Validator
	validator.Init(e)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// Init Controller
	authCtl := controller.NewAuthController()
	borrowerCtl := controller.NewBorrowerController()
	verificationCtl := controller.NewVerificationController()

	// Root Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to CicilAja API 😍")
	})

	// Auth Routes
	auth := e.Group("/auth")
	auth.POST("/login", authCtl.BorrowerLogin)

	//Group route for verification
	verification := e.Group("/verifications")
	verification.POST("/send-email", verificationCtl.HandleSendEmailVerification, customMiddleware.VerifyToken())
	verification.GET("/verify-borrower/:email/:verificationToken", verificationCtl.HandleVerifyBorrower)

	//Grup route for borrower
	borrower := e.Group("/borrowers")
	borrower.POST("/create", borrowerCtl.HandleCreateNewBorrower)
	// Make other borrower endpoints restrict
	borrower.Use(customMiddleware.VerifyToken())
	borrower.Use(customMiddleware.CheckVerificationStatus())
	borrower.GET("/", borrowerCtl.HandleGetBorrowerByEmail)
	borrower.PUT("/update/:borrowerID", borrowerCtl.HandleUpdateBorrower)

	return e
}
