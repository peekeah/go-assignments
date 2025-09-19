package main

import "fmt"

func RunEmployeePayroll() {
	type RawEmployeeData struct {
		Name        string
		Role        string
		HoursWorked int
		Salary      int
	}

	intialEmp := []RawEmployeeData{
		{Name: "Alice", Role: "FullTime", HoursWorked: 0, Salary: 60000},
		{Name: "Bob", Role: "Contractor", HoursWorked: 160, Salary: 500},
		{Name: "Charlie", Role: "FullTime", HoursWorked: 0, Salary: 45000},
		{Name: "Diana", Role: "Contractor", HoursWorked: 100, Salary: 800},
		{Name: "Ethan", Role: "FullTime", HoursWorked: 0, Salary: 75000},
		{Name: "Fiona", Role: "Contractor", HoursWorked: 120, Salary: 600},
		{Name: "George", Role: "FullTime", HoursWorked: 0, Salary: 55000},
		{Name: "Hannah", Role: "Contractor", HoursWorked: 140, Salary: 700},
		{Name: "Ian", Role: "FullTime", HoursWorked: 0, Salary: 50000},
		{Name: "Julia", Role: "Contractor", HoursWorked: 180, Salary: 650},
	}

	employee := []Employee{}

	for _, val := range intialEmp {
		if val.Role == "Contractor" {
			employee = append(employee, Contractor{
				val.Name,
				float64(val.HoursWorked),
				float64(val.Salary),
			})
		} else {
			employee = append(employee, Fulltime{
				val.Name,
				float64(val.Salary),
			})
		}
	}

	emp1 := employee[0]

	salary1 := emp1.CalculateSalary()
	fmt.Printf("salary of emp1 is %.2f\n", salary1)
}
