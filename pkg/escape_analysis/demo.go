package main

import "fmt"

// stackAllocated demonstrates a variable that does NOT escape to the heap.
func stackAllocated() int {
	x := 42 // x is allocated on the stack
	return x
}

// heapAllocated demonstrates a variable that escapes to the heap.
func heapAllocated() *int {
	y := 99 // y escapes to the heap because its address is returned
	return &y
}

// go build -gcflags=-m .
func main() {
	a := stackAllocated()
	b := heapAllocated()
	fmt.Println("stackAllocated:", a)
	fmt.Println("heapAllocated:", *b)
}
