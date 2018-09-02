package main

import "fmt"

func main() {
	i, j := 24, 85
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(*p)
	p = &j
	fmt.Println(*p)
	*p = *p / 37
	fmt.Println(j)
	fmt.Println(*p)
}
