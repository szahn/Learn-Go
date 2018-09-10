package main

import (
	"fmt"
	"math"
)

func add(a int, b int) int {
	return a + b
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(num int) (x int, y int) {
	x = num * 8 / 9
	y = x - 5
	return
}

//https://tour.golang.org/moretypes/24
func compute(f func(a int, b int) int) int {
	return f(2, 3)
}

func computef(f func(a float64, b float64) float64) float64 {
	return f(2, 3)
}

func main() {
	fmt.Println("Sum", add(1, 2))
	a, b := swap("a", "b")
	fmt.Println("Swap", a, b)
	fmt.Println(split(20))
	fmt.Println(compute(func(a int, b int) int {
		return a * b
	}))
	fmt.Println(computef(math.Pow))
}
