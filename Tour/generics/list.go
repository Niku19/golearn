package main

import (
	"fmt"
)

// List represents a singly-linked list that holds values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

// PushFront adds a new element to the front of the list.
// Returns a new list with the new head.
func PushFront[T any](head *List[T], val T) *List[T] {
	return &List[T]{val: val, next: head}
}

// PushBack adds a new element to the end of the list.
// Returns the (unchanged) head of the list.
func PushBack[T any](head *List[T], val T) *List[T] {
	if head == nil {
		return &List[T]{val: val}
	}
	curr := head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = &List[T]{val: val}
	return head
}

// PrintList prints all elements in the list.
func PrintList[T any](head *List[T]) {
	for curr := head; curr != nil; curr = curr.next {
		fmt.Printf("%v -> ", curr.val)
	}
	fmt.Println("nil")
}

// Length returns the number of elements in the list.
func Length[T any](head *List[T]) int {
	count := 0
	for curr := head; curr != nil; curr = curr.next {
		count++
	}
	return count
}

// Contains checks if the list contains a specific value.
func Contains[T comparable](head *List[T], val T) bool {
	for curr := head; curr != nil; curr = curr.next {
		if curr.val == val {
			return true
		}
	}
	return false
}

func main() {
	var list *List[int]

	list = PushFront(list, 3)
	list = PushFront(list, 2)
	list = PushBack(list, 4)
	list = PushFront(list, 1)

	PrintList(list)
	fmt.Println("Length:", Length(list))
	fmt.Println("Contains 3?", Contains(list, 3))
	fmt.Println("Contains 5?", Contains(list, 5))
}
