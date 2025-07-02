package main

import "fmt"

func main() {
	// Using new: allocates zeroed storage, returns pointer
	p1 := new([]int)                    // *([]int), pointer to nil slice
	fmt.Println("new([]int):", p1, *p1) // address, nil

	// Using make: allocates and initializes slice, map, or channel
	s1 := make([]int, 3) // []int{0,0,0}
	fmt.Println("make([]int, 3):", s1)

	// new for map
	p2 := new(map[string]int) // *map[string]int, pointer to nil map
	// (*p2)["one"] = 1                  // This line would panic because p2 is nil
	fmt.Println("new(map):", p2, *p2) // address, map[one:1]

	// make for map
	m1 := make(map[string]int)
	m1["one"] = 1
	fmt.Println("make(map):", m1)

	// new for channel
	p3 := new(chan int) // *chan int, pointer to nil channel
	// (*p3) <- 1                         // This line would panic because p3 is nil
	fmt.Println("new(chan):", p3, *p3) // address, nil

	// make for channel
	c1 := make(chan int, 2)
	c1 <- 1
	fmt.Println("make(chan):", c1)

	// new for struct
	type Person struct {
		Name string // zeroed string
		Age  int    // zeroed int
	}

	p4 := new(Person)                    // *Person, pointer to zeroed struct
	fmt.Println("new(Person):", p4, *p4) // address, zeroed
	// make for struct (not applicable, use new)
	// make cannot be used for structs, only slices, maps, and channels
	// p5 := make(Person) // invalid, cannot use make for struct

	p5 := &Person{Name: "Alice", Age: 30}       // &Person, pointer to initialized struct
	fmt.Println("Initialized Person:", p5, *p5) // address, initialized
}
