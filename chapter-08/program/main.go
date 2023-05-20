package main

import (
	"fmt"
	"github.com/clebersonp/chapter-08/magazine"
)

func main() {
	// Struct literals allow us to use short variable declarations
	adamBean := magazine.Subscriber{
		Name: "Adam Bean",
		Rate: 4.99,
	}
	fmt.Printf("Type: %T, value: %v\n", adamBean, adamBean)
	fmt.Printf("Zero Values for Address: %#v\n", adamBean.Address)

	// Struct literals let you create a struct and set its fields at the same time
	employeeJohn := magazine.Employee{
		Name:   "John",
		Salary: 4000,
		Address: magazine.Address{
			Street:     "Work street",
			City:       "Santi Andreas",
			State:      "San Paulo",
			PostalCode: "08417235",
		},
	}
	fmt.Printf("Type: %T, value: %v\n", employeeJohn, employeeJohn)
	fmt.Printf("Accessing anonymous fields => %v, %v, %v, %v\n",
		employeeJohn.Address.State, employeeJohn.Address.City, employeeJohn.Address.Street, employeeJohn.Address.PostalCode)
	// So now that the Address struct type is embedded within the Employee struct types, you don’t have to write out
	//employeeJohn.Address.City to get at the City field; you can just write employeeJohn.City
	fmt.Printf("Accessing anonymous fields with embedded approach => %v, %v, %v, %v\n",
		employeeJohn.State, employeeJohn.City, employeeJohn.Street, employeeJohn.PostalCode)

	var s interface{} = magazine.Subscriber{
		Name:   "Example interface",
		Rate:   55.10,
		Active: true,
		Address: magazine.Address{
			Street:     "Work street",
			City:       "Santi Andreas",
			State:      "San Paulo",
			PostalCode: "08417235",
		},
	}
	_, isSubscriber := s.(magazine.Subscriber)
	_, isEmployee := s.(magazine.Employee)
	fmt.Printf("Type: %T, value: %v, Is type of magazine.Subscriber? %t, Is type of magazine.Employee? %t\n", s, s, isSubscriber, isEmployee)
}
