package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	//alex := person{firstName: "Alex", lastName: "Andersson"}
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Andersson"

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "test@test.123",
			zipCode: 12345,
		},
	}

	fmt.Println(jim)
	fmt.Printf("%+v", jim)
}
