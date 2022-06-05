// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	model "github.com/bagastri07/api-cicil-aja/api/model"
	mock "github.com/stretchr/testify/mock"
)

// LoanTicketRepository is an autogenerated mock type for the LoanTicketRepository type
type LoanTicketRepository struct {
	mock.Mock
}

// DeleteLoanTicketById provides a mock function with given fields: borrowerID, loanTicketID
func (_m *LoanTicketRepository) DeleteLoanTicketById(borrowerID uint64, loanTicketID string) (*model.LoanTicket, error) {
	ret := _m.Called(borrowerID, loanTicketID)

	var r0 *model.LoanTicket
	if rf, ok := ret.Get(0).(func(uint64, string) *model.LoanTicket); ok {
		r0 = rf(borrowerID, loanTicketID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.LoanTicket)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64, string) error); ok {
		r1 = rf(borrowerID, loanTicketID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllLoanTickets provides a mock function with given fields: borrowerID, statuses
func (_m *LoanTicketRepository) GetAllLoanTickets(borrowerID uint64, statuses string) (*model.LoanTickets, error) {
	ret := _m.Called(borrowerID, statuses)

	var r0 *model.LoanTickets
	if rf, ok := ret.Get(0).(func(uint64, string) *model.LoanTickets); ok {
		r0 = rf(borrowerID, statuses)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.LoanTickets)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64, string) error); ok {
		r1 = rf(borrowerID, statuses)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllLoanTicketsForAdmin provides a mock function with given fields: statuses
func (_m *LoanTicketRepository) GetAllLoanTicketsForAdmin(statuses string) (*model.LoanTickets, error) {
	ret := _m.Called(statuses)

	var r0 *model.LoanTickets
	if rf, ok := ret.Get(0).(func(string) *model.LoanTickets); ok {
		r0 = rf(statuses)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.LoanTickets)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(statuses)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllLoanTicketsForAmbassador provides a mock function with given fields: ambassadorID, statuses
func (_m *LoanTicketRepository) GetAllLoanTicketsForAmbassador(ambassadorID uint64, statuses string) (*model.LoanTickets, error) {
	ret := _m.Called(ambassadorID, statuses)

	var r0 *model.LoanTickets
	if rf, ok := ret.Get(0).(func(uint64, string) *model.LoanTickets); ok {
		r0 = rf(ambassadorID, statuses)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.LoanTickets)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64, string) error); ok {
		r1 = rf(ambassadorID, statuses)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLoanLoanTicketByIdForAmbassador provides a mock function with given fields: ambassadorID, loanTicketID
func (_m *LoanTicketRepository) GetLoanLoanTicketByIdForAmbassador(ambassadorID uint64, loanTicketID string) (*model.LoanTicketAndBorrower, error) {
	ret := _m.Called(ambassadorID, loanTicketID)

	var r0 *model.LoanTicketAndBorrower
	if rf, ok := ret.Get(0).(func(uint64, string) *model.LoanTicketAndBorrower); ok {
		r0 = rf(ambassadorID, loanTicketID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.LoanTicketAndBorrower)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64, string) error); ok {
		r1 = rf(ambassadorID, loanTicketID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLoanTicketById provides a mock function with given fields: borrowerID, loanTicketID
func (_m *LoanTicketRepository) GetLoanTicketById(borrowerID uint64, loanTicketID string) (*model.LoanTicket, error) {
	ret := _m.Called(borrowerID, loanTicketID)

	var r0 *model.LoanTicket
	if rf, ok := ret.Get(0).(func(uint64, string) *model.LoanTicket); ok {
		r0 = rf(borrowerID, loanTicketID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.LoanTicket)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64, string) error); ok {
		r1 = rf(borrowerID, loanTicketID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLoanTicketByIdForAdmin provides a mock function with given fields: loanTicketID
func (_m *LoanTicketRepository) GetLoanTicketByIdForAdmin(loanTicketID string) (*model.LoanTicket, error) {
	ret := _m.Called(loanTicketID)

	var r0 *model.LoanTicket
	if rf, ok := ret.Get(0).(func(string) *model.LoanTicket); ok {
		r0 = rf(loanTicketID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.LoanTicket)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(loanTicketID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MakeNewLoanTicket provides a mock function with given fields: borrowerID, payload
func (_m *LoanTicketRepository) MakeNewLoanTicket(borrowerID uint64, payload *model.MakeLoanTicketPayload) (*model.LoanTicket, error) {
	ret := _m.Called(borrowerID, payload)

	var r0 *model.LoanTicket
	if rf, ok := ret.Get(0).(func(uint64, *model.MakeLoanTicketPayload) *model.LoanTicket); ok {
		r0 = rf(borrowerID, payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.LoanTicket)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64, *model.MakeLoanTicketPayload) error); ok {
		r1 = rf(borrowerID, payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReviewLoanTikcetByAmbassador provides a mock function with given fields: ambassadorID, loanTicketID
func (_m *LoanTicketRepository) ReviewLoanTikcetByAmbassador(ambassadorID uint64, loanTicketID string) (*model.LoanTicket, error) {
	ret := _m.Called(ambassadorID, loanTicketID)

	var r0 *model.LoanTicket
	if rf, ok := ret.Get(0).(func(uint64, string) *model.LoanTicket); ok {
		r0 = rf(ambassadorID, loanTicketID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.LoanTicket)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64, string) error); ok {
		r1 = rf(ambassadorID, loanTicketID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStatusLoanTicketByIDForAdmin provides a mock function with given fields: loanTicketID, status
func (_m *LoanTicketRepository) UpdateStatusLoanTicketByIDForAdmin(loanTicketID string, status string) (*model.LoanTicket, error) {
	ret := _m.Called(loanTicketID, status)

	var r0 *model.LoanTicket
	if rf, ok := ret.Get(0).(func(string, string) *model.LoanTicket); ok {
		r0 = rf(loanTicketID, status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.LoanTicket)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(loanTicketID, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
