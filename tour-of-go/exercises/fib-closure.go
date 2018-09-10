package main

import "fmt"

//https://tour.golang.org/moretypes/26

func fib(n int) int {
	if n == 0 || n == 1 || n == 2 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

func fibClosure(n int) func() int {
	return func() int { return fib(n) }
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(fibClosure(i)())
	}
}
