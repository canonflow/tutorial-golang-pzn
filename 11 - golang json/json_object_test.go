package _1___golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
}

func TestJsonObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Nathan",
		MiddleName: "Garzya",
		LastName:   "Santoso",
		Age:        22,
		Married:    false,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
	/*
		=== RUN   TestJsonObject
		{"FirstName":"Nathan","MiddleName":"Garzya","LastName":"Santoso","Age":22,"Married":false}
		--- PASS: TestJsonObject (0.00s)
		PASS
	*/
}
