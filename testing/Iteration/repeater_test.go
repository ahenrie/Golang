package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	result := Repeat("a")
	want := "aaaaa"

	if result != want {
		t.Errorf("Got: '%q' Wanted: '%q'", result, want)
	}
}

func ExampleRepeat() {
	result := Repeat("1")
	fmt.Println(result)
	// Output: 11111
}
