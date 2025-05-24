package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name    string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
	Email   string `required:"true" max:"10"`
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value) // Mengambil tipe-nya
	fmt.Println("Type Name", valueType.Name())

	// Ambil Struct Field
	for i := 0; i < valueType.NumField(); i++ {
		var structField reflect.StructField = valueType.Field(i) // Mengambil field-nya
		fmt.Println(structField.Name, "with type", structField.Type)

		// Ambil Tag dari Field
		fmt.Println(structField.Tag.Get("required"), structField.Tag.Get("max"))
	}
}

func IsValid(value any) (result bool) {
	result = true
	t := reflect.TypeOf(value)

	// Iterasi tiep fieldnya
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			// Ambil datanya
			data := reflect.ValueOf(value).Field(i).Interface()

			// Karena required
			//return data != ""
			result = data != ""
			if !result {
				return result
			}
		}
	}

	return result
}

func main() {
	/* ===== INTRO =====
	- Dalam bahasa pemrograman, biasanya ada fitur Reflection, dimana kita bisa melihat struktur kode kita pada saat
	  aplikasi sedang berjalan
	- Hal ini bisa dilakukan di golang dengan menggunakan package Reflect
	- Fitur ini mungkin tidak bisa dibahas secara lengkap dalam 1 video, Anda bisa eksplorasi package reflec ini secara
	  otodidak
	- Reflection sangat berguna ketika kita INGIN MEMBUAT LIBRARY yg general sehingga mudah digunakan
	*/
	readField(Sample{"Nathan"})
	/*
		Type Name Sample
		Name with type string
		true 10
	*/
	readField(Person{"Nathan", "Indonesia", "abc@gmail.com"})
	/*
		Type Name Person
		Name with type string
		true 10
		Address with type string
		true 10
		Email with type string
		true 10
	*/

	person1 := Person{
		Name:    "Nathan",
		Address: "Indonesia",
		Email:   "nathan@gmail.com",
	}
	person2 := Person{
		Name:    "Nathan",
		Address: "Indonesia",
		Email:   "",
	}
	fmt.Println(IsValid(person1)) // true
	fmt.Println(IsValid(person2)) // false
}
