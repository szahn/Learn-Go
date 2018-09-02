package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

//structs to pointers: https://tour.golang.org/moretypes/4

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	pv = &Vertex{1, 2} // has type *Vertex
)

func main() {
	v := Vertex{1, 2}
	v.Y++
	p := &v
	p.X += 3e3
	v.Y *= 2
	fmt.Println(v)
	fmt.Println(p)

	fmt.Println(v1, pv, v2, v3)

}
