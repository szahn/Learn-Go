package main

import (
	"fmt"
	"math"

	"golang.org/x/tour/pic"
)

/* https://tour.golang.org/moretypes/18
Implement Pic. It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers.
When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.*/

func Pic(dx, dy int) [][]uint8 {

	fmt.Printf("Pic %vx%v\r\n", dx, dy)
	max := 255.0

	buffer := make([][]uint8, dy, dy)

	for y := 0; y < dy; y++ {
		buffer[y] = make([]uint8, dx, dx)

		color := math.Round((float64(y) / max) * max)
		altcolor := max - color

		for x := 0; x < dx; x++ {
			if y > x {
				buffer[y][x] = uint8(altcolor)
			} else {
				buffer[y][x] = uint8(color)
			}
		}
	}

	return buffer
}

func main() {
	pic.Show(Pic)
}
