package main

type Employee interface {
	CalculateSalary() float64
}

type Fulltime struct {
	Name   string
	Salary float64
}

type Contractor struct {
	Name        string
	Salary      float64 // per hour
	HoursWorked float64
}

func (c Contractor) CalculateSalary() float64 {
	return (c.HoursWorked * c.Salary)
}

func (ft Fulltime) CalculateSalary() float64 {
	return ft.Salary
}
