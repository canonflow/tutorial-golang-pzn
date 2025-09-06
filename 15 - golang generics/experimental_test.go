package main

import (
	"maps"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
)

func ExperimentalMin[T constraints.Ordered](first T, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestExperimentalMin(t *testing.T) {
	assert.Equal(t, 100, ExperimentalMin(100, 200))
	assert.Equal(t, 100.0, ExperimentalMin(100.0, 200.0))

	/*
		=== RUN   TestExperimentalMin
		--- PASS: TestExperimentalMin (0.00s)
		PASS
	*/
}

func TestExperimentalMaps(t *testing.T) {
	first := map[string]string{
		"Name": "Nathan",
	}

	second := map[string]string{
		"Name": "Nathan",
	}

	assert.True(t, maps.Equal(first, second))

	/*
		=== RUN   TestExperimentalMaps
		--- PASS: TestExperimentalMaps (0.00s)
		PASS
	*/
}

func TestExperimentalSlices(t *testing.T) {
	first := []string{"Nathan"}
	second := []string{"Nathan"}

	assert.True(t, slices.Equal(first, second))

	/*
		=== RUN   TestExperimentalSlices
		--- PASS: TestExperimentalSlices (0.00s)
		PASS
	*/
}
