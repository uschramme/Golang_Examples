package main

import (
	"fmt"
)

// variadic functions accept a variable number of arguments
// fmt.Println() is a well-known example

// calculates the average
func Average(values ...float64) float64 {
	sum := 0.0
	for _, num := range values {
		sum += num
	}
	return sum / float64(len(values))
}

// calculates the average, too
// you have to pass at leat one value
// 0 values is not allowed!
func BetterAverage(value float64, values ...float64) float64 {
	sum := 0.0
	for _, num := range values {
		sum += num
	}
	sum += value
	return sum / float64(len(values)+1)
}

func main() {
	fmt.Println(Average(2.6, 3.8, 14.258))

	// you can pass a slice, too
	// but a three-dotted ellipsis is necessary
	fmt.Println(Average([]float64{18.66, 2.45, 7.77777}...))

	// you cann pass 0 values but this can lead to problems
	// here you get NaN (not a number)
	fmt.Println(Average())

	// function BetterAverage avoids this problem because it doesn´t accept 0 values
	fmt.Println(BetterAverage(17.23, 44.19, 62.915644, -13.8))
}
