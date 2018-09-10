package main

import "fmt"

//https://tour.golang.org/methods/15
func main() {
	//Type assertions
	//https://tour.golang.org/methods/15
	var i interface{} = 32
	s, ok := i.(string)
	fmt.Println(s, ok)

}
