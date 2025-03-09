package domain

type Employee struct {
	Name              string  `json:"name"`
	DateOfBirth       string  `json:"dateOfBirth"`
	ContractStartDate string  `json:"contractStartDate"`
	IsSpecialContract string  `json:"isSpecialContract"`
	VacationDays      float64 `json:"-"`
}

type EmployeeData struct {
	Employees map[string]*Employee `json:"employees"`
}
