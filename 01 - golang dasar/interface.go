package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

func SayHello(hasName HasName) {
	fmt.Println("Hello,", hasName.GetName())
}

func main() {
	/* ===== INTRO =====
	- Interface adalah tipe data abstract, dia tidak memiliki implementasi langsung
	- Sebuah interface biasanya berisikan definisi - definisi method
	- Biasanya digunakan sebagai KONTRAK
	*/

	/* ===== IMPLEMENTASI =====
	- Setiap tipe data yang SESUAI dengan kontrak interface, secara OTOMATIS dianggap sbg interface tsb
	- Sehingga kita TIDAK PERLU mengimplementasikan interface secara manual
	- Hal ini berbeda dengan bahasa lain dimana ketika membuat interface, kita harus menyebutkan secara
	  eksplisit akan menggunakan interface mana
	*/
	p := Person{"Nathan"}
	animal := Animal{"Lion"}
	SayHello(p)      // Hello, Nathan
	SayHello(animal) // Hello, Lion
}
