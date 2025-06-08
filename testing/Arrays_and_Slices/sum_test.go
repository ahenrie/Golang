package arr_slices

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	got := sum(numbers)
	want := 15

	if got != want {
		t.Errorf("Got: %d Want: %d Given: %v", got, want, numbers)
	}
}

func Examplesum() {
	total := sum([]int{100, 1})
	fmt.Println(total)
	// Output: 101
}
