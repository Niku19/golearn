package main

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	first, second := 0, 1
	return func() int {
		currentFib := first
		first, second = second, first+second
		return currentFib
	}
}
