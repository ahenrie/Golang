package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10, 10}
	got := Perimeter(rectangle.Width, rectangle.Height)
	want := 40.0

	if got != want {
		t.Errorf("got: %.2f want: %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	t.Run("rectagles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		got := Area(rectangle)
		want := 72.0

		if got != want {
			t.Errorf("got: %.2f want: %.2f", got, want)
		}
	})

	t.Run("circle", func(t *testing.T) {
		circle = Circle{10}
		got := Area(circle)
		want := 314.1592653589793

		if got != want {
			t.Errorf("got: %g want: %g", got, want)
		}
	})
}
