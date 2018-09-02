package main

import (
	"fmt"
)

//https://tour.golang.org/flowcontrol/13
//https://blog.golang.org/defer-panic-and-recover

func main() {
	i := 0
	defer fmt.Println("Goodbye Cruel World!", i)

	i++

	for j := 0; j < 10; j++ {
		defer fmt.Println(j)
	}

	fmt.Println("Hello World", i)
}
