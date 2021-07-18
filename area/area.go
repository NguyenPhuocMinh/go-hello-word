package area

import "math"

type Shape interface {
	area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	With, Height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) area() float64 {
	return r.With * r.Height
}

func GetArea(s Shape) float64 {
	return s.area()
}
