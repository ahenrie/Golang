package main

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := [...]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given the array: %v", got, want, numbers)
	}
}

func ExampleSum() {
	testNumbers := [...]int{1000, 1, 1, 1, 1}
	sum := Sum(testNumbers)
	fmt.Println(sum)
	// Output: 1004
}
