package main

import (
	"reflect"
	"testing"
)

func TestLinkedList_Constructor(t *testing.T) {
	queue := NewLinkedListQueue()

	expected := &LinkedListQueue{
		front: nil,
		rear:  nil,
		size:  0,
	}

	if !reflect.DeepEqual(queue, expected) {
		t.Errorf("Expected queue state: %+v, but got: %+v", expected, queue)
	}
}
func TestLinkedListQueue_Enqueue(t *testing.T) {
	q := NewLinkedListQueue()
	q.Enqueue(1)

	expected := &LinkedListQueue{
		front: q.front,
		rear:  q.rear,
		size:  1,
	}

	if !reflect.DeepEqual(q, expected) {
		t.Errorf("Expected queue state: %+v, but got: %+v", expected, q)
	}
}

func TestLinkedListQueue_Dequeue(t *testing.T) {
	q := NewLinkedListQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	item := q.Dequeue()
	if item != 1 {
		t.Errorf("Expected dequeued item to be 1, but got %v", item)
	}

	if q.Size() != 2 {
		t.Errorf("Expected queue size to be 2 after dequeue, but got %d", q.Size())
	}
}

func TestLinkedListQueue_IsEmpty(t *testing.T) {
	q := NewLinkedListQueue()

	if !q.IsEmpty() {
		t.Error("Expected queue to be empty")
	}

	q.Enqueue(1)
	if q.IsEmpty() {
		t.Error("Expected queue not to be empty after enqueue")
	}
}

func TestLinkedListQueue_Front(t *testing.T) {
	q := NewLinkedListQueue()
	q.Enqueue(1)
	q.Enqueue(2)

	front := q.Front()
	if front != 1 {
		t.Errorf("Expected front item to be 1, but got %v", front)
	}
}

func TestLinkedListQueue_Rear(t *testing.T) {
	q := NewLinkedListQueue()
	q.Enqueue(1)
	q.Enqueue(2)

	rear := q.Rear()
	if rear != 2 {
		t.Errorf("Expected rear item to be 2, but got %v", rear)
	}
}

func TestLinkedListQueue_Size(t *testing.T) {
	q := NewLinkedListQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	size := q.Size()
	if size != 3 {
		t.Errorf("Expected queue size to be 3, but got %d", size)
	}
}

func TestLinkedListQueue_Display(t *testing.T) {
	q, _ := NewCircularArrayQueue(3)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	done := capture()
	q.Display()
	capturedOutput, err := done()

	expected := "Current Queue:\n1 2 3 \n"

	if capturedOutput != expected {
		t.Errorf("Expected output: %q, but got: %q", expected, capturedOutput)
	}

	if err != nil {
		t.Errorf("Error while capturing output: %v", err)
	}
}
