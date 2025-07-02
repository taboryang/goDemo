package demo

import "fmt"

type Counter struct {
	Value int
}

// Method with value receiver (struct object)
func (c Counter) IncrementByValue() {
	c.Value++
}

// Method with pointer receiver (struct pointer)
func (c *Counter) IncrementByPointer() {
	c.Value++ // This modifies the original struct
}

func DemoReceiverDifference() {
	// call flexible methods with value and pointer receivers
	c1 := &Counter{Value: 10}
	c2 := Counter{Value: 10}

	c1.IncrementByValue()                            // Go auto-dereferences (*c1)
	fmt.Println("After IncrementByValue:", c1.Value) // Output: 10

	c2.IncrementByPointer()                            // Go auto-takes address (&c2)
	fmt.Println("After IncrementByPointer:", c2.Value) // Output: 11
}
