package main

import (
	"fmt"
)

// Generic function with type parameter T
func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

// Generic type with type parameter T
type Pair[T any, U any] struct {
	First  T
	Second U
}

type Comparable interface {
	int | float64 | ~string
}

func GMin[T Comparable](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func main() {
	// Using generic function
	ints := []int{1, 2, 3}
	strs := []string{"hello", "world"}
	PrintSlice(ints)
	PrintSlice(strs)

	// Using generic type
	p := Pair[int, string]{First: 42, Second: "answer"}
	fmt.Println(p)

	fmt.Println("GMin(1, 2) is: ", GMin(1, 2))
	fmt.Println("GMin(3.14, 2.71) is: ", GMin(3.14, 2.71))
	fmt.Println("GMin(\"apple\", \"banana\") is: ", GMin("apple", "banana"))

	q := Pair[int, string]{First: 24, Second: "twenty-four"}
	fmt.Println(q)
	// GMin(p, q) // This will not compile because Pair does not implement Comparable
}
