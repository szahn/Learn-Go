package main

import (
	"fmt"
)

const Pi = 3.14

//https://tour.golang.org/basics/16

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(n int) int {
	return n*10 + 1
}

func needFloat(f float64) float64 {
	return f * 0.1
}

func main() {
	const World = "世界"
	fmt.Println("Hello, ", World)
	fmt.Println("Happy", Pi, "Day")
	fmt.Printf("Int %v Float %v\r\n", needInt(10), needFloat(99))
}
