package main

import "fmt"

type City struct {
	belongsTo string
}

type Employee struct {
	firstName, lastName string
	phoneNumber         int
	City
}

func (e *Employee) changeLastName(n string) {
	e.lastName = n
}

func (e Employee) changePhoneNumber() {
	e.phoneNumber = 15029373
	fmt.Println(e.phoneNumber)
}

func (e *Employee) changeCityName(c string) {
	e.City.belongsTo = c
}

func main() {
	e := Employee{
		firstName:   "Mayank",
		lastName:    "Kumar",
		phoneNumber: 7362525,
		City:        City{"Gaya"},
	}
	fmt.Println("The firstName and lastName before change is:", e.firstName, e.lastName, e.City)
	e.changeLastName("chauhan")
	e.changeCityName("Chennai")
	fmt.Println("The firstName and lastName after change is:", e.firstName, e.lastName, e.City)
	e.changePhoneNumber()
	fmt.Println(e.phoneNumber)
}
