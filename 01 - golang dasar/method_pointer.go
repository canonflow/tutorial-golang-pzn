package main

import "fmt"

type Man struct {
	Name string
}

func (man Man) Married() {
	man.Name = "Mr. " + man.Name
}

func (man *Man) MarriedByPointer() {
	man.Name = "Mr. " + man.Name
}

func main() {
	/* ===== INTRO =====
	- Walaupun method akan menempel di struct, tapi sebenarnya data struct yang diakses di dalam method
	  adalah PASS BY VALUE
	- Sangat direkomendasikan menggunakan pointer di method, sehingga tidak boros memory karena
	  harus selalu duplikasi ketika memanggil method
	*/
	nathan := Man{Name: "Nathan"}
	nathan.Married()
	fmt.Println(nathan) // {Nathan}

	nathan.MarriedByPointer()
	fmt.Println(nathan) // {Mr. Nathan}
}
