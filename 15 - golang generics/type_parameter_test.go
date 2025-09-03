package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Length[T any](params T) T {
	fmt.Println(params)
	return params
}

func TestLength(t *testing.T) {
	var result string = Length[string]("Nathan")
	assert.Equal(t, "Nathan", result)

	var resultNumber int = Length[int](100)
	assert.Equal(t, 100, resultNumber)
	/*
		=== RUN   TestLength
		Nathan
		100
		--- PASS: TestLength (0.00s)
		PASS
	*/
}

func TestSample(t *testing.T) {
	assert.True(t, true)
}
