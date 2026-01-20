package main

import "fmt"

// List represents a singly-linked list that holds values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

// Add a new element at the beginning
func (l *List[T]) Add(value T) *List[T] {
	return &List[T]{next: l, val: value}
}

// Append element at the end
func (l *List[T]) Append(value T) {
	current := l
	for current.next != nil {
		current = current.next
	}
	current.next = &List[T]{val: value}
}

// Find checks if value exists in list
func (l *List[T]) Find(value T, equal func(a, b T) bool) bool {
	for current := l; current != nil; current = current.next {
		if equal(current.val, value) {
			return true
		}
	}
	return false
}

// Length of list
func (l *List[T]) Length() int {
	count := 0
	for current := l; current != nil; current = current.next {
		count++
	}
	return count
}

// Print list
func (l *List[T]) Print() {
	for current := l; current != nil; current = current.next {
		fmt.Print(current.val, " -> ")
	}
	fmt.Println("nil")
}

func main() {
	// Create list of integers
	var list *List[int]
	list = list.Add(10)
	list = list.Add(20)
	list = list.Add(30)

	list.Append(5)

	fmt.Print("List: ")
	list.Print()

	fmt.Println("Length:", list.Length())

	// Find example
	found := list.Find(20, func(a, b int) bool {
		return a == b
	})

	fmt.Println("20 present?:", found)
}
