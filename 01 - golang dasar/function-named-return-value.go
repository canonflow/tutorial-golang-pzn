package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Nathan"
	middleName = "Garzya"
	lastName = "Santoso"
	return firstName, middleName, lastName
}

func main() {
	firstName, middleName, lastName := getCompleteName()
	fmt.Println(firstName, middleName, lastName) // Nathan Garzya Santoso
}
