package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func AddInt(a int, b int) int {
	return a + b
}

func AddFloat(a float64, b float64) float64 {
	return a + b
}

/*
Say we want to create a simple add function that can add whatever datatypes.
How can we use generics to accomplish this?
*/

func AddGenerics[T int | float64](a T, b T) T {
	return a + b
}

// There are lots of different types in Go so the approach is above is too verbose.
// We can create our own datatype to interface with a generic function.
type Num interface {
	int | int8 | int16 | float32 | float64
}

func AddInteface[T Num](a T, b T) T {
	return a + b
}

// This is still too verbose. We can use constraints.
func Add[T constraints.Ordered](a T, b T) T {
	return a + b
}

func main() {
	result := Add(1, 2.1)
	fmt.Printf("result: %+v\n", result)
}
