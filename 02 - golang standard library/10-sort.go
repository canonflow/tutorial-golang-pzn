package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (userSlice UserSlice) Len() int {
	return len(userSlice)
}

func (userSlice UserSlice) Less(i, j int) bool {
	return userSlice[i].Age < userSlice[j].Age
}

func (userSlice UserSlice) Swap(i, j int) {
	//temp := userSlice[i]
	//userSlice[i] = userSlice[j]
	//userSlice[j] = temp
	userSlice[i], userSlice[j] = userSlice[j], userSlice[i]
}

func main() {
	/* ===== INTRO =====
	- Package sort adalah package yg berisikan utilitas utk proses pengurutan
	- Agar data kita bisa diurutkan, kita harus mengimplementasikan kontrak di interface sort.Interface

	sort.Interface {
		Len() int
		Less(i, j int) bool
		Swap(i, j int) int
	}
	*/

	users := UserSlice{
		User{"Nathan", 30},
		User{"Garzya", 35},
		User{"Santoso", 25},
		User{"Canonflow", 23},
	}

	sort.Sort(users)
	fmt.Println(users) // [{Canonflow 23} {Santoso 25} {Nathan 30} {Garzya 35}]
}
