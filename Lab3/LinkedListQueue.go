package main

import (
	"fmt"
)

// Node представляє вузол для звязного списку
type Node struct {
	data any
	next *Node
}

// LinkedListQueue представляє чергу на основі звязного списку
type LinkedListQueue struct {
	front *Node
	rear  *Node
	size  int
}

// NewLinkedListQueue функція генератор
func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{}
}

// Enqueue додає елемент в кінець черги
func (q *LinkedListQueue) Enqueue(item interface{}) {
	newNode := &Node{data: item, next: nil}

	if q.rear == nil {
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.next = newNode
		q.rear = newNode
	}
	q.size++
}

// Dequeue видаляє и повертає елемент з початку черги
func (q *LinkedListQueue) Dequeue() interface{} {
	if q.IsEmpty() {
		fmt.Println("Queue is empty")
		return nil
	}
	item := q.front.data
	q.front = q.front.next
	if q.front == nil {
		q.rear = nil
	}
	q.size--
	return item
}

// IsEmpty повертає true, якщо черга пуста
func (q *LinkedListQueue) IsEmpty() bool {
	return q.size == 0
}

// Size повертає поточний розмір черги
func (q *LinkedListQueue) Size() int {
	return q.size
}

// Front повертає елемент в початку черги без його видалення
func (q *LinkedListQueue) Front() interface{} {
	if q.IsEmpty() {
		fmt.Println("Queue is empty")
		return nil
	}
	return q.front.data
}

// Rear повертає елемент в кінці черги без його видалення
func (q *LinkedListQueue) Rear() interface{} {
	if q.IsEmpty() {
		fmt.Println("Queue is empty")
		return nil
	}
	return q.rear.data
}

// Display відображає поточний стан черги
func (q *LinkedListQueue) Display() {
	if q.IsEmpty() {
		fmt.Println("Queue is empty")
		return
	}

	current := q.front
	fmt.Print("Current Queue: ")
	for current != nil {
		fmt.Printf("%v ", current.data)
		current = current.next
	}
	fmt.Println()
}
