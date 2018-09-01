package main

import(
	"fmt"
	"math/cmplx"
)

var (
	ToBe	bool	= false
	MaxInt64	uint64	= 1<<64 -1
	MaxInt8	uint8	= 1<<8 - 1
	MaxInt16	uint16	= 1<<16 - 1
	z		complex128 = cmplx.Sqrt(-5 + 12i)
)


func main(){
	fmt.Printf("Type %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type %T Value: %v\n", MaxInt64, MaxInt64)
	fmt.Printf("Type %T Value: %v\n", MaxInt16, MaxInt16)
	fmt.Printf("Type %T Value: %v\n", MaxInt8, MaxInt8)
	fmt.Printf("Type %T Value: %v\n", z, z)

	//Variables declared without an explicit initial value are given their zero value.
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}