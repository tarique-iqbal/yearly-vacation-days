package main

import (
	"fmt"
	"log"
	"yearly-vacation-days/internal/container"
	"yearly-vacation-days/internal/utils"
)

func main() {
	year, err := utils.GetYearFromCLI()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	iocContainer := container.NewContainer()

	errorsList, employees := iocContainer.VacationDaysService.CalculateAll(year)
	if len(errorsList) > 0 {
		for _, err := range errorsList {
			fmt.Println("Vacation days error: ", err)
		}
	}

	genericMap := make(map[string]any)
	for key, value := range employees.Employees {
		genericMap[key] = value
	}
	sortedIDs := utils.GetSortedEmployeeIDs(genericMap)

	fmt.Println("Vacation Days Report:")
	for _, id := range sortedIDs {
		employee := employees.Employees[id]

		var vacationDays string

		if employee.VacationDays == 0.0 {
			vacationDays = "Not applicable"
		} else if employee.VacationDays == float64(int(employee.VacationDays)) {
			vacationDays = fmt.Sprintf("%.0f", employee.VacationDays)
		} else {
			vacationDays = fmt.Sprintf("%.2f", employee.VacationDays)
		}

		fmt.Printf("%s: %s\n", employee.Name, vacationDays)
	}
}
