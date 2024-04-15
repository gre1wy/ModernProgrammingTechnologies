package main

import (
	"fmt"
)

// Stack represents a basic stack data structure
type Stack struct {
	items []interface{}
}

// Push adds an item to the top of the stack
func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

// Pop removes and returns the item at the top of the stack
func (s *Stack) Pop() interface{} {
	if len(s.items) == 0 {
		return nil
	}
	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item
}

// Peek returns the item at the top of the stack without removing it
func (s *Stack) Peek() interface{} {
	if len(s.items) == 0 {
		return nil
	}
	return s.items[len(s.items)-1]
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of items in the stack
func (s *Stack) Size() int {
	return len(s.items)
}

func main() {
	stack := Stack{}

	// Pushing items onto the stack
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// Popping items from the stack
	fmt.Println(stack.Pop()) // Output: 3
	fmt.Println(stack.Pop()) // Output: 2

	// Peek at the top item without popping it
	fmt.Println(stack.Peek()) // Output: 1

	// Check if the stack is empty
	fmt.Println(stack.IsEmpty()) // Output: false

	// Get the number of items in the stack
	fmt.Println(stack.Size()) // Output: 1
}
