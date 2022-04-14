package model

type DataResponse struct {
	Data interface{} `json:"data"`
}

type MessageResponse struct {
	Message interface{} `json:"messages"`
}

type MessageDataResponse struct {
	Message interface{} `json:"messages"`
	Data    interface{} `json:"data"`
}

type CalculateEstimateLoanTicketResponse struct {
	LoanTotal   float64 `json:"loan_total"`
	Interest    float64 `json:"interest"`
	MonthlyBill float64 `json:"monthly_bill"`
}

type OauthResp struct {
}
