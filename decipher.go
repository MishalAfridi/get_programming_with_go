package main

// Vignere cipher - write a program to decipher the string below using a keyword
// string.Repeat
// complete exercise without importing any packages other than fmt to print
// the deciphered message.
// use range in a loop and without a loop - range keyword splits a string into runes
// whereas keyword[0] results in bytes
// solve issue of wrapping around without using an if statement - use %
//

import (
	"fmt"
	"strings"
)

func main() {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	length := len(keyword) % len(cipherText)
	keywordStr := strings.Repeat(keyword, length)

	for i := 0; i < len(cipherText); i++ {
		c := cipherText[i] //byte a.k.a int8
		k := keywordStr[i]
		c -= k - byte('A')

		// not sure how to solve wrapping around
		// without a condition, with modulus
		if c < 'A' {
			c += 26
		}

		fmt.Printf("%c", c)
	}
	fmt.Println("")
}
