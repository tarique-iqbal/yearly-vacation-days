package utils

import (
	"testing"
	"time"
	"yearly-vacation-days/internal/utils"

	"github.com/stretchr/testify/assert"
)

func TestCalculateYearMonthDayDifference(t *testing.T) {
	testCases := []struct {
		pastDate       string
		currentDate    string
		expectedYears  int
		expectedMonths int
		expectedDays   int
	}{
		// Simple difference (exact years, months, days)
		{"01.01.2000", "01.01.2020", 20, 0, 0},
		// Leap year test (Feb 29 to Mar 1)
		{"29.02.2000", "01.03.2001", 1, 0, 1},
		// Feb 29 → Feb 28 (1 full year)
		{"29.02.2020", "28.02.2021", 1, 0, 0},
		// Feb 29 → Mar 1 (1 year, 1 day)
		{"29.02.2020", "01.03.2021", 1, 0, 1},
		// End-of-month adjustment (Jan 31 to Feb 28)
		{"31.01.2019", "28.02.2020", 1, 0, 28},
		// Mid-year changes
		{"15.06.2010", "15.09.2015", 5, 3, 0},
		// Negative day case (adjusting previous month)
		{"15.05.2018", "10.06.2020", 2, 0, 26},
	}

	layout := "02.01.2006"

	for _, tc := range testCases {
		past, _ := time.Parse(layout, tc.pastDate)
		current, _ := time.Parse(layout, tc.currentDate)

		years, months, days, err := utils.CalculateYearMonthDayDifference(past, current)

		assert.NoError(t, err)
		assert.Equal(t, tc.expectedYears, years, "Years mismatch for %v : %v", tc.pastDate, tc.currentDate)
		assert.Equal(t, tc.expectedMonths, months, "Months mismatch for %v : %v", tc.pastDate, tc.currentDate)
		assert.Equal(t, tc.expectedDays, days, "Days mismatch for %v : %v", tc.pastDate, tc.currentDate)
	}
}

func TestRoundToDecimalPlaces(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		places   int
		expected float64
	}{
		{"Round to 2 decimal places", 12.3456, 2, 12.34},
		{"Round to 1 decimal place", 7.89, 1, 7.8},
		{"Round to 0 decimal places", 9.99, 0, 9.0},
		{"Round negative number", -5.6789, 3, -5.678},
		{"Round very small number", 0.000456, 4, 0.0004},
		{"No rounding needed", 100.0, 2, 100.0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := utils.RoundToDecimalPlaces(tc.value, tc.places)
			assert.Equal(t, tc.expected, result)
		})
	}
}
