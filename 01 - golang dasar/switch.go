package main

import "fmt"

func main() {
	name := "Nathan"
	switch name {
	case "Nathan":
		fmt.Println("Hello, Nathan!")
	case "Garzya":
		fmt.Println("Hello, Garzya!")
	default:
		fmt.Println("Hello, Boleh Kenalan?")
	}

	// ===== SWITCH SHORT STATEMENT =====
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("name is too long!")
	case false:
		fmt.Println("name is correct!")
	}

	// ===== SWITCH WITHOUT CONDITION =====
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("name is too long!")
	case length > 5:
		fmt.Println("name is slightly long!")
	default:
		fmt.Println("name is correct!")
	}
}
