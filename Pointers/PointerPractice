package main

import "fmt"

type employeeDetails struct {
	firstName, lastName string
	password            string
}

func changePassword(e *string) string {
	pass := "1234567890"
	*e = pass
	return *e
}

func main() {

	e := employeeDetails{
		firstName: "Mayank",
		lastName:  "priya",
		password:  "pass3115",
	}
	fmt.Printf("I am %v%v and my password is %v\n", e.firstName, e.lastName, e.password)
	changePassword(&e.password)
	fmt.Printf("Now I am %v%v and my password is %v\n", e.firstName, e.lastName, e.password)

	// a := 100
	// var b *int
	// b= &a
	// fmt.Printf("My type is %T and my address is %v and my value is %v\n", b, b, *b)
	// *b = 1000
	// fmt.Printf("Now My type is %T and my address is %v and my value is %v\n", b, b, *b)
}
