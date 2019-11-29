package main

import "fmt"

// Decipher the quote from Julius Caesar:
// L fdph, L vdz, L frqtxhuhg.
// Julius Caesar
// Your program will need to shift uppercase and
// lowercase letters by –3. Remember that 'a' becomes 'x', 'b'
// becomes 'y', and 'c' becomes 'z', and likewise for uppercase
// letters.

func main() {
	quote := "L fdph, L vdz, L frqtxhuhg."

	for i := 0; i < len(quote); i++ {
		c := quote[i]

		if c >= 'a' && c <= 'z' {
			c -= 3
			if c > 'z' {
				c += 26
			}
		} else if c >= 'A' && c <= 'Z' {
			c -= 3
			if c > 'Z' {
				c += 26
			}
		}
		fmt.Printf("%c", c)
	}
	fmt.Println("")
}
