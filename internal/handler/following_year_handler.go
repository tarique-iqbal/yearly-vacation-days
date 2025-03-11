package handler

import (
	"fmt"
	"time"
	"yearly-vacation-days/internal/domain"
	"yearly-vacation-days/internal/utils"
)

const (
	thirtyYearsOld         = 30
	additionalVacationDays = 1
)

type FollowingYearHandler struct {
	BaseHandler
}

func (handler *FollowingYearHandler) Calculate(employee domain.Employee, givenYear int, yearlyVacationDays int) (float64, error) {
	layout := "02.01.2006"
	joiningDate, err := time.Parse(layout, employee.ContractStartDate)
	if err != nil {
		return 0, fmt.Errorf("joining-date error: %w", err)
	}
	dateOfBirth, err := time.Parse(layout, employee.DateOfBirth)
	if err != nil {
		return 0, fmt.Errorf("date-of-birth error: %w", err)
	}
	joiningYear := joiningDate.Year()

	if joiningYear < givenYear {
		calculatedVacationDays := float64(yearlyVacationDays)

		if handler.doesAgeCriteriaQualify(dateOfBirth, givenYear) &&
			handler.isEmploymentOnFifthYear(joiningDate, givenYear) {
			calculatedVacationDays += additionalVacationDays
		}

		return calculatedVacationDays, nil
	} else {
		return 0, nil
	}
}

func (handler *FollowingYearHandler) doesAgeCriteriaQualify(dateOfBirth time.Time, givenYear int) bool {
	lastDayOfGivenYear := time.Date(givenYear, 12, 31, 23, 59, 59, 0, time.UTC)

	years, _, _, _ := utils.CalculateYearMonthDayDifference(dateOfBirth, lastDayOfGivenYear)

	return years >= thirtyYearsOld
}

func (handler *FollowingYearHandler) isEmploymentOnFifthYear(contractStartDate time.Time, givenYear int) bool {
	years := givenYear - contractStartDate.Year()

	return years%5 == 0
}
