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

	// Learn Slices
	var s []int = primes[1:4]
	fmt.Println(s)

	// Slices are like references to arrays
	// Changing the elements of a slice modifies the corresponding elements of its underlying array.
	// Other slices that share the same underlying array will see those changes.
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	slicea := names[0:2]
	sliceb := names[1:3]
	fmt.Println(slicea, sliceb)

	sliceb[0] = "XXX"
	fmt.Println(slicea, sliceb)
	fmt.Println(names)

	// A slice literal is like an array literal without the length.
	structs := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(structs)

	// 	A slice has both a length and a capacity.

	// The length of a slice is the number of elements it contains.

	// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

	// The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).

	// You can extend a slice's length by re-slicing it, provided it has sufficient capacity.

	fmt.Println("Slice len")

	slicelen := []int{2, 3, 5, 7, 11, 13}
	printSlice(slicelen)

	// Slice the slice to give it zero length.
	slicelen = slicelen[:0]
	printSlice(slicelen)

	// Extend its length.
	slicelen = slicelen[:4]
	printSlice(slicelen)

	// Drop its first two values.
	slicelen = slicelen[2:]
	printSlice(slicelen)

	// Nil slices
	fmt.Println("Nil slices")
	var nilslice []int
	fmt.Println(nilslice, len(nilslice), cap(nilslice))
	if nilslice == nil {
		fmt.Println(nil)
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
