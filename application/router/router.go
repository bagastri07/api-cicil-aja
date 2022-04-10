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
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.PATCH},
	}))

	// Init Controller
	authCtl := controller.NewAuthController()
	borrowerCtl := controller.NewBorrowerController()
	verificationCtl := controller.NewVerificationController()
	loanTicketCtl := controller.NewLoanTicketController()
	loanBillCtl := controller.NewLoanBillController()

	// Root Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to CicilAja API 😍")
	})

	// Auth Routes
	auth := e.Group("/auth")
	auth.POST("/login", authCtl.BorrowerLogin)
	auth.GET("/oauth-login", authCtl.BorrowerLoginWithGoogle)
	auth.Any("/oauth-callback", authCtl.BorrowerLoginGoogleCallback)

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
	borrower.GET("", borrowerCtl.HandleGetCurrentBorrower)
	borrower.PUT("", borrowerCtl.HandleUpdateBorrower)
	borrower.PUT("/update-bank-information", borrowerCtl.HandleUpdateBorrowerBankAccount)
	borrower.POST("/upload-ktm", borrowerCtl.HandleUploadBorrowerDocument)
	borrower.PATCH("/change-password", borrowerCtl.HandleChangePassword)

	// Group Loan Ticket
	loanTicket := e.Group("/loan-tickets")
	loanTicket.Use(customMiddleware.VerifyToken())

	// Loan Ticket Request
	loanTicket.POST("", loanTicketCtl.HandleMakeLoanTicket)
	loanTicket.GET("", loanTicketCtl.HandleGetAllTicket)
	loanTicket.GET("/:loanTicketID", loanTicketCtl.HandleGetLoanTicketByID)
	loanTicket.DELETE("/:loanTicketID", loanTicketCtl.HandleDeleteLoanTicketByID)

	// Grip Loan Bill
	loanBill := e.Group("/loan-bills")
	loanBill.Use(customMiddleware.VerifyToken())

	// Loan Bills Request
	loanBill.GET("", loanBillCtl.HandleGetAllLoanBill)
	loanBill.PATCH("/:loanBillID", loanBillCtl.HandlePayLoanBillByID)

	// Auth group
	e.POST("/_admin/auth/login", authCtl.AdminLogin)

	// Group fo admin
	admin := e.Group("/_admin")
	admin.Use(customMiddleware.VerifyToken())
	admin.Use(customMiddleware.IsAdmin())

	adminLoanTicket := admin.Group("/loan-tickets")
	adminLoanTicket.GET("", loanTicketCtl.HandleGetAllTicketForAdmin)
	adminLoanTicket.GET("/:loanTicketID", loanTicketCtl.HandleGetLoanTicketByIDForAdmin)
	adminLoanTicket.POST("/:loanTicketID/accept", loanTicketCtl.HandleAcceptLoanTicketByIDForAdmin)

	return e
}
