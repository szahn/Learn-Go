package main

import (
	"fmt"
	"math/rand"
	"math"
)

func main(){
	fmt.Println("Hello, ", rand.Intn(10))
	fmt.Printf("Hello %g\r\n", rand.Intn(10))
	fmt.Println("Hello, ", math.Pi)
}