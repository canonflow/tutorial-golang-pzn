package main

import "fmt"

func main() {
	// map[keyType]valueType
	person := map[string]string{
		"name":    "Nathan",
		"address": "Kediri",
	}

	fmt.Println(person)            // map[address:Kediri name:Nathan]
	fmt.Println(person["name"])    // Nathan
	fmt.Println(person["address"]) // Kediri
	fmt.Println(person["salah"])   // karena key "salah" tidak ada, maka akan menggunakan default value (kalo string, itu string kosong. Kalo number ya 0)

	/* ===== FUNCTION =====
	1. len(map) -> mendapatkan jumlah data di map
	2. map[key] -> mengambil data di map dengan key
	3. map[key] = value -> mengubah data di map dengan key
	4. make(map[TypeKey]TypeValue) -> membuat map baru
	5. delete(map, key) -> menghapus data di map dengan key
	*/
	book := make(map[string]string)
	book["title"] = "Buku Go-Lang"
	book["author"] = "Nathan Garzya"
	book["wrong"] = "Ups"
	delete(book, "wrong")
	fmt.Println(book) // map[author:Nathan Garzya title:Buku Go-Lang]
}
