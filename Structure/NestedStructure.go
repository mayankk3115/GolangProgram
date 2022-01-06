package main

import "fmt"

type Salaried interface {
	getSalary() float64
}

type Salary struct {
	basic     float64
	insurance float64
	hra       float64
}

func (s Salary) getSalary() float64 {
	return s.basic + s.insurance + s.hra
}

type Employee struct {
	firstName, lastName string
	salary              Salaried
}

func main() {

	e := Employee{
		firstName: "Priya",
		lastName:  "Sakrapani",
		salary:    Salary{1200, 14000, 8000},
	}

	fmt.Printf("The Name is %v %v and the salary is %v\n", e.firstName, e.lastName, e.salary.getSalary())
}
