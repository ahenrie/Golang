package pointers

import "fmt"

func Deref(p *int) int {
	num := *p
	fmt.Printf("The Address: %p --> Contains: %d\n", p, num)
	return num
}
