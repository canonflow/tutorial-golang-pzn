package main

import "fmt"

func sumAll(numbers ...int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}

	return sum
}

func main() {
	sum := sumAll(1, 2, 3, 4, 5)
	fmt.Println(sum) // 15

	// Slice as Argument
	numbers := []int{1, 2, 3, 4, 5}
	total := sumAll(numbers...)
	fmt.Println(total) // 15
}
