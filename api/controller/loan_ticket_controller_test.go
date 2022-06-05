package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/bagastri07/api-cicil-aja/api/model"
	"github.com/bagastri07/api-cicil-aja/api/repository"
	mocksLoanTicketRepo "github.com/bagastri07/api-cicil-aja/mocks/repository"
	"github.com/bagastri07/api-cicil-aja/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestNewLoanTicketController(t *testing.T) {
	e := echo.New()
	loanTicketRepo := new(mocksLoanTicketRepo.LoanTicketRepository)

	type args struct {
		e    *echo.Echo
		repo repository.LoanTicketRepository
	}

	tests := []struct {
		name string
		args args
		want *LoanTicketController
	}{
		{
			name: "success",
			args: args{
				e:    e,
				repo: loanTicketRepo,
			},
			want: &LoanTicketController{
				e:                    e,
				loanTicketRepository: loanTicketRepo,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLoanTicketController(tt.args.e, tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLoanTicketController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoanBillController_HandleCalculateEstimateLoanTicket(t *testing.T) {
	type args struct {
		payload string
	}
	tests := []struct {
		name       string
		args       args
		path       string
		wantErr    bool
		wantStatus int
		wantBody   string
	}{
		{
			name: "Success, when Loan Tenure is 3",
			args: args{
				payload: `{
					"loan_amount": 50000,
					"loan_tenure_in_months": "3"
				}`,
			},
			path:       "/loan-tickets/calculate-estimation",
			wantErr:    false,
			wantStatus: 200,
			wantBody: `{
				"data": {
					"loan_total": 60000,
					"interest": 10000,
					"monthly_bill": 20000
				}
			}`,
		},
		{
			name: "Success, when Loan Tenure is 6",
			args: args{
				payload: `{
					"loan_amount": 50000,
					"loan_tenure_in_months": "6"
				}`,
			},
			path:       "/loan-tickets/calculate-estimation",
			wantErr:    false,
			wantStatus: 200,
			wantBody: `{
				"data": {
					"loan_total": 62500,
					"interest": 12500,
					"monthly_bill": 10416.666666666666
				}
			}`,
		},
		{
			name: "Success, when Loan Tenure is 12",
			args: args{
				payload: `{
					"loan_amount": 50000,
					"loan_tenure_in_months": "12"
				}`,
			},
			path:       "/loan-tickets/calculate-estimation",
			wantErr:    false,
			wantStatus: 200,
			wantBody: `{
				"data": {
					"loan_total": 67500,
					"interest": 17500,
					"monthly_bill": 5625
				}
			}`,
		},
		{
			name: "Error Bad Request",
			args: args{
				payload: `{
					"loan_amount": 50000,
					"loan_tenure_in_months": "13"
				}`,
			},
			path:       "/loan-tickets/calculate-estimation",
			wantErr:    true,
			wantStatus: 400,
			wantBody: `{
				"message": {
					"errors": [
						{
							"Field": "LoanTenureInMonths",
							"Msg": "This field should one of [3 6 12]"
						}
					]
				}
			}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			payload := new(model.CalculateEstimateLoanTicketPayload)
			json.Unmarshal([]byte(tt.args.payload), payload)

			req := httptest.NewRequest(http.MethodPost, tt.path, strings.NewReader(tt.args.payload))
			rec := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, rec)

			validator.Init(e)

			req.Header.Set("Content-Type", "application/json")

			ctl := &LoanTicketController{
				e: e,
			}

			err := ctl.HandleCalculateEstimateLoanTicket(c)
			assert.Equal(t, tt.wantErr, err != nil)
			if err == nil {
				assert.JSONEq(t, tt.wantBody, rec.Body.String())
				assert.Equal(t, tt.wantStatus, rec.Code)
			}
		})
	}
}
