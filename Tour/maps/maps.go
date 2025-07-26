package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex
var z map[int]int

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	z = make(map[int]int)
	z[6] = 9
	fmt.Println(z[6])

	// Other way to create map
	var ma = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}

	//	If the top-level type is just a type name, you can omit it from the elements of the literal.
	ma = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(ma)

	// Mutating Maps
	fmt.Println("Mutating Maps")
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
