package shapes

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Perimeter() float64 {
	return math.Pi * c.Radius * 2
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Perimeter() float64 {
	// https://math.stackexchange.com/a/80401
	return t.Base + t.Height + math.Sqrt((t.Base*t.Base)+(t.Height*t.Height))
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height * 0.5
}
