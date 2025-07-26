package main

import (
	"fmt"
)

// walkRecursive is a helper function that performs an in-order traversal
// of the binary tree and sends each node's value to the channel.
// It's called recursively for left, then current, then right subtrees.
func walkRecursive(t *Tree, ch chan int) {
	// Base case: if the current node is nil, return.
	if t == nil {
		return
	}

	// 1. Recursively walk the left subtree.
	walkRecursive(t.Left, ch)

	// 2. Send the current node's value to the channel.
	ch <- t.Value

	// 3. Recursively walk the right subtree.
	walkRecursive(t.Right, ch)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
// This function starts the recursive traversal in a goroutine
// and ensures the channel is closed once all values have been sent.
func Walk(t *Tree, ch chan int) {
	// defer close(ch) ensures that the channel 'ch' is closed
	// after the walkRecursive function completes its execution.
	// This is crucial so that receivers using 'range' over the channel
	// know when there are no more values to expect.
	defer close(ch)
	walkRecursive(t, ch)
}

// Same determines whether the trees t1 and t2 contain the same values.
// It does this by concurrently walking both trees and comparing the values
// received from their respective channels.
func Same(t1, t2 *Tree) bool {
	// Create two unbuffered channels to receive values from the tree traversals.
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Start two goroutines to walk each tree concurrently.
	// Each goroutine will send its tree's values to its respective channel.
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	// Iterate 10 times because tree.New(k) is specified to construct
	// trees holding values k, 2k, ..., 10k, meaning exactly 10 values.
	for i := 0; i < 10; i++ {
		// Receive a value from each channel.
		// 'ok' will be false if the channel has been closed and no more values are available.
		val1, ok1 := <-ch1
		val2, ok2 := <-ch2

		// Compare the received values and the 'ok' status.
		// If values are different (val1 != val2) OR
		// if one channel closed but the other didn't (ok1 != ok2),
		// then the trees are not "Same".
		if val1 != val2 || ok1 != ok2 {
			return false
		}
	}

	// If the loop completes without returning false, it means all 10 values
	// from both trees matched, and both channels were still open (or closed simultaneously
	// after the 10th value, which is implicitly handled by the ok check).
	// Therefore, the trees are considered "Same".
	return true
}

func main() {
	// --- 2. Test the Walk function ---
	fmt.Println("--- Testing Walk function ---")
	// Create a channel for the Walk function to send values to.
	chWalk := make(chan int)
	var tree Tree

	// Start a goroutine to walk a tree generated with k=1 (values 1, 2, ..., 10).
	go Walk(tree.New(1), chWalk)

	// Read and print 10 values from the channel.
	// This loop will correctly terminate because the 'Walk' function
	// defers the closing of 'chWalk'.
	fmt.Println("Values from tree.New(1) via Walk:")
	for i := 0; i < 10; i++ {
		fmt.Println(<-chWalk)
	}

	// --- 4. Test the Same function ---
	fmt.Println("\n--- Testing Same function ---")

	// Test case 1: Two identical trees (tree.New(1) vs tree.New(1))
	// Expected: true
	fmt.Println("Same(tree.New(1), tree.New(1)):", Same(tree.New(1), tree.New(1)))

	// Test case 2: Two different trees (tree.New(1) vs tree.New(2))
	// Expected: false (values will be 1,2,...10 vs 2,4,...20)
	fmt.Println("Same(tree.New(1), tree.New(2)):", Same(tree.New(1), tree.New(2)))

	// Test case 3: Two identical trees with a different base (tree.New(2) vs tree.New(2))
	// Expected: true
	fmt.Println("Same(tree.New(2), tree.New(2)):", Same(tree.New(2), tree.New(2)))
}
