package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := range 10 {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	// calling fibonacci
	f := fibonacci()
	fmt.Printf("Type of f is %T\n", f)
	for range 10 {
		fmt.Printf("%v ", f())
	}
}
