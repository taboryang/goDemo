package demo

import "fmt"

// ExportedStruct is visible outside the package
type ExportedStruct struct {
	Name string
}

// unexportedStruct is only visible inside the package
type unexportedStruct struct {
	age int
}

// ExportedMethod is visible outside the package
func (e *ExportedStruct) ExportedMethod() {
	fmt.Println("This is an exported method. Name:", e.Name)
}

// unexportedMethod is only visible inside the package
func (e *ExportedStruct) unexportedMethod() {
	fmt.Println("This is an unexported method. Name:", e.Name)
}

// ExportedFunc is visible outside the package
func ExportedFunc() {
	fmt.Println("This is an exported function.")
}

// unexportedFunc is only visible inside the package
func unexportedFunc() {
	fmt.Println("This is an unexported function.")
}
