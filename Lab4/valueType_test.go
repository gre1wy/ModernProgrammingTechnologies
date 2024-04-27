package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type CustomStructValue struct {
	Value int
}

func NewCustomStruct(value int) CustomStructValue {
	return CustomStructValue{Value: value}
}

func TestIntegerPassedByValue(t *testing.T) {
	value := 10

	IncrementValue(value)

	assert.Equal(t, 10, value, "Expected not incremented value")
}

func IncrementValue(v int) {
	v++
}

func TestCustomStructPassedByValue(t *testing.T) {
	customStruct := NewCustomStruct(10)

	IncrementStruct(customStruct)

	assert.Equal(t, 10, customStruct.Value, "Expected not incremented value")
}

func IncrementStruct(s CustomStructValue) {
	s.Value++
}
