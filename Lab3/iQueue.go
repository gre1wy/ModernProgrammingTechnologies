package main

type Queue interface {
	Enqueue(item interface{})
	Dequeue() interface{}
	IsEmpty() bool
	Size() int
	Front() interface{}
	Rear() interface{}
	Display()
}
