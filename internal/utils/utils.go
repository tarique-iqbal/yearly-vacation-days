package utils

import (
	"math"
	"time"
)

func CalculateYearMonthDayDifference(pastTime time.Time, currentTime time.Time) (int, int, int, error) {
	years := currentTime.Year() - pastTime.Year()
	months := int(currentTime.Month()) - int(pastTime.Month())
	days := currentTime.Day() - pastTime.Day()

	if days < 0 {
		prevMonth := currentTime.AddDate(0, -1, 0)
		daysInPrevMonth := time.Date(prevMonth.Year(), prevMonth.Month()+1, 0, 0, 0, 0, 0, time.UTC).Day()
		days += daysInPrevMonth
		months--

		if pastTime.Month() == 2 && pastTime.Day() == 29 && currentTime.Month() == 3 {
			days += 1
		}
	}

	if months < 0 {
		months += 12
		years--
	}

	if pastTime.Month() == 2 && pastTime.Day() == 29 && currentTime.Month() == 2 && currentTime.Day() == 28 {
		days -= 30
		months -= 11
		years++
	}

	return years, months, days, nil
}

func RoundToDecimalPlaces(value float64, places int) float64 {
	multiplier := math.Pow(10, float64(places))

	return math.Trunc(value*multiplier) / multiplier
}
