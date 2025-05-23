package main

import "fmt"

type Player struct {
	Name string
}

func (player Player) sayHello(name string) {
	fmt.Println("Hello", name, "! My name is", player.Name)
}

func main() {
	/* ===== INTRO =====
	- Struct bisa digunakan sebagai parameter
	- Jika kita ingin menambahkan method ke dalam structs, sehingga seakan - akan sebuah struct memiliki function
	- method adalah function yg menempel di dalam structs
	*/
	budi := Player{"Budi"}
	budi.sayHello("Agus") // Hello Agus ! My name is Budi
 
}
