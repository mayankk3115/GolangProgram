package main

import "fmt"

type Student struct {
	firstName, lastName string
	admissionNo         int
}

type College struct {
	affliatedTo string
	collegeName string
	Student
}

func main() {

	c := College{
		affliatedTo: "UPTU",
		collegeName: "Galgotias",
		Student:     Student{firstName: "Mayank", lastName: "kumar", admissionNo: 123456},
	}

	fmt.Printf("The name of college is:%v and the name of student is:%v\n", c.collegeName, c.firstName)
}
