package service

import (
	"fmt"
	"yearly-vacation-days/internal/domain"
	"yearly-vacation-days/internal/handler"
	"yearly-vacation-days/internal/repository"
)

const (
	ordinaryVacationDays = 26
	specialVacationDays  = 27
	employeeDataFilePath = "data/employees.json"
)

type VacationDaysService struct {
	employeeRepository repository.EmployeeRepository
	handler            handler.Handler
}

func NewVacationDaysService(repo repository.EmployeeRepository, handler handler.Handler) *VacationDaysService {
	return &VacationDaysService{
		employeeRepository: repo,
		handler:            handler,
	}
}

func (vds *VacationDaysService) CalculateAll(year int) ([]error, domain.EmployeeData) {
	var yearlyVacationDays int

	employees, err := vds.employeeRepository.LoadEmployees(employeeDataFilePath)
	if err != nil {
		panic("Error loading employees: " + err.Error())
	}

	var errorsList []error

	for _, employee := range employees.Employees {
		if employee.IsSpecialContract == "yes" {
			yearlyVacationDays = specialVacationDays
		} else {
			yearlyVacationDays = ordinaryVacationDays
		}

		vacationDays, err := vds.handler.Calculate(*employee, year, yearlyVacationDays)
		if err != nil {
			errorsList = append(errorsList, fmt.Errorf("employee %s: %w", employee.Name, err))
			continue
		}

		employee.VacationDays = vacationDays
	}

	return errorsList, employees
}
