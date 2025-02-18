// interface is named collection of method signatures
// @todo: read -> https://research.swtch.com/interfaces

package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

// define types
type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// implement interface for types
func (r rect) area() float64 {
	return r.height * r.width
}

func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// function with interface type
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius:", c.radius)
	}
}

// use a type switch
func detectShape(g geometry) {
	switch s := g.(type) {
	case circle:
		fmt.Println("circle with radius:", s.radius)
	case rect:
		fmt.Println("rect of dimensions:", s.width, s.height)
	default:
		fmt.Println("unkown shape")
	}
}

func main() {
	r := rect{width: 5, height: 12}
	c := circle{radius: 7}

	// both shapes implement interface and used in functions with interface type arg
	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)

	detectShape(r)
	detectShape(c)
}
