package main

import (
	"github.com/clebersonp/interfaces/without"
)

func main() {
	cars := []without.Car{{Model: "Porsche 911 GT3"}, {Model: "Camaro 2018"}}
	for _, car := range cars {
		car.Accelerate()
	}

	motors := []without.Motorcycle{{Model: "BMW 2023 K 1600 GTL"}, {Model: "KAWASAKI Vulcan 900 Classic"}}
	for _, moto := range motors {
		moto.Accelerate()
	}

}
