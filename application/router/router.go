package router

import (
	"net/http"

	"github.com/bagastri07/api-cicil-aja/api/controller"
	"github.com/bagastri07/api-cicil-aja/api/repository"
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

	//init repo
	loanTicketRepo := repository.NewLoanTicketRepository()

	// Init Controller
	authCtl := controller.NewAuthController()
	borrowerCtl := controller.NewBorrowerController()
	verificationCtl := controller.NewVerificationController()
	loanTicketCtl := controller.NewLoanTicketController(e, loanTicketRepo)
	loanBillCtl := controller.NewLoanBillController()
	ambassadorCtl := controller.NewAmbassadorController()
	commissionCtl := controller.NewCommissionTransactionController()

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
	borrower.GET("", borrowerCtl.HandleGetCurrentBorrower)
	borrower.PUT("", borrowerCtl.HandleUpdateBorrower)
	borrower.PUT("/update-bank-information", borrowerCtl.HandleUpdateBorrowerBankAccount)
	borrower.POST("/upload-ktm", borrowerCtl.HandleUploadKTMBorrowerDocument)
	borrower.POST("/upload-ktp", borrowerCtl.HandleUploadKTPBorrowerDocument)
	borrower.PATCH("/change-password", borrowerCtl.HandleChangePassword)

	// Group Loan Ticket
	loanTicket := e.Group("/loan-tickets")
	loanTicket.Use(customMiddleware.VerifyToken())

	// Loan Ticket Request
	loanTicket.POST("", loanTicketCtl.HandleMakeLoanTicket)
	loanTicket.GET("", loanTicketCtl.HandleGetAllTicket)
	loanTicket.GET("/:loanTicketID", loanTicketCtl.HandleGetLoanTicketByID)
	loanTicket.DELETE("/:loanTicketID", loanTicketCtl.HandleDeleteLoanTicketByID)
	loanTicket.POST("/calculate-estimation", loanTicketCtl.HandleCalculateEstimateLoanTicket)

	// Group Loan Bill
	loanBill := e.Group("/loan-bills")
	loanBill.Use(customMiddleware.VerifyToken())

	// Loan Bills Request
	loanBill.GET("", loanBillCtl.HandleGetAllLoanBill)
	loanBill.PATCH("/:loanBillID", loanBillCtl.HandlePayLoanBillByID)

	//A Group Ambassador
	ambasaddor := e.Group("/ambassadors")
	ambasaddor.Use(customMiddleware.VerifyToken())

	// Ambassador Request
	ambasaddor.POST("/register", ambassadorCtl.HandleRegisterAsAmbassador)

	// Auth group
	e.POST("/_admin/auth/login", authCtl.AdminLogin)

	// Group fo adminEndpoint
	adminEndpoint := e.Group("/_admin")
	adminEndpoint.Use(customMiddleware.VerifyToken())
	adminEndpoint.Use(customMiddleware.IsAdmin())

	adminBorrower := adminEndpoint.Group("/borrowers")
	adminBorrower.GET("", borrowerCtl.HandleGetAllBorrowersForAdmin)
	adminBorrower.GET("/:borrowerID", borrowerCtl.HandleGetBorrowerByIDForAdmin)

	adminLoanTicket := adminEndpoint.Group("/loan-tickets")
	adminLoanTicket.GET("", loanTicketCtl.HandleGetAllTicketForAdmin)
	adminLoanTicket.GET("/:loanTicketID", loanTicketCtl.HandleGetLoanTicketByIDForAdmin)
	adminLoanTicket.PATCH("/:loanTicketID/update-status", loanTicketCtl.HandleUpdateStatusLoanTicketByIDForAdmin)

	adminAmbassador := adminEndpoint.Group("/ambassadors")
	adminAmbassador.GET("", ambassadorCtl.HandleGetAcceptedAmbassadorForAdmin)
	adminAmbassador.GET("/registrations", ambassadorCtl.HandleGetAllAmbassadorRegistrationsForAdmin)
	adminAmbassador.GET("/number-of-tickets", ambassadorCtl.HandleGetAllAmbassadorsAndNumberOfTickets)
	adminAmbassador.PATCH("/:registrationID/update-status", ambassadorCtl.HandleUpdateRegistrationStatusForAdmin)

	// group for ambassador
	ambasaddorEndpoint := e.Group("/_ambassador")
	ambasaddorEndpoint.Use(customMiddleware.VerifyToken())
	ambasaddorEndpoint.Use(customMiddleware.IsAmbassador())

	// group for loan tickets
	ambasaddorLoanTicket := ambasaddorEndpoint.Group("/loan-tickets")
	ambasaddorLoanTicket.GET("", loanTicketCtl.HandleGetAllLoanTicketForAmbassador)
	ambasaddorLoanTicket.GET("/:loanTicketID", loanTicketCtl.HandleGetLoanTicketByIDForAmbassador)
	ambasaddorLoanTicket.PATCH("/:loanTicketID/reviewed", loanTicketCtl.HandleReviewLoanTicketByAmbassador)

	// Group For Comission Transaction
	ambasaddorComission := ambasaddorEndpoint.Group("/commissions")
	ambasaddorComission.GET("/balance-detail", commissionCtl.HandleGetBalanceDetailAmbassador)
	ambasaddorComission.GET("/comission-history", commissionCtl.HandleGetAllComissionHistory)

	// Group for Borrower
	ambassadorBorrower := ambasaddorEndpoint.Group("/borrowers")
	ambassadorBorrower.GET("/:borrowerID", borrowerCtl.HandleGetBorrowerByIDForAdmin)

	return e
}
