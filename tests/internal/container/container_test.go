package container_test

import (
	"testing"
	"yearly-vacation-days/internal/container"

	"github.com/stretchr/testify/assert"
)

func TestNewContainer(t *testing.T) {
	c := container.NewContainer()

	assert.NotNil(t, c.EmployeeRepo, "EmployeeRepo should be initialized")
	assert.NotNil(t, c.VacationDaysHandler, "VacationDaysHandler should be initialized")
	assert.NotNil(t, c.VacationDaysService, "VacationDaysService should be initialized")
}
