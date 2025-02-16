package pointers

import "fmt"

func Add(value int, number *int) {
	*number += value
	fmt.Printf("The pointer at: %p had %d added to it. Now: %d\n", number, value, *number)
}
