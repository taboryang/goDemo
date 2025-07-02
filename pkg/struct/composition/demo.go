package abc

import "fmt"

// Animal is an interface with a Speak method
type Animal interface {
	Speak() string
}

// Dog is a base struct
type Dog struct {
	Name string
}

// Speak implements Animal for Dog
func (d *Dog) Speak() string {
	return d.Name + " says: Woof! abc"
}

// Pet is a struct that "inherits" Dog via composition
type Pet struct {
	Dog   // embedded struct (composition)
	Owner string
}

// Speak overrides Dog's Speak method
func (p Pet) Speak() string {
	return fmt.Sprintf("%s is owned by %s and says: Woof!", p.Name, p.Owner)
}

func Demo() {
	var a Animal

	d := Dog{Name: "Buddy"}
	a = &d
	fmt.Println(a.Speak()) // Buddy says: Woof!

	p := Pet{
		Dog:   Dog{Name: "Max"},
		Owner: "Alice",
	}
	a = &p
	fmt.Println(a.Speak()) // Max is owned by Alice and says: Woof!
}
