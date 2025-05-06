package main

import "fmt"

func main() {
	var name1 = "nathan"
	var name2 = "nathan"

	var result1 bool = name1 == name2
	var result2 bool = name1 != name2

	fmt.Println(result1) // true
	fmt.Println(result2) // false

	// Selain == dan !=, akan membandingkan per-alphabet
	var a = "a"
	var b = "b"
	var result3 bool = b > a
	fmt.Println(result3) // true
}
