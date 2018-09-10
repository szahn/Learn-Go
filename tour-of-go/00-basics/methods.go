package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X float64
	Y float64
}

//https://tour.golang.org/methods/1

//By value (copy of)
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//By reference (pointer recieve)
//https://tour.golang.org/methods/4
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleF(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	fmt.Println(Vertex{-12, -4}.Abs())
	v := Vertex{3, 3}
	v.Scale(1.3)
	fmt.Println(v)
	//Pointers and functions: https://tour.golang.org/methods/5, https://tour.golang.org/methods/6
	ScaleF(&v, 2)
	v2 := Vertex{2, 2}
	//https://tour.golang.org/methods/6
	//https://tour.golang.org/methods/8
	v2.Scale(2)
	fmt.Println(v2)
}
