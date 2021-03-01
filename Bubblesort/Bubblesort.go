/*
Bubblesort example
2021 - Ulrich Schramme
*/

package main

import (
	"fmt"
)

// type 'compare'
type compare func(arg1 float64, arg2 float64) bool

// function to compare for ascending (tpye 'compare')
func compasc(arg1 float64, arg2 float64) bool {
	return arg1 > arg2
}

// function to compare for ascending (tpye 'compare')
func compdesc(arg1 float64, arg2 float64) bool {
	return arg1 < arg2
}

// simple implementation of the bubblesort-algorithm
func bubblesort(numbers *[]float64, compfunc compare) {
	max := len(*numbers) - 1
	if max > 0 {
		flag := true // swapped?
		for flag {
			flag = false
			for i := 0; i < max; i++ {
				if compfunc((*numbers)[i], (*numbers)[i+1]) {
					flag = true
					(*numbers)[i], (*numbers)[i+1] = (*numbers)[i+1], (*numbers)[i] // swap values
				}
			}
		}
	}
}

func main() {
	fmt.Println("Testing Bubblesort...")
	fmt.Println("---------------------")
	numbers := []float64{5.72, 16.88, 3.52, 105.16, 44.69}
	fmt.Println("Unsorted:\t\t", numbers)
	bubblesort(&numbers, compasc)
	fmt.Println("Ascending:\t\t", numbers)
	bubblesort(&numbers, compdesc)
	fmt.Println("Descending:\t\t", numbers)
}
