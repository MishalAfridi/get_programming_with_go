package main

//write a program with celsius, fahrenheit, and kelvin types
//and methods to convert them from any temperature type
//to any other temperature type

import "fmt"

type celsius float64
type kelvin float64
type fahrenheit float64

//converts ˚C to ˚F and ˚K
func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

//converts ˚K to ˚C and ˚F
func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}
func (k kelvin) fahrenheit() fahrenheit {
	return k.celsius().fahrenheit()
}

//converts ˚F to ˚C and ˚K
func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}
func (f fahrenheit) kelvin() kelvin {
	return f.celsius().kelvin()
}

func main() {
	var c celsius = 35
	var k kelvin = 300
	var f fahrenheit = 7

	kToC := k.celsius()
	fToC := f.celsius()
	cToF := c.fahrenheit()
	kToF := k.fahrenheit()
	cToK := c.kelvin()
	fToK := f.kelvin()

	fmt.Print(k, "˚K is ", kToC, "˚C and ", kToF, "˚F\n")
	fmt.Print(c, "˚C is ", cToF, "˚F and ", cToK, "˚K\n")
	fmt.Print(f, "˚F is ", fToC, "˚C and ", fToK, "˚K\n")
}
