package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome,", name)
	}
}

func main() {
	/* INTRO
	- Sebelumnya setiap membuat function, kita akan selalu memberikan sebuah nama pada function tsb
	- Namun kadang ada kalanya lebih mudah membuat function secara langsung di variable / parameter tanpa
	  harus membuat function terlebih dahulu
	- Hal tersebut dinamakan anonymous function / function tanpa nama
	*/
	blacklist := func(name string) bool {
		return name == "Anjing"
	}
	registerUser("Nathan", blacklist) // Welcome, Nathan
	registerUser("Anjing", func(name string) bool {
		return name == "Anjing"
	}) // You are blocked Anjing
}
