package main

import (
	"fmt"
)

// calculates the least common multiple
func euclid(a, b int) int {
	if b == 0 {
		return a
	}
	return euclid(b, a%b) // recursive call
}

func main() {
	const text = "Least common multiple of"
	// Examples
	x, y := 65, 26
	fmt.Printf("%s %d and %d = %d\n", text, x, y, euclid(x, y))
	x, y = 224, 12
	fmt.Printf("%s %d and %d = %d\n", text, x, y, euclid(x, y))
	x, y = 1605, 24
	fmt.Printf("%s %d and %d = %d\n", text, x, y, euclid(x, y))
	x, y = 77, 392
	fmt.Printf("%s %d and %d = %d\n", text, x, y, euclid(x, y))
	x, y = 109, 15
	fmt.Printf("%s %d and %d = %d\n", text, x, y, euclid(x, y))
}
