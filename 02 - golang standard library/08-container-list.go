package main

import (
	"container/list"
	"fmt"
)

func main() {
	/* ===== INTRO =====
	- Package container/list adalah implementasi struktur data DOUBLE LINKED LIST di golang
	*/

	//data := list.New()
	var data *list.List = list.New()
	data.PushBack("Nathan")
	data.PushBack("Garzya")
	data.PushBack("Santoso")

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
