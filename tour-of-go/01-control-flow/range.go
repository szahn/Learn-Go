package main

import "fmt"

//https://tour.golang.org/moretypes/16

func main() {
	pow := make([]int, 10)

	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
