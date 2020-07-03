package main

// Rectangle holds the height and width
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter returns the perimeter given width and height
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area returns the area given width and height
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
