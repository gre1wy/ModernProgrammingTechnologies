package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayPassedByReference(t *testing.T) {
	value := []int{1, 2, 3}

	IncrementEach(value)

	expected := []int{2, 3, 4}

	assert.Equal(t, expected, value, "Expected incremented value")

}

func IncrementEach(array []int) {
	for i := range array {
		array[i]++
	}
}

func TestCustomClassPassedByReference(t *testing.T) {
	customClass := CustomStructReference{Value: 10}

	IncrementReference(&customClass)

	assert.Equal(t, 11, customClass.Value, "Expected incremented value")
}

type CustomStructReference struct {
	Value int
}

func (c *CustomStructReference) Increment() {
	c.Value++
}

func IncrementReference(c *CustomStructReference) {
	c.Increment()
}
