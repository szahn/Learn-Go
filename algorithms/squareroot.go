package main

import (
	"fmt"
	"math"
)

//Use Newton's method to find the square root
//https://tour.golang.org/flowcontrol/8
func Sqrt(x float64) float64 {
	z := x / 2
	last := z
	precision := .000001

	for {
		square := z * z
		if square == x {
			return z
		} else if last-z < precision && math.Abs(x-square) < precision {
			return z
		}

		last = z
		z -= (z*z - x) / (2 * z)
	}
}

func main() {

	var assertSqrt = func(n float64) {
		sqrts := []float64{Sqrt(n), math.Sqrt(n)}
		fmt.Printf("%v: %v %v %v\r\n", n, sqrts[0], sqrts[1], sqrts[0] == sqrts[1])
	}

	assertSqrt(9)
	assertSqrt(13)
	assertSqrt(25)
	assertSqrt(99)
	assertSqrt(100)
	assertSqrt(999)
	assertSqrt(1000)
	assertSqrt(10000)
	assertSqrt(12345)
}
