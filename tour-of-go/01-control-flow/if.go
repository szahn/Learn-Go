package main

import (
	"fmt"
	"math"
)

func pow(n, exp, lim float64) float64 {
	if v := math.Pow(n, exp); v < lim {
		return v
	} else {
		fmt.Println("%v > %v\n", v, lim)
		return lim
	}

}

func main() {
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
}
