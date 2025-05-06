package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var d = 5

	var e = 2
	var c = a + b - d*e

	fmt.Println(c) // 10

	// ===== AUGMENTED ASSIGNMENT =====
	var i = 10
	i += 10
	fmt.Println(i) // 20

	i += 5
	fmt.Println(i) // 25

	// ===== UNARY OPERATOR =====
	var j = 1
	j++
	j++
	fmt.Println(j) // 3

	j--
	j--
	fmt.Println(j) // 1
}
