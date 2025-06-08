package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	result := Repeat("a", 6)
	want := "aaaaaa"

	if result != want {
		t.Errorf("Got: '%q' Wanted: '%q'", result, want)
	}
}

func ExampleRepeat() {
	result := Repeat("1", 10)
	fmt.Println(result)
	// Output: 1111111111
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 3)
	}
}
