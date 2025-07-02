package main

import (
	demo "dops-git.fortinet-us.com/demo/pkg/visibility/util"
)

func main() {
	s := demo.ExportedStruct{Name: "Alice"}
	s.ExportedMethod()
	// s.unexportedMethod() // This will not work because unexportedMethod is not accessible outside the package
	demo.ExportedFunc()
	// demo.unexportedFunc() // This will not work because unexportedFunc is not accessible outside the package
	// unexported := demo.UnexportedStruct{Age: 30}
}
