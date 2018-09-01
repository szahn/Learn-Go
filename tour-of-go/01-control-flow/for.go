package main

import "fmt"

//https://tour.golang.org/flowcontrol/1

func basicLoop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func loop() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func main() {
	basicLoop()
	loop()
}
