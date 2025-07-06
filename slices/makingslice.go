package main

import (
	"fmt"
)

// Creating a slice with make
// Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.

// The make function allocates a zeroed array and returns a slice that refers to that array:
// a := make([]int, 5)  // len(a)=5
// To specify a capacity, pass a third argument to make:

// b := make([]int, 0, 5) // len(b)=0, cap(b)=5

// b = b[:cap(b)] // len(b)=5, cap(b)=5
// b = b[1:]      // len(b)=4, cap(b)=4

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

	// Calling many slices
	Manyslice()

	// Append file
	appendit()

	// Copy slices
	sl := []int{1, 2, 3}
	t := make([]int, len(sl), (cap(sl)+1)*2)
	copy(t, sl)
	sl = t
	printSlices(sl)

	// To append one slice to another, use ... to expand the second argument to a list of arguments.
	first := []string{"John", "Paul"}
	second := []string{"George", "Ringo", "Pete"}
	first = append(first, second...) // equivalent to "append(a, b[0], b[1], b[2])"
	// a == []string{"John", "Paul", "George", "Ringo", "Pete"}
	fmt.Printf("First is : %v", first)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
