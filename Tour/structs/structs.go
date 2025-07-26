package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Printf("%v\n", Vertex{1, 2})
	// Go fields are accessed using dot operator
	fmt.Printf("%v", Vertex{1, 2}.X)

	// Struct fields can be accessed through a struct pointer.
	// (*p).a can be cumbersome so language lets use use p.a
	v := Vertex{1, 2}
	p := &v
	// 1e1 means 10
	p.X = 1e1
	fmt.Println(v)

	// Struct Literals
	fmt.Println("Struct Literals")
	v1 := Vertex{1, 2} // has type Vertex
	v2 := Vertex{X: 1} // Y:0 is implicit
	v3 := Vertex{}     // X:0 and Y:0
	point := &Vertex{21, 2}
	fmt.Println(v1, point, v2, v3)
	fmt.Println(point.X)
}
