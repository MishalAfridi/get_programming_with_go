package main

import (
	"fmt"
	"strings"
)

func main() {
	plainText := "your message goes here"
	keyword := "GOLANG"
	str := strings.Replace(plainText, " ", "", len(plainText))
	convertedText := strings.ToUpper(str)
	length := len(keyword) % len(convertedText)
	keywordStr := strings.Repeat(keyword, length)

	for i, c := range convertedText {
		// c is type rune, int32
		k := keywordStr[i] //uint8
		c += rune(k) - 'A'

		if c > 'Z' {
			c -= 26
		}

		fmt.Printf("%c", c) // int32
	}
	fmt.Println("")
}
