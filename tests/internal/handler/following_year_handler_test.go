package handler

import (
	"testing"
	"yearly-vacation-days/internal/domain"
	"yearly-vacation-days/internal/handler"

	"github.com/stretchr/testify/assert"
)

const (
	givenYear              = 2024
	ordinaryVacationDays   = 26
	specialVacationDays    = 27
	additionalVacationDays = 1
)

func TestCalculate_OrdinaryVacationDays(t *testing.T) {
	handler := handler.FollowingYearHandler{}

	employee := domain.Employee{
		Name:              "Hans Müller",
		DateOfBirth:       "30.12.1970",
		ContractStartDate: "01.07.2001",
		IsSpecialContract: "no",
	}

	expectedVacationDays := float64(ordinaryVacationDays)

	vacationDays, err := handler.Calculate(employee, givenYear, ordinaryVacationDays)

	assert.NoError(t, err)
	assert.Equal(t, expectedVacationDays, vacationDays)
}

func TestCalculate_SpecialVacationDays(t *testing.T) {
	handler := handler.FollowingYearHandler{}

	employee := domain.Employee{
		Name:              "John Doe",
		DateOfBirth:       "15.05.1980", // Will be over 30 in 2024
		ContractStartDate: "10.08.2019", // 5th work year in 2024
		IsSpecialContract: "yes",
	}

	expectedVacationDays := float64(specialVacationDays + additionalVacationDays)

	vacationDays, err := handler.Calculate(employee, givenYear, specialVacationDays)

	assert.NoError(t, err)
	assert.Equal(t, expectedVacationDays, vacationDays)
}

func TestFollowingYearHandler_Calculate_ErrorCases(t *testing.T) {
	handler := handler.FollowingYearHandler{}

	tests := []struct {
		name        string
		employee    domain.Employee
		givenYear   int
		expectedErr string
	}{
		{
			name: "Invalid contract start date",
			employee: domain.Employee{
				Name:              "John Doe",
				ContractStartDate: "invalid-date",
				DateOfBirth:       "30.12.1970",
			},
			givenYear:   2023,
			expectedErr: "joining-date error",
		},
		{
			name: "Invalid date of birth",
			employee: domain.Employee{
				Name:              "Jane Doe",
				ContractStartDate: "01.07.2001",
				DateOfBirth:       "invalid-date",
			},
			givenYear:   2023,
			expectedErr: "date-of-birth error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := handler.Calculate(tt.employee, tt.givenYear, 26)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), tt.expectedErr)
		})
	}
}
