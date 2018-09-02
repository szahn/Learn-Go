package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	//Slices: https://tour.golang.org/moretypes/7
	var s []int = primes[1:4]
	fmt.Println(s)

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)
	fmt.Println(q[1:])
	fmt.Println(q[:1])

	//https://tour.golang.org/moretypes/9
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	ss := []struct {
		i int
		b bool
	}{
		{2, true},
		{i: 3, b: false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(ss)

	var sss []int
	fmt.Println(sss, len(sss), cap(sss))
	if sss == nil {
		fmt.Println("nil!")
	}
}
