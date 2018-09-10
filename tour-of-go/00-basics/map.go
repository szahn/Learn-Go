package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

//https://tour.golang.org/moretypes/19

func main() {
	var m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40, -74,
	}

	m["Google"] = Vertex{
		37, -122,
	}

	var p = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}

	a := make(map[string]int)
	a["Answer"] = 42
	a["Question"] = 43
	fmt.Println("The value is", a["Answer"])
	fmt.Println("The value is", a["Question"])

	//https://tour.golang.org/moretypes/22

	delete(m, "Question")

	_, exists := m["Question"]
	fmt.Println("Exists", exists)

	fmt.Println(m["Bell Labs"])
	fmt.Println(m["Google"])
	fmt.Println(p)
}
