package main

import "fmt"

func fuelGauge(fuel int) {
	fmt.Println("The fuel check copy ", fuel)
}

func calculateFuel(planet string) int {
	var fuel int
	switch planet {
	case "Venus":
		fuel = 300000
	case "Mercury":
		fuel = 500000
	case "Mars":
		fuel = 700000
	case "Jupiter":
		fuel = 400000
	case "Saturn":
		fuel = 9000000
	case "Pluto":
		fuel = 10000000
	case "Uranus":
		fuel = 110000000
	case "Neptune":
		fuel = 1300000000
	default:
		fuel = 0

	}
	return fuel
}

func greetPlanet(planet string) string {
	fmt.Println("Welcome to the new world we call ", planet)
	return planet
}

func cantFly() {
	fmt.Println("We do not have the available fuel to fly there.")
}

func flyToPlanet(planet string, fuel int) int {
	var fuelRemaining, fuelCost int
	fuelRemaining = fuel
	fuelCost = calculateFuel(planet)
	if fuelRemaining >= fuelCost {
		greetPlanet(planet)
		fuelCost -= fuelRemaining
	} else {
		cantFly()
	}

	return fuelRemaining
}
func fuelGauge1(refuel int) {
	fmt.Println("The fuel check copy ", refuel)
}

func toHome(refuel int) {
	fmt.Println("Now about Home")
	if refuel >= 30000 {
		fmt.Println("We are coming Home ! oh hooo")
	} else {
		fmt.Println("Last ride to Humaninty :) ")
	}
}

func main() {

	var fuel int
	fuel = 500000

	fuelGauge(fuel)
	fmt.Println("Where you wanna fly MAVERICK! ?")
	var planet string
	fmt.Scan(&planet)
	fuel = flyToPlanet(planet, fuel)

	refuel := fuel - calculateFuel(planet)
	fuelGauge1(refuel)
	toHome(refuel)

}
