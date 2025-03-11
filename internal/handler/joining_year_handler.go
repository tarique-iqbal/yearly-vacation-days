package handler

import (
	"fmt"
	"time"
	"yearly-vacation-days/internal/domain"
	"yearly-vacation-days/internal/utils"
)

const monthsInAYear = 12

type JoiningYearHandler struct {
	BaseHandler
}

func (handler *JoiningYearHandler) Calculate(employee domain.Employee, givenYear int, yearlyVacationDays int) (float64, error) {
	layout := "02.01.2006"
	joiningDate, err := time.Parse(layout, employee.ContractStartDate)
	if err != nil {
		return 0, fmt.Errorf("contract-start-date error: %w", err)
	}

	joiningYear := joiningDate.Year()
	if joiningYear == givenYear {
		firstDayOfNextYear := time.Date(givenYear+1, time.January, 1, 0, 0, 0, 0, time.UTC)

		years, mths, days, _ := utils.CalculateYearMonthDayDifference(joiningDate, firstDayOfNextYear)

		if years == 1 {
			return float64(yearlyVacationDays), nil
		}

		months := float64(mths)
		if days == 17 {
			months += 0.5
		}

		calculatedVacationDays := float64(yearlyVacationDays) / monthsInAYear * months

		return utils.RoundToDecimalPlaces(calculatedVacationDays, 2), nil
	} else {
		return handler.next.Calculate(employee, givenYear, yearlyVacationDays)
	}
}
