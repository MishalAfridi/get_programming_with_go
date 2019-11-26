package main

import (
	"fmt"
	"math/rand"
)

func main() {
	piggyBank := 0
	coins := [3]int{5, 10, 25}

	for {
		amount := coins[rand.Intn(3)]
		piggyBank += amount
		fmt.Printf("Putting ¢%v into your bank...\n", amount)

		if piggyBank >= 2000 {
			fmt.Printf("You have reached your target!\nFinal balance is: $%4.2f\n", float64(piggyBank)/100)
			break
		}

		fmt.Printf("New balance is: $%4.2f\n\n", float64(piggyBank)/100)
	}

}
