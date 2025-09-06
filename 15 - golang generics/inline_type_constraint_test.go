package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func FindMin[T interface{ int | int64 | float64 }](first T, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestFindMin(t *testing.T) {
	assert.Equal(t, 100, FindMin(100, 200))
	assert.Equal(t, int64(100), FindMin(int64(100), int64(200)))
	assert.Equal(t, 100.0, FindMin(100.0, 200.0))

	/*
		=== RUN   TestFindMin
		--- PASS: TestFindMin (0.00s)
		PASS
	*/
}

// Generic Type di Parameter
func GetFirst[T []E, E any](data T) E {
	first := data[0]
	return first
}

func TestGetFirst(t *testing.T) {
	names := []string{
		"Nathan",
		"Garzya",
		"Santoso",
	}

	first := GetFirst[[]string, string](names)
	assert.Equal(t, "Nathan", first)
	/*
		=== RUN   TestGetFirst
		--- PASS: TestGetFirst (0.00s)
		PASS
	*/
}
