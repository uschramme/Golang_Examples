package main

import (
	"fmt"
)

/* If you are a C/C++ developer...
Go doesn't support pointer arithmetic.
Instead of using a pointer on an array you should use a slice.
Slices are convenient and safe. */

func main() {
	// A pointer stores the adress of a variable. I use an int as example. Other types work as well
	number := 1961      // Declaring an int variable
	pnum := &number     // Declaring a pointer to number. The ampersand is called address operator
	fmt.Println(pnum)   // Prints the adress of number (not very interesting...)
	fmt.Println(*pnum)  // Prints the value. The asterisk is called dereference operator
	*pnum = 2021        // Assigning a value to the pointer
	fmt.Println(number) // number is changed as well
	fmt.Println()

	// You can declare a pointer to a pointer
	ppnum := &pnum
	**ppnum = 1981      // Assigning a value to ppnum. You have to use two asterisks for dereferencing
	fmt.Println(number) // number is changed as well
	fmt.Println()

	// Pointer as function argument
	ReverseNumber(&number)
	fmt.Println(number)
	fmt.Println()

	// Pointers are an effective way to pass big objects to functions or collections
	// It is cheaper to pass an address than a big copy
	// You have to pass the adress by using the ampersand
	mybook := Book{"Fear and Loathing in Las Vegas", "	Hunter S. Thompson"}
	PrintBook(&mybook)
	fmt.Println()

	// Functions can return pointers as well
	message := Pointer2String()
	fmt.Println(*message)
	fmt.Println()

	// Pointers to functions are implicitly in Go
	f := func() {
		fmt.Println("Pointers to functions are implicitly in Go!")
	}
	f()
}

// Example for passing an argument by reference (by pointer)
// The function just multiplies the given number by -1
// The given number keeps its value after leaving the function
func ReverseNumber(arg *int) {
	*arg *= -1
}

// A struct for testing
type Book struct {
	Title  string
	Author string
}

// The function expects a pointer to a Book-struct
func PrintBook(book *Book) {
	fmt.Printf("'%s' is written by %s.\n", book.Title, book.Author)
}

// This function returns a pointer to a string
func Pointer2String() *string {
	message := "Returning a pointer is easy!"
	return &message
}
