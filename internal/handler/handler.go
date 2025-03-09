package handler

import "yearly-vacation-days/internal/domain"

type Handler interface {
	SetNext(handler Handler)
	Calculate(employee domain.Employee, givenYear int, yearlyVacationDays int) (float64, error)
}
