package service_test

import (
	"errors"
	"fmt"
	"testing"
	"yearly-vacation-days/internal/domain"
	"yearly-vacation-days/internal/handler"
	"yearly-vacation-days/internal/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	ordinaryVacationDays = 26
	specialVacationDays  = 27
	employeeDataFilePath = "data/employees.json"
)

type MockEmployeeRepository struct {
	mock.Mock
}

func (m *MockEmployeeRepository) LoadEmployees(filename string) (domain.EmployeeData, error) {
	args := m.Called(filename)
	return args.Get(0).(domain.EmployeeData), args.Error(1)
}

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Calculate(employee domain.Employee, givenYear int, yearlyVacationDays int) (float64, error) {
	args := m.Called(employee, givenYear, yearlyVacationDays)
	return args.Get(0).(float64), args.Error(1)
}

func (m *MockHandler) SetNext(handler handler.Handler) {
	m.Called(handler)
}

func TestCalculateAll_Success(t *testing.T) {
	mockRepo := new(MockEmployeeRepository)
	mockHandler := new(MockHandler)

	mockEmployees := domain.EmployeeData{
		Employees: map[string]*domain.Employee{
			"1": {
				Name:              "Hans Müller",
				DateOfBirth:       "30.12.1970",
				ContractStartDate: "01.07.2001",
				IsSpecialContract: "no",
			},
			"2": {
				Name:              "John Doe",
				DateOfBirth:       "15.05.1980",
				ContractStartDate: "10.08.2005",
				IsSpecialContract: "yes",
			},
		},
	}

	mockRepo.On("LoadEmployees", employeeDataFilePath).Return(mockEmployees, nil)

	mockHandler.On("Calculate", *mockEmployees.Employees["1"], 2024, ordinaryVacationDays).Return(26.0, nil)
	mockHandler.On("Calculate", *mockEmployees.Employees["2"], 2024, specialVacationDays).Return(27.0, nil)

	vds := service.NewVacationDaysService(mockRepo, mockHandler)

	_, employees := vds.CalculateAll(2024)

	assert.Equal(t, 26.0, employees.Employees["1"].VacationDays)
	assert.Equal(t, 27.0, employees.Employees["2"].VacationDays)

	mockRepo.AssertExpectations(t)
	mockHandler.AssertExpectations(t)
}

func TestCalculateAll_EmployeeLoadFailure(t *testing.T) {
	mockRepo := new(MockEmployeeRepository)
	mockHandler := new(MockHandler)

	mockRepo.On("LoadEmployees", employeeDataFilePath).Return(domain.EmployeeData{}, errors.New("file not found"))

	vds := service.NewVacationDaysService(mockRepo, mockHandler)

	assert.Panics(t, func() {
		vds.CalculateAll(2024)
	})

	defer func() {
		if r := recover(); r != nil {
			assert.Equal(t, "Error loading employees: file not found", r)
		}
	}()

	vds.CalculateAll(2024)

	mockRepo.AssertExpectations(t)
}

func TestCalculateAll_ErrorHandling(t *testing.T) {
	mockRepo := new(MockEmployeeRepository)
	mockHandler := new(MockHandler)

	mockEmployees := domain.EmployeeData{
		Employees: map[string]*domain.Employee{
			"1": {
				Name:              "Alice Doe",
				ContractStartDate: "01.01.2020",
				DateOfBirth:       "15.05.1990",
				IsSpecialContract: "no",
			},
			"2": {
				Name:              "Bob Smith",
				ContractStartDate: "03.03.2015",
				DateOfBirth:       "20.07.1985",
				IsSpecialContract: "yes",
			},
		},
	}

	mockRepo.On("LoadEmployees", employeeDataFilePath).Return(mockEmployees, nil)

	mockHandler.On("Calculate", mock.Anything, mock.Anything, mock.Anything).Return(0.0, errors.New("calculation error"))

	vds := service.NewVacationDaysService(mockRepo, mockHandler)

	errorsList, _ := vds.CalculateAll(2024)

	assert.Len(t, errorsList, 2, "There should be errors for both employees")

	expectedError1 := fmt.Sprintf("employee %s: %s", mockEmployees.Employees["1"].Name, "calculation error")
	expectedError2 := fmt.Sprintf("employee %s: %s", mockEmployees.Employees["2"].Name, "calculation error")

	assert.EqualError(t, errorsList[0], expectedError1)
	assert.EqualError(t, errorsList[1], expectedError2)

	mockRepo.AssertExpectations(t)
	mockHandler.AssertExpectations(t)
}
