package _1_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func DisplayNumber(number int) {
	fmt.Println("Display Number:", number)
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Ups ...")

	// Just in case biar goroutine-nya sudah selesai sebelum aplikasi-nya selesai
	time.Sleep(1 * time.Second)

	/*
		=== RUN   TestCreateGoroutine
		Ups ...
		Hello World
		--- PASS: TestCreateGoroutine (1.00s)
		PASS
		ok      golang-goroutines       1.136s
	*/
}

func TestManyGoroutine(t *testing.T) {
	for i := range 100000 {
		go DisplayNumber(i)
	}

	time.Sleep(7 * time.Second)
}
