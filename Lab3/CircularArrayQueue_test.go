package main

import (
	"reflect"
	"testing"
)

func TestCircularArrayQueue_Constructor(t *testing.T) {
	capacity := 5
	queue, err := NewCircularArrayQueue(capacity)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
		return
	}

	expected := &CircularArrayQueue{
		items:    make([]interface{}, capacity),
		size:     0,
		capacity: capacity,
		head:     0,
		tail:     0,
	}

	if !reflect.DeepEqual(queue, expected) {
		t.Errorf("Expected queue state: %+v, but got: %+v", expected, queue)
	}
}

func TestCircularArrayQueue_Constructor_NegativeCapacity(t *testing.T) {
	capacity := -5
	_, err := NewCircularArrayQueue(capacity)

	expectedError := "capacity must be non-negative"

	if err == nil {
		t.Error("Expected error, but got nil")
		return
	}
	if err.Error() != expectedError {
		t.Errorf("Expected error message: %s, but got: %s", expectedError, err.Error())
	}
}

func TestCircularArrayQueue_Initialization(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
		elements []interface{}
		expected CircularArrayQueue
	}{
		{
			name:     "WithElements",
			capacity: 5,
			elements: []interface{}{1, 2, 3},
			expected: CircularArrayQueue{
				items:    []interface{}{1, 2, 3, nil, nil},
				size:     3,
				capacity: 5,
				head:     0,
				tail:     3,
			},
		},
		{
			name:     "Empty",
			capacity: 5,
			elements: nil,
			expected: CircularArrayQueue{
				items:    []interface{}{nil, nil, nil, nil, nil},
				size:     0,
				capacity: 5,
				head:     0,
				tail:     0,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			q, _ := NewCircularArrayQueue(test.capacity)
			for _, elem := range test.elements {
				q.Enqueue(elem)
			}
			if !reflect.DeepEqual(*q, test.expected) {
				t.Errorf("Expected queue state: %+v, but got: %+v", test.expected, *q)
			}
		})
	}
}
func TestCircularArrayQueue_Enqueue(t *testing.T) {
	q, _ := NewCircularArrayQueue(5)
	q.Enqueue(1)
	q.Enqueue("2")
	q.Enqueue("goat")

	expected := &CircularArrayQueue{
		items:    []interface{}{1, "2", "goat", nil, nil},
		size:     3,
		capacity: 5,
		head:     0,
		tail:     3,
	}

	if !reflect.DeepEqual(q, expected) {
		t.Errorf("Expected queue state: %+v, but got: %+v", expected, q)
	}
}

func TestCircularArrayQueue_Dequeue(t *testing.T) {
	q, _ := NewCircularArrayQueue(5)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Dequeue()
	q.Enqueue(4)
	q.Enqueue(5)
	q.Enqueue(6)
	q.Dequeue()
	q.Enqueue(7)
	q.Dequeue()

	expected := &CircularArrayQueue{
		items:    []interface{}{6, 7, 3, 4, 5},
		size:     4,
		capacity: 5,
		head:     3,
		tail:     2,
	}

	if !reflect.DeepEqual(q, expected) {
		t.Errorf("Expected queue state: %+v, but got: %+v", expected, q)
	}

}

func TestCircularArrayQueue_IsEmpty(t *testing.T) {
	q, _ := NewCircularArrayQueue(5)
	if !q.IsEmpty() {
		t.Error("Expected queue to be empty")
	}
	q.Enqueue(1)
	if q.IsEmpty() {
		t.Error("Expected queue not to be empty after enqueue")
	}
}

func TestCircularArrayQueue_IsFull(t *testing.T) {
	q, _ := NewCircularArrayQueue(3)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	if !q.IsFull() {
		t.Error("Expected queue to be full")
	}
}

func TestCircularArrayQueue_Front(t *testing.T) {
	q, _ := NewCircularArrayQueue(3)
	q.Enqueue(1)
	q.Enqueue(2)
	front := q.Front()
	if front != 1 {
		t.Errorf("Expected front item to be 1, but got %v", front)
	}
}

func TestCircularArrayQueue_Rear(t *testing.T) {
	q, _ := NewCircularArrayQueue(3)
	q.Enqueue('h')
	q.Enqueue(2)
	rear := q.Rear()
	if rear != 2 {
		t.Errorf("Expected rear item to be 2, but got %v", rear)
	}
}

func TestCircularArrayQueue_Display(t *testing.T) {
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
