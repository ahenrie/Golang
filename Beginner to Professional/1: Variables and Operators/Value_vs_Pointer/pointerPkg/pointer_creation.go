package pointerpkg

import (
	"fmt"
	"time"
)

func createPointers {
	// variable declaration with var
	var count1 *int
	// variable with new()
	count2 := new(int)
	// Create a number to use its address
	countTemp := 4
	// Using &, create a pointer to an existing variable
	count3 := &countTemp
	// Create a pointer from some types without a temp variable
	t := &time.Time{}
	// Print each out
	fmt.Printf("count1: %#v\n", count1)
	fmt.Printf("count2: %#v\n", count2)
	fmt.Printf("count3: %#v\n", count3)
	fmt.Printf("time : %#v\n", t)
}
