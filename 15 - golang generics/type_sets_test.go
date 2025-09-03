package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Age int

type Number interface {
	~int | int8 | int16 | int32 | int64 | float32 | float64
}

func Min[T Number](first T, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestTypeSets(t *testing.T) {
	assert.Equal(t, int(100), Min[int](100, 200))
	assert.Equal(t, int64(100), Min[int64](100, 200))
	assert.Equal(t, float64(100), Min[float64](100.0, 200.0))

	//Min[string]("NATHAN", "GARZYA") // Error, karena 'string' tidak termasuk di dalam Type Set Number
	/*
		=== RUN   TestTypeSets
		--- PASS: TestTypeSets (0.00s)
		PASS
	*/
}

func TestTypeApproximation(t *testing.T) {
	assert.Equal(t, Age(100), Min[Age](Age(100), Age(200.0))) // Age diperbolehkan, karena memakai Type Approximation untuk tipe dasat 'int'

	/*
		=== RUN   TestTypeApproximation
		--- PASS: TestTypeApproximation (0.00s)
		PASS
	*/
}
