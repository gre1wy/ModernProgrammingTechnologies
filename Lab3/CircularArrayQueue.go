package main

import (
	"errors"
	"fmt"
)

type CircularArrayQueue struct {
	items    []interface{}
	size     int // Теперішній розмір
	capacity int // Максимальний розмір
	head     int // Індекс початку черги
	tail     int // Індекс кінця черги
}

// NewCircularArrayQueue функція конструктор
func NewCircularArrayQueue(capacity int) (*CircularArrayQueue, error) {
	if capacity < 0 {
		return nil, errors.New("capacity must be non-negative")
	}
	return &CircularArrayQueue{
		items:    make([]interface{}, capacity),
		capacity: capacity,
		size:     0,
		head:     0,
		tail:     0,
	}, nil
}

// Enqueue додає елемент в кінець черги
func (q *CircularArrayQueue) Enqueue(item interface{}) {
	if q.IsFull() {
		fmt.Println("Queue is full")
		return
	}
	q.items[q.tail] = item
	q.tail = (q.tail + 1) % q.capacity
	q.size++
}

// Dequeue "видаляє" и повертає елемент з початку черги
func (q *CircularArrayQueue) Dequeue() interface{} {
	if q.IsEmpty() {
		fmt.Println("Queue is empty")
		return nil
	}
	item := q.items[q.head]
	q.head = (q.head + 1) % q.capacity
	q.size--
	return item
}

// IsEmpty повертає true, якщо черга пуста
func (q *CircularArrayQueue) IsEmpty() bool {
	return q.size == 0
}

// IsFull повертає true, якшо черга заповнена
func (q *CircularArrayQueue) IsFull() bool {
	return q.size == q.capacity
}

// Size повертає поточний розмір черги
func (q *CircularArrayQueue) Size() int {
	return q.size
}

// Front повертає елемент в початку черги без його видалення
func (q *CircularArrayQueue) Front() interface{} {
	if q.IsEmpty() {
		fmt.Println("Queue is empty")
		return nil
	}
	return q.items[q.head]
}

// Rear повертає елемент в кінці черги без його видалення
func (q *CircularArrayQueue) Rear() interface{} {
	if q.IsEmpty() {
		fmt.Println("Queue is empty")
		return nil
	}
	return q.items[(q.tail-1+q.capacity)%q.capacity]
}

// Display відображає поточний стан черги
func (q *CircularArrayQueue) Display() {
	if q.IsEmpty() {
		fmt.Println("Queue is empty")
		return
	}
	fmt.Println("Current Queue:")
	i := q.head
	for {
		fmt.Printf("%v ", q.items[i])
		i = (i + 1) % q.capacity
		if i == q.tail {
			break
		}
	}
	fmt.Println()
}
