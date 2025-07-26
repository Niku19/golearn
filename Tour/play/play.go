package main

import (
	"fmt"
	"sort" // For sorting map keys for consistent output
	"sync"
	"time"
)

// SafeCounter is a concurrency-safe counter for integer values.
type SafeCounter struct {
	mu     sync.Mutex
	counts map[int]int
}

// Inc increments the count for the given value.
func (c *SafeCounter) Inc(val int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counts[val]++
}

// GetCounts returns a copy of the counts map.
func (c *SafeCounter) GetCounts() map[int]int {
	c.mu.Lock()
	defer c.mu.Unlock()
	// Return a copy to prevent external modification without mutex
	copyMap := make(map[int]int)
	for k, v := range c.counts {
		copyMap[k] = v
	}
	return copyMap
}

func main() {
	numGoroutines := 100 // Number of goroutines to launch

	// --- DEMONSTRATION OF THE BUG ---
	fmt.Println("--- BUGGY EXAMPLE (Loop Variable Capture) ---")
	fmt.Println("Expected: 0, 1, 2, ... 99 (each counted once)")
	fmt.Println("Actual (likely): Many '99's, or other final/later values, counted multiple times")

	var wgBuggy sync.WaitGroup
	buggyCounter := &SafeCounter{counts: make(map[int]int)} // Initialize safe counter

	for i := 0; i < numGoroutines; i++ {
		wgBuggy.Add(1)
		go func() { // The anonymous function forms a closure over 'i'
			defer wgBuggy.Done()
			time.Sleep(50 * time.Millisecond) // Simulate some work
			fmt.Printf("Buggy Goroutine: Value of i = %d\n", i)
			buggyCounter.Inc(i) // Increment count for the captured 'i'
		}()
	}

	wgBuggy.Wait() // Wait for all buggy goroutines to finish

	fmt.Println("\nBuggy Counts:")
	printSortedCounts(buggyCounter.GetCounts())
	fmt.Println("--- BUGGY EXAMPLE FINISHED ---\n")

	// --- DEMONSTRATION OF THE CORRECT FIX ---
	fmt.Println("--- CORRECTED EXAMPLE (Pass by Value to Closure) ---")
	fmt.Println("Expected: 0, 1, 2, ... 99 (each counted once)")
	fmt.Println("Actual: 0, 1, 2, ... 99 (each counted once)")

	var wgCorrect sync.WaitGroup
	correctCounter := &SafeCounter{counts: make(map[int]int)} // Initialize another safe counter

	for i := 0; i < numGoroutines; i++ {
		wgCorrect.Add(1)
		// FIX: Pass the loop variable 'i' as an argument to the anonymous function.
		go func(val int) { // 'val' is a new variable for each goroutine
			defer wgCorrect.Done()
			time.Sleep(50 * time.Millisecond) // Same delay
			fmt.Printf("Correct Goroutine: Value of val = %d\n", val)
			correctCounter.Inc(val) // Increment count for the specific 'val'
		}(i) // Pass 'i' by value
	}

	wgCorrect.Wait() // Wait for all corrected goroutines to finish

	fmt.Println("\nCorrected Counts:")
	printSortedCounts(correctCounter.GetCounts())
	fmt.Println("--- CORRECTED EXAMPLE FINISHED ---")
}

// printSortedCounts helps visualize the map contents in a consistent order.
func printSortedCounts(counts map[int]int) {
	keys := make([]int, 0, len(counts))
	for k := range counts {
		keys = append(keys, k)
	}
	sort.Ints(keys) // Sort the keys

	for _, k := range keys {
		fmt.Printf("  Value %d appeared %d time(s)\n", k, counts[k])
	}
}
