package main

import (
	"fmt"
	"runtime"
)

//https://tour.golang.org/flowcontrol/9

func main() {
	switch os := runtime.GOOS; os {
	case "darwin":
		{
			fmt.Println("OSX")
		}
	case "Linux":
		fmt.Println("Linux")
	default:
		{
			fmt.Printf("OS: %s\r\n", os)
		}
	}

}
