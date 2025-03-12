package container

import (
	"yearly-vacation-days/internal/handler"
	"yearly-vacation-days/internal/repository"
	"yearly-vacation-days/internal/service"
)

type Container struct {
	EmployeeRepo        repository.EmployeeRepository
	VacationDaysHandler handler.Handler
	VacationDaysService *service.VacationDaysService
}

func NewContainer() *Container {
	employeeRepo := repository.NewEmployeeRepository()

	var joiningYearHandler handler.Handler = &handler.JoiningYearHandler{}
	var followingYearHandler handler.Handler = &handler.FollowingYearHandler{}

	joiningYearHandler.SetNext(followingYearHandler)

	vacationService := service.NewVacationDaysService(employeeRepo, joiningYearHandler)

	return &Container{
		EmployeeRepo:        employeeRepo,
		VacationDaysHandler: joiningYearHandler,
		VacationDaysService: vacationService,
	}
}
