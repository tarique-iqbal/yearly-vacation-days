package handler

import (
	"testing"
	"yearly-vacation-days/internal/domain"
	"yearly-vacation-days/internal/handler"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockNextHandler struct {
	mock.Mock
	next handler.Handler
}

func (m *MockNextHandler) Calculate(employee domain.Employee, givenYear int, yearlyVacationDays int) (float64, error) {
	args := m.Called(employee, givenYear, yearlyVacationDays)
	return args.Get(0).(float64), args.Error(1)
}

func (m *MockNextHandler) SetNext(next handler.Handler) {
	m.next = next
}

func TestJoiningYearHandler_Calculate(t *testing.T) {
	handler := handler.JoiningYearHandler{}

	testCases := []struct {
		name                 string
		contractStartDate    string
		givenYear            int
		yearlyVacationDays   int
		expectedVacationDays float64
		expectError          bool
	}{
		{
			name:                 "Employee joins at start of the year",
			contractStartDate:    "01.01.2022",
			givenYear:            2022,
			yearlyVacationDays:   26,
			expectedVacationDays: 26.0,
		},
		{
			name:                 "Employee joins mid-year",
			contractStartDate:    "01.07.2022",
			givenYear:            2022,
			yearlyVacationDays:   26,
			expectedVacationDays: 13.0,
		},
		{
			name:                 "Employee joins end of year",
			contractStartDate:    "01.12.2022",
			givenYear:            2022,
			yearlyVacationDays:   26,
			expectedVacationDays: 2.16,
		},
		{
			name:                 "Employee joins end of year on 15.12.2022",
			contractStartDate:    "15.12.2022",
			givenYear:            2022,
			yearlyVacationDays:   26,
			expectedVacationDays: 1.08,
		},
		{
			name:                 "Employee joins mid-month (16th)",
			contractStartDate:    "15.06.2022",
			givenYear:            2022,
			yearlyVacationDays:   26,
			expectedVacationDays: 14.08,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			employee := domain.Employee{
				Name:              "Test Employee",
				ContractStartDate: tc.contractStartDate,
			}

			vacationDays, err := handler.Calculate(employee, tc.givenYear, tc.yearlyVacationDays)

			assert.NoError(t, err)
			assert.Equal(t, tc.expectedVacationDays, vacationDays, "Expected %v, got %v", tc.expectedVacationDays, vacationDays)
		})
	}
}

func TestJoiningYearHandler_CalculateExpectError(t *testing.T) {
	handler := handler.JoiningYearHandler{}

	tc := struct {
		name                 string
		contractStartDate    string
		givenYear            int
		yearlyVacationDays   int
		expectedVacationDays float64
	}{
		name:               "Invalid date format",
		contractStartDate:  "invalid-date",
		givenYear:          2022,
		yearlyVacationDays: 26,
	}

	t.Run(tc.name, func(t *testing.T) {
		employee := domain.Employee{
			Name:              "Test Employee",
			ContractStartDate: tc.contractStartDate,
		}

		_, err := handler.Calculate(employee, tc.givenYear, tc.yearlyVacationDays)

		assert.Error(t, err, "Expected error but got none")
		assert.Contains(t, err.Error(), "contract-start-date error")
	})
}

func TestJoiningYearHandler_CallsNextHandler(t *testing.T) {
	mockNextHandler := new(MockNextHandler)
	handler := handler.JoiningYearHandler{}
	handler.SetNext(mockNextHandler)

	employee := domain.Employee{
		Name:              "Test Employee",
		ContractStartDate: "01.01.2020",
	}

	mockNextHandler.On("Calculate", employee, 2023, 26).Return(26.0, nil)

	vacationDays, err := handler.Calculate(employee, 2023, 26)

	assert.NoError(t, err)
	assert.Equal(t, 26.0, vacationDays)

	mockNextHandler.AssertExpectations(t)
}
