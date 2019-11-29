package main

import (
	"fmt"
)

// CIPHER the Spanish message “Hola Estación Espacial Internacional”
// with ROT13. Modify listing 9.7 to use the range keyword.
// Now when you use ROT13 on Spanish text, characters with accents
// are preserved.

func main() {
	message := "Hola Estación Espacial Internacional"

	for _, character := range message {
		if character >= 'a' && character <= 'z' {
			character += 13
			if character < 'z' {
				character -= 26
			}
		} else if character >= 'A' && character <= 'Z' {
			character += 13
			if character < 'Z' {
				character -= 26
			}
		}
		fmt.Printf("%c", character)
	}
}
