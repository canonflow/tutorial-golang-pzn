package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter func(string) string) {
	fmt.Println("Hello", filter(name))
}

func sayHelloWithFilter2(name string, filter Filter) {
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	}
	return name
}

func main() {
	/* Function as Parameter
	- Function bisa digunakan sebagai parameter untuk function lain
	*/
	sayHelloWithFilter("Nathan", spamFilter) // Hello Nathan

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter) // Hello ...

	/* Function Type Declaration
	- Kadang jika function terlalu panjang, agak ribet untuk menuliskannya di dalam parameter
	- Type Declaration juga bisa digunakan untuk membuat alias function, sehingga akan mempermudah kita menggunakan function sebagai parameter
	*/

	sayHelloWithFilter2("Garzya", spamFilter) // Hello Garzya
}
