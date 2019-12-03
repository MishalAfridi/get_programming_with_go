package main

//convert strings to bools
// "true" "yes" or "1" are true - all strings
// "false" "no" or "0" are true - all strings
// display error for anything else

import "fmt"

func main() {
	x := "1"
	var y bool

	switch x {
	case "true", "yes", "1":
		y = true
	case "false", "no", "0":
		y = false
	default:
		fmt.Println(y, "is not valid")
	}
	fmt.Println("Ready for launch: ", y)
}
