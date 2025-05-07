package main

import "fmt"

func main() {
	name := "Nath"
	if name == "Nathan" {
		fmt.Println("Hello, Nathan!")
	} else if name == "Garzya" {
		fmt.Println("Hello, Joko!")
	} else {
		fmt.Println("Hello, Boleh Kenalan?")
	}

	// ===== IF SHORT STATEMENT =====
	if length := len(name); length > 5 {
		fmt.Println("name is too long")
	} else {
		fmt.Println("name is correct")
	}
}
