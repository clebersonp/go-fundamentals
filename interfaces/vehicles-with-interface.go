package main

import (
	"github.com/clebersonp/interfaces/with"
)

func main() {
	vehicles := []with.Vehicle{
		with.Car{Model: "Porsche 911 GT3"},
		with.Car{Model: "Camaro 2018"},
		with.Motorcycle{Model: "BMW 2023 K 1600 GTL"},
		with.Motorcycle{Model: "KAWASAKI Vulcan 900 Classic"},
	}
	for _, vehicle := range vehicles {
		vehicle.Brake()
	}
}
