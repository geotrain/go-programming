package main

import (
	"fmt"
)

func main() {
	// Count = 0
	count := 0

	// For every pair of 4 digit numbers
	for a := 1000; a <= 9999; a++ {
		for b := a; b <= 9999; b++ { // don't count twice
			n := a * b

			// if a * b even ended
			s := fmt.Sprintf("%d", n)
			if s[0] == s[len(s)-1] {
				// count = count + 1
				count++
			}
		}
	}

	// print count
	fmt.Println(count)
}
