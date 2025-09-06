package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type GetterSetter[T any] interface {
	GetValue() T
	SetValue(value T)
}

func ChangeValue[T any](param GetterSetter[T], newValue T) T {
	param.SetValue(newValue)
	return param.GetValue()
}

type MyData[T any] struct {
	Value T
}

func (m *MyData[T]) GetValue() T {
	return m.Value
}

func (m *MyData[T]) SetValue(value T) {
	m.Value = value
}

func TestInterface(t *testing.T) {
	myData := MyData[string]{}
	result := ChangeValue(&myData, "Nathan")

	assert.Equal(t, "Nathan", result)
	/*
		=== RUN   TestInterface
		--- PASS: TestInterface (0.00s)
		PASS
	*/
}
