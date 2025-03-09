package repository

import (
	"encoding/json"
	"os"
	"testing"
	"yearly-vacation-days/internal/domain"
	"yearly-vacation-days/internal/repository"

	"github.com/stretchr/testify/assert"
)

func TestLoadEmployees_Success(t *testing.T) {
	testData := domain.EmployeeData{
		Employees: map[string]*domain.Employee{
			"1": {
				Name:              "Hans Müller",
				DateOfBirth:       "30.12.1970",
				ContractStartDate: "01.07.2001",
				IsSpecialContract: "no",
			},
		},
	}

	tempFile, err := os.CreateTemp(os.TempDir(), "employees_*.json")
	assert.NoError(t, err)
	defer os.Remove(tempFile.Name())

	jsonData, _ := json.Marshal(testData)
	_, err = tempFile.Write(jsonData)
	assert.NoError(t, err)
	tempFile.Close()

	repo := repository.NewEmployeeRepository()
	employees, err := repo.LoadEmployees(tempFile.Name())

	assert.NoError(t, err)
	assert.Equal(t, testData, employees)
}

func TestLoadEmployees_FileNotFound(t *testing.T) {
	repo := repository.NewEmployeeRepository()
	_, err := repo.LoadEmployees("nonexistent.json")

	assert.Error(t, err)
}

func TestLoadEmployees_InvalidJSON(t *testing.T) {
	tempFile, err := os.CreateTemp(os.TempDir(), "invalid_*.json")
	assert.NoError(t, err)
	defer os.Remove(tempFile.Name())

	_, err = tempFile.Write([]byte("{invalid json}"))
	assert.NoError(t, err)
	tempFile.Close()

	repo := repository.NewEmployeeRepository()
	_, err = repo.LoadEmployees(tempFile.Name())

	assert.Error(t, err)
}
