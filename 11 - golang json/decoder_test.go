package _1___golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestDecoder(t *testing.T) {
	reader, _ := os.Open("customer.json")
	decoder := json.NewDecoder(reader)

	customer := &Customer{}
	decoder.Decode(customer)
	fmt.Println(customer)
	/*
		=== RUN   TestDecoder
		&{Nathan Garzya Santoso 0 false [] []}
		--- PASS: TestDecoder (0.00s)
		PASS
	*/
}
