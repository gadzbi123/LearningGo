package main

import "fmt"

type MyFloat float64

func absFloat(x MyFloat) float64 {
	if x < 0 {
		return float64(x)
	}
	return float64(x)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) scale(f float64) {
	v.X *= f
	v.Y *= f
}
func (v *Vertex) add(vv Vertex) {
	v.X += vv.X
	v.Y += vv.Y
}
func add(v *Vertex, vv Vertex) {
	v.X += vv.X
	v.Y += vv.Y
}
func main() {
	myVal := MyFloat(3.2)
	fmt.Println(absFloat(myVal))
	vert := Vertex{X: 2.3, Y: 22}
	vert.scale(2)
	fmt.Println(vert)
	vert2 := Vertex{X: 2.1, Y: 4.2}
	vert2.add(Vertex{2.2, 2.1})
	add(&vert2, Vertex{3.3, 2.2})
	fmt.Println(vert2)
}
