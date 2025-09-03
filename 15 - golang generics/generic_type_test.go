package main

import (
	"fmt"
	"testing"
)

type Bag[T any] []T

func PrintBag[T any](bag Bag[T]) {
	for _, val := range bag {
		fmt.Println(val)
	}
}

func TestBag(t *testing.T) {
	numbers := Bag[int]{1, 2, 3, 4, 5}
	PrintBag[int](numbers)

	names := Bag[string]{"Nathan", "Garzya", "santoso"}
	fmt.Println(names)
	PrintBag[string](names)

	/*
		=== RUN   TestBag
		1
		2
		3
		4
		5
		[Nathan Garzya santoso]
		Nathan
		Garzya
		santoso
		--- PASS: TestBag (0.00s)
		PASS
	*/
}
