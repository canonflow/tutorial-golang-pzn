package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Kalau pake any, gk bisa soalnya gk semua tipe data any bisa dibandingkan
func IsSame[T comparable](val1, val2 T) bool {
	if val1 == val2 {
		return true
	} else {
		return false
	}
}

func TestComparable(t *testing.T) {
	assert.Equal(t, true, IsSame[string]("nathan", "nathan"))
	assert.Equal(t, true, IsSame[int](100, 100))
	/*
		=== RUN   TestComparable
		--- PASS: TestComparable (0.00s)
		PASS
	*/
}
