package main

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return roundToTwoDecimalPlaces(math.Pi * 2 * c.Radius)
}

func (r Rectangle) Perimeter() float64 {
	return (r.Height + r.Width) * 2
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

// Generic functions
func Perimeter(r Rectangle) float64 {
	return (r.Height + r.Width) * 2
}

func Area(r Rectangle) float64 {
	return r.Height * r.Width
}

// Helper functionot round float to 2 decimal places
func roundToTwoDecimalPlaces(value float64) float64 {
	return math.Round(value*100) / 100
}
