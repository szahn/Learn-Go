package main

import "fmt"

//every number after the first two is the sum of the two preceding ones

func fib(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}

	return fib(n-2) + fib(n-1)
}

func main() {
	for i := 1; i < 100; i++ {
		fmt.Println("Fibonnacci", i, fib(i))
	}
}
