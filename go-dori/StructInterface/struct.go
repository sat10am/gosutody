package StructInterface

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

// func Perimeter(rectangle Rectangle) float64 {
// 	return 2 * (rectangle.Width + rectangle.Height)
// }

// func Area(rectangle Rectangle) float64 {
// 	return rectangle.Width * rectangle.Height
// }

type Circle struct {
	Radius float64
}

type Shape interface {
	Area() float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
