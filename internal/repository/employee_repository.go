package repository

import (
	"encoding/json"
	"os"
	"yearly-vacation-days/internal/domain"
)

type EmployeeRepository interface {
	LoadEmployees(filename string) (domain.EmployeeData, error)
}

type EmployeeRepositoryImplement struct{}

func NewEmployeeRepository() EmployeeRepository {
	return &EmployeeRepositoryImplement{}
}

func (r *EmployeeRepositoryImplement) LoadEmployees(filename string) (domain.EmployeeData, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return domain.EmployeeData{}, err
	}

	var data domain.EmployeeData
	err = json.Unmarshal(file, &data)
	if err != nil {
		return domain.EmployeeData{}, err
	}

	return data, nil
}
