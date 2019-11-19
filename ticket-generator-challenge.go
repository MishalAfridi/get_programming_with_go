package main

import (
	"fmt"
	"math/rand"
	"os"
	"text/tabwriter"
)

func main() {
	spacelines := [3]string{"Space Adventures", "SpaceX", "Virgin Galactic"}
	types := [2]string{"One-way", "Round-trip"}

	const departureDate = "13/10/2020"
	const hoursPerDay = 24
	const distance = 62100000 //km

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 5, 0, 1, ' ', tabwriter.AlignRight)
	fmt.Fprintln(w, "\t      \t Spaceline \tDays\t Trip Type \t Price \t")
	fmt.Fprintln(w, "\t  =========================================")

	for i := 0; i < 10; i++ {
		spaceline := spacelines[rand.Intn(3)]
		speed := (rand.Intn(14) + 16) // km/s
		speedPerHour := speed * 3600
		durationInDays := distance / speedPerHour / hoursPerDay
		tripType := types[rand.Intn(2)]
		price := speed + 20

		if tripType == "Round-trip" {
			price = price * 2
		}

		fmt.Fprintln(w, "\t", spaceline, "\t", durationInDays, "\t", tripType, "\t $", price, "\t")
	}
	w.Flush()
}
