package main

import(
	"fmt"
)

func add(a int, b int) int{
	return a + b
}

func swap(x, y string) (string, string){
	return y, x
}

func split(num int) (x int, y int){
	x = num * 8 / 9
	y = x - 5
	return
}

func main(){
	fmt.Println("Sum", add(1, 2))
	a, b := swap("a", "b")
	fmt.Println("Swap", a, b)
	fmt.Println(split(20))
}