package main

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		assertCorrectMessage(t, got, want)
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		assertCorrectMessage(t, got, want)
	})
}

func ExampleSum() {
	testNumbers := [...]int{1000, 1, 1, 1, 1}
	sum := Sum(testNumbers)
	fmt.Println(sum)
	// Output: 1004
}

// Helper function for comparison of ints
func assertCorrectMessage(t testing.TB, got int, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
