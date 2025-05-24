package main

import (
	"fmt"
	"os"
)

func main() {
	/* ===== INTRO =====
	- Bisa digunakan di semua sistem operasi
	*/
	args := os.Args
	for _, arg := range args {
		fmt.Println(arg)
	}

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println(err.Error())
	}
}
