package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	kiddo := person{firstName: "Sushma",
		lastName: "Hegde",
		contactInfo: contactInfo{
			email:   "sushma.hegde@gmail.com",
			zipCode: 560067},
	}
	kiddo.updateName("kiddo")
	kiddo.print()
}

func (personPtr *person) updateName(newName string) {
	(*personPtr).firstName = newName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
