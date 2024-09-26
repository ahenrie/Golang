package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

/*
Map
For the map, we have an input: [1, 2, 3], a function that mutates the input (n) => n * 2
Output = [2, 4, 6]

Without generics:
func MapValues(values []int, mapFunc func(int) int) []int {
	var newValues []int
	for _, v := range values {
		newValue := mapFunc(v)
		newValues = append(newValues, newValue)
	}
	return newValues
}

func main() {
	result := MapValues([]int{1, 2, 3}, func(n int) int {
		return n * 2
	})
	fmt.Printf("result: %+v\n", result)
}
*/

func MapValues[T constraints.Ordered](values []T, mapFunc func(T) T) []T {
	var newValues []T
	for _, v := range values {
		newValue := mapFunc(v)
		newValues = append(newValues, newValue)
	}
	return newValues
}

func main() {
	result := MapValues([]int{1, 2, 3}, func(n int) int {
		return n * 2
	})

	fmt.Printf("result: %+v\n", result)

	resultFloat := MapValues([]float64{12.2, 123.2, 13.1}, func(n float64) float64 {
		return n * 1.2
	})

	fmt.Printf("result: %+v\n", resultFloat)
}
