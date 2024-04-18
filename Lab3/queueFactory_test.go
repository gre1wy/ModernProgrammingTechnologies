package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestDetermineLimit(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
		err      error
	}{
		{input: "5\n", expected: 5, err: nil},
		{input: "-5\n", expected: 0, err: fmt.Errorf("wrong format of size: entry number>=0")},
		{input: "abc\n", expected: 0, err: fmt.Errorf("wrong format of size: entry number>=0")},
	}

	for _, tc := range testCases {
		reader := strings.NewReader(tc.input)
		limit, err := determineLimit(reader)
		if limit != tc.expected {
			t.Errorf("For input: %s, expected: %d, got: %d", tc.input, tc.expected, limit)
		}
		if (err == nil && tc.err != nil) || (err != nil && tc.err == nil) || (err != nil && err.Error() != tc.err.Error()) {
			t.Errorf("For input: %s, expected error: %v, got error: %v", tc.input, tc.err, err)
		}
	}
}

func TestGetQueue_LinkedListQueue(t *testing.T) {
	input := strings.NewReader("")
	queue, err := getQueue("LinkedListQueue", input)

	if err != nil {
		t.Errorf("Expected no error for LinkedListQueue, got: %v", err)
	}

	if _, ok := queue.(*LinkedListQueue); !ok {
		t.Errorf("Expected LinkedListQueue type, got: %T", queue)
	}
}

func TestGetQueue_CircularArrayQueue(t *testing.T) {
	input := strings.NewReader("5\n")
	queue, err := getQueue("CircularArrayQueue", input)

	if err != nil {
		t.Errorf("Expected no error for CircularArrayQueue, got: %v", err)
	}

	if _, ok := queue.(*CircularArrayQueue); !ok {
		t.Errorf("Expected CircularArrayQueue type, got: %T", queue)
	}
}

func TestGetQueue_WrongQueueType(t *testing.T) {
	input := strings.NewReader("")
	_, err := getQueue("InvalidQueueType", input)

	expectedError := "wrong queue type passed"
	if err == nil || err.Error() != expectedError {
		t.Errorf("Expected error: %s, got: %v", expectedError, err)
	}
}
