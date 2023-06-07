package main

import (
	"fmt"
	"github.com/clebersonp/interfaces/with_pointer_type"
	"reflect"
)

func main() {
	// To use types with pointer types in the implementation interfaces methods we need to use &(address memory pointer)
	// in the instantiation
	// When Go decides whether a value satisfies an interface, pointer methods aren’t included for direct values.
	// But they are included for pointers. So the solution is to assign a pointer to a slice of Vehicle to the
	// vehicles variable, instead of a direct Car or Motorcycle value
	vehicles := []with_pointer_type.Vehicle{
		&with_pointer_type.Car{Model: "Porsche 911 GT3"},
		&with_pointer_type.Car{Model: "Camaro 2018"},
		&with_pointer_type.Motorcycle{Model: "BMW 2023 K 1600 GTL"},
		&with_pointer_type.Motorcycle{Model: "KAWASAKI Vulcan 900 Classic"},
	}
	for _, vehicle := range vehicles {
		vehicle.Brake()
		fmt.Println(vehicle.GetModel())
		// with type assertion is possible to call specifics and concrete methods from the types that implement the interface
		// using pointer type to type assertion when the concrete struct has a pointer receiver on interface implemented method
		if car, ok := vehicle.(*with_pointer_type.Car); ok {
			car.CloseDoor()
		}
		if motorcycle, ok := vehicle.(*with_pointer_type.Motorcycle); ok {
			motorcycle.CloseHelmetVisor()
		}

		// type switch with type assertions
		switch v := vehicle.(type) {
		case *with_pointer_type.Car:
			fmt.Println("Car:", v)
		case *with_pointer_type.Motorcycle:
			fmt.Println("Motorcycle:", v)
		default:
			fmt.Println("Unknown")
		}

		// using Golang reflection
		car := &with_pointer_type.Car{Model: "Beetle 69"}
		var vehicle2 with_pointer_type.Vehicle = car
		fmt.Printf("Type '%v' is equals to type '%v'? %t\n", reflect.TypeOf(car), reflect.TypeOf(vehicle2), reflect.TypeOf(car) == reflect.TypeOf(vehicle2))

	}
}
