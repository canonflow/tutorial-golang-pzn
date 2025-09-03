package main

import (
	"fmt"
	"testing"
)

func MultipleParameter[T1 any, T2 any](param1 T1, param2 T2) {
	fmt.Println(param1)
	fmt.Println(param2)
}

func TestMultipleParameter(t *testing.T) {
	MultipleParameter[string, int]("Nathan", 100)
	MultipleParameter[int, bool](10, true)
	/*
		=== RUN   TestMultipleParameter
		Nathan
		100
		10
		true
		--- PASS: TestMultipleParameter (0.00s)
		PASS
	*/
}
