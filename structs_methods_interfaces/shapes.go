package main

import "math"

// Shape is the generic interface
type Shape interface {
	Area() float64
}

// Rectangle holds the height and width
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle holds t he radius
type Circle struct {
	Radius float64
}

// Triangle holds the base and Height
type Triangle struct {
	Base   float64
	Height float64
}

// Perimeter returns the perimeter given width and height
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area returns the area given width and height
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area returns the area given radius
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Area returns the area given the base and height
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
