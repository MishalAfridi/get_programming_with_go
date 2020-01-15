package main

import "fmt"

// kelvinToCelcius converts ˚K to ˚C
func kelvinToCelcius(k float64) float64 {
	k -= 273.15
	return k
}

// celciusToFahreheit converts ˚C to ˚F
func celciusToFahreheit(c float64) float64 {
	f := (c * 9.0 / 5.0) + 32.0
	return f
}

// kelvinToFahrenheit converts ˚K to ˚F by reusing methods
func kelvinToFahrenheit(k float64) float64 {
	c := kelvinToCelcius(k)
	f := celciusToFahreheit(c)
	return f
}

func main() {
	kelvin := 294.0
	celcius := kelvinToCelcius(kelvin)
	fmt.Print(kelvin, "˚K is ", celcius, "˚C\n")

	kelvin = 233
	celcius = kelvinToCelcius(kelvin)
	fmt.Print(kelvin, "˚K is ", celcius, "˚C\n")

	celcius = 30.0
	fahrenheit := celciusToFahreheit(celcius)
	fmt.Print(celcius, "˚C is ", fahrenheit, "˚F\n")

	kelvin = 0.0
	fahrenheit = kelvinToFahrenheit(kelvin)
	fmt.Print(kelvin, "˚K is ", fahrenheit, "˚F\n")
}
