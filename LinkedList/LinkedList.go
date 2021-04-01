/* a singly linked list
best way to understand the principle is to use the debugger...
don´t use it for your applications
you will find a (much better) doubly linked list in "container/list" */

package main

import (
	"errors"
	"fmt"
)

// list-node
// every node has a pointer to the next node
type node struct {
	next  *node
	value int
}

// the list-struct
type LinkedList struct {
	length int
	data   *node
}

// add a node
func (list *LinkedList) Add(element *node) {
	if list.data == nil {
		// list is still empty
		list.data = element
		list.length++ // increase length
	} else {
		curr := list.data
		// move to last node
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = element // add the new one after the last
		list.length++       // increase length
	}
}

// delete a node
func (list *LinkedList) Delete(val int) {
	if list.data == nil {
		return // makes no sense with an empty list...
	}
	// first node is to be deleted
	if list.data.value == val {
		list.data = list.data.next
		list.length-- // decrease length
		return
	}
	// try the other nodes...
	prev := list.data
	curr := list.data.next
	for curr.next != nil {
		prev = prev.next
		curr = curr.next
		if curr.value == val {
			prev.next = curr.next
			list.length-- // decrease length
		}
	}
}

// is a value in list?
func (list *LinkedList) FindValue(val int) bool {
	if list.data == nil {
		return false // list is empty
	} else {
		curr := list.data
		// move to last node
		for curr != nil {
			if curr.value == val {
				return true // value fount -> return
			}
			curr = curr.next
		}
		return false
	}
}

// Get value at a certain index (starts at 0)
func (list *LinkedList) GetValue(index int) (int, error) {
	if list.data == nil {
		return 0, errors.New("List is empty") // empty list
	}
	if index < 0 || index >= list.length {
		return 0, errors.New("Index out of bounds") // invalid index
	}
	pos := 0
	curr := list.data
	if pos == index {
		return curr.value, nil
	}
	for pos != index {
		pos++
		curr = curr.next
	}
	return curr.value, nil
}

// print all elements
func (list *LinkedList) Traverse() {
	if list.data == nil {
		fmt.Println("List is empty!")
	} else {
		curr := list.data
		// move to last node
		for curr != nil {
			fmt.Printf("%d\t", curr.value)
			curr = curr.next
		}
	}
}

func main() {
	var list LinkedList
	// adding some nodes
	list.Add(&node{value: 17})
	list.Add(&node{value: 1502})
	list.Add(&node{value: 122})
	list.Add(&node{value: 308})
	list.Add(&node{value: 1961})
	list.Add(&node{value: 999})
	// some tests
	fmt.Println("Length of list:", list.length)
	list.Traverse()
	fmt.Println()
	val, err := list.GetValue(0)
	if err == nil {
		fmt.Println("Value at index 0", val)
	} else {
		fmt.Println(err)
	}
	val, err = list.GetValue(2)
	if err == nil {
		fmt.Println("Value at index 2", val)
	} else {
		fmt.Println(err)
	}
	fmt.Println()
	fmt.Println("Is 308 in list?", list.FindValue(308))
	fmt.Println("Is 520 in list?", list.FindValue(520))
	fmt.Println()
	list.Delete(17)
	fmt.Println("Length of list after deleting 17:", list.length)
	list.Traverse()
	fmt.Println()
	fmt.Println()
	list.Delete(308)
	fmt.Println("Length of list after deleting 308:", list.length)
	list.Traverse()
	fmt.Println()
	val, err = list.GetValue(2)
	if err == nil {
		fmt.Println("Value at index 2", val)
	} else {
		fmt.Println(err)
	}
}
