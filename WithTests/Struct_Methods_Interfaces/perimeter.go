package main

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func Perimeter(width, height float64) float64 {
	return (width + height) * 2
}

func Area(width, height float64) float64 {
	return width * height
}

func Area(c Circle) float64 {
	return math.Pi * c.Radius * c.Radius
}
