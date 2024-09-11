package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}
func (c *Circle) Circ() float64 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Height, Width float64
}

func (c *Rectangle) Area() float64 {
	return c.Height * c.Width
}

type Areable interface {
	Area() float64
}

type Circumferencable interface {
	Circ() float64
}

func main() {
	circ := &Circle{5.}
	rect := &Rectangle{10., 20.}
	var x Areable = circ
	fmt.Println("circ area", x.Area())
	x = rect
	fmt.Println("rect area", x.Area())
	var y Circumferencable = circ
	fmt.Println("circ circ", y.Circ())
	// y = rect
	// fmt.Println("rect circ", y.Circ())
}
