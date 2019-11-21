package main

import (
	"fmt"
	"math/rand"
)

// save some money to buy a gift for your friend
// write a program that randomly places nickels (0.05$), dimes (0.10$)
// and quarters (0.25$) into an empty piggy bank until it contains at least
// $20.00
// Display the running balance after each deposit - formatted with
// an appriopriate width and precision

func main() {
	var piggyBank float64
	coins := [3]float64{0.05, 0.10, 0.25}
	var coinType string

	for {
		amount := coins[rand.Intn(3)]
		piggyBank += amount

		switch amount {
		case 0.05:
			coinType = "nickel"
		case 0.10:
			coinType = "dime"
		case 0.25:
			coinType = "quarter"
		}

		if piggyBank >= 20.00 {
			fmt.Printf("You have reached your target!\nFinal balance is: $%4.2f\n", piggyBank)
			break
		}

		fmt.Printf("Putting a %v into your bank...\n", coinType)
		fmt.Printf("New balance is: $%4.2f\n\n", piggyBank)
	}
}
