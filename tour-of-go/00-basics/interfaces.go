package main

import (
	"fmt"
	"math"
)

type AbsoluteValue interface {
	Abs() float64
}

type MyFloat float64

type Vertex struct {
	X float64
	Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X) + math.Sqrt(v.Y*v.Y)
}

func (f MyFloat) Abs() float64 {
	return math.Abs(float64(f))
}

//https://tour.golang.org/methods/9
//https://tour.golang.org/methods/10
func main() {
	v := &Vertex{2, 3}
	fmt.Printf("%v %v \r\n", *v, v.Abs())

	var v2 AbsoluteValue
	v2 = MyFloat(-12)
	fmt.Println(v2.Abs())

	v2 = v
	fmt.Println(v2.Abs())

	//The empty interface
	//https://tour.golang.org/methods/14
	var i interface{}
	fmt.Printf("%v %T\r\n", i, i)
	i = 32
	fmt.Printf("%v %T\r\n", i, i)

}
