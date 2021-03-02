/*
Generic stack example
A stack is a LIFO (last in first out) data collection
Our stack may hold values of any type
2021 - Ulrich Schramme
*/
package main

import (
	"errors"
	"fmt"
)

// our stack is based on a slice of empty interfaces
// empty interfaces can hold values of any type
type stack []interface{}

// pushing a value on the stack
func (stk *stack) push(arg interface{}) {
	*stk = append(*stk, arg)
}

// returns the topmost value
func (stk *stack) peek() (interface{}, error) {
	index := len(*stk) - 1
	if index < 0 {
		return nil, errors.New("Error! Stack is empty")
	}
	return (*stk)[index], nil
}

// returns the topmost value and removes it from the stack
func (stk *stack) pop() (interface{}, error) {
	value, err := stk.peek()
	if err == nil {
		*stk = (*stk)[:len(*stk)-1] // resizing the slice
	}
	return value, err
}

func main() {
	var mystack stack
	var value interface{}
	var err error
	const line = "-------------------------------\n"

	mystack.push(14555)                                         // pushing an int
	mystack.push(true)                                          // pushing a bool
	mystack.push("Another String ;-)")                          // pushing a string
	mystack.push(13.772)                                        // pushing a float
	mystack.push("This stack contains different types of data") // pushing a string
	mystack.push("Testing...")                                  // pushing a string

	fmt.Println("Testing our generic stack...")
	fmt.Println(line)
	fmt.Println("How many elements?", len(mystack))
	value, err = mystack.peek()
	if err == nil {
		fmt.Println("The result of peek() is:", value)
	} else {
		fmt.Println(err)
	}
	fmt.Println(line)
	fmt.Println("Printing all elements")
	fmt.Println(line)
	// executing pop() until the stack is empty
	for len(mystack) > 0 {
		value, err = mystack.pop()
		if err == nil {
			fmt.Println("The result of pop() is:", value)
		} else {
			fmt.Println(err)
		}
	}
	fmt.Println(line)
	fmt.Println("Execute peek() on an empty stack") // execute peek() or pop() on an empty stack causes an error
	value, err = mystack.peek()
	if err == nil {
		fmt.Println("The result of peek() is:", value)
	} else {
		fmt.Println(err)
	}
}
