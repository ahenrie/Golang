package pointers

import "fmt"

func Swap(a *int, b *int) {
	fmt.Printf("a = %d, b = %d\n", *a, *b)

	*a = *a ^ *b
	*b = *b ^ *a
	*a = *a ^ *b

	fmt.Printf("a = %d, b = %d\n", *a, *b)
}
