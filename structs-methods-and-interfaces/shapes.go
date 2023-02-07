package structs_methods_and_interfaces

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.Width * rectangle.Height
}

type Circle struct {
	Radius float64
}

func (circle Circle) Area() float64 {
	return math.Pi * circle.Radius * circle.Radius
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
