package main

import "fmt"

func getFullName() (string, string) {
	return "Nathan", "Garzya"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName) // Nathan Garzya

	// Ignoring some return value
	firstName, _ = getFullName()
	fmt.Println(firstName) // Nathan
}
