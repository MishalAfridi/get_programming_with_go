package main

import (
	"fmt"
)

func main() {
	const hoursPerDay = 24

	var days = 28
	var distance = 56000000 // km
	// how fast would a ship need to travel in km/h in order to reach Malacandra/Mars in 28 days
	var speed = distance / (days * hoursPerDay)

	fmt.Println(speed, "km/h")

}