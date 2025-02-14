package pointers

import "fmt"

func Allocate(num int) *int {
	returnPointer := &num
	fmt.Printf("Number: %d --> Address: %p\n", num, returnPointer)
	return returnPointer
}
