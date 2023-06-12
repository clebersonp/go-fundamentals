package main

import (
	"encoding/json"
	"fmt"
	"github.com/clebersonp/go-enum/car"
	"log"
)

func main() {
	bmw := car.New(car.BMW, car.White, "320i GT")
	ferrari := car.New(car.Ferrari, car.Red, "SF90 Stadala")

	fmt.Printf("I own a '%s' and dream about a '%s'\n", bmw, ferrari)

	fmt.Printf("==============JSON Marshal==================\n\n")
	cars := []*car.Car{bmw, ferrari}
	carsJson, err := json.MarshalIndent(cars, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(carsJson))

	fmt.Printf("=================JSON Unmarshal=============\n\n")
	porscheJson := []byte(`{"brand":"Porsche", "color":"White", "model":"Taycan"}`)
	var porsche car.Car
	err = json.Unmarshal(porscheJson, &porsche)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Another Car I love: %v\n", porsche)
}
