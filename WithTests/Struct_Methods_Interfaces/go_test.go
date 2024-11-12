package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10, 10}
	got := Perimeter(rectangle)
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
}

func TestAreaInterfaces(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(t, rectangle, 72)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}

func TestMethods(t *testing.T) {

	t.Run("rectangle area", func(t *testing.T) {
		rec := Rectangle{10, 2}
		got := rec.Area()
		want := 20.0

		if got != want {
			t.Errorf("got: %.2f want: %.2f", got, want)
		}
	})

	t.Run("rectangle perimeter", func(t *testing.T) {
		rec := Rectangle{10, 2}
		got := rec.Perimeter()
		want := 24.0

		if got != want {
			t.Errorf("got: %.2f want: %.2f", got, want)
		}
	})

	t.Run("circle area", func(t *testing.T) {
		cir := Circle{10}
		got := cir.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got: %.2f want: %.2f", got, want)
		}
	})

	t.Run("circle perimeter", func(t *testing.T) {
		cir := Circle{10}
		got := cir.Perimeter()
		want := 62.83

		if got != want {
			t.Errorf("got: %.2f want: %.2f", got, want)
		}
	})
}

/*
	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		got := Area(circle)
		want := 314.1592653589793

		if got != want {
			t.Errorf("got: %g want: %g", got, want)
		}
	})
}
*/
