package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	/* FUNCTION VALUE
	1. function adalah first class citizen
	2. function juga merupakan tipe data, dan bisa disimpan di dalam variable
	*/
	goodbye := getGoodBye
	fmt.Println(goodbye("Nathan")) // Good Bye Nathan
}
