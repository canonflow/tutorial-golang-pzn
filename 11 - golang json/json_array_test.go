package _1___golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Nathan",
		MiddleName: "Garzya",
		LastName:   "Santoso",
		Hobbies:    []string{"Reading", "Swimming", "Coding"},
	}

	bytes, _ := json.Marshal(customer)

	fmt.Println(string(bytes))
	/*
		PASS
		=== RUN   TestJsonArray
		{"FirstName":"Nathan","MiddleName":"Garzya","LastName":"Santoso","Age":0,"Married":false,"Hobbies":["Reading","Swimming","Coding"]}
		--- PASS: TestJsonArray (0.00s)
	*/
}

func TestJsonArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Nathan","MiddleName":"Garzya","LastName":"Santoso","Age":0,"Married":false,"Hobbies":["Reading","Swimming","Coding"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(*customer)
	fmt.Println(customer.Hobbies)
	/*
		=== RUN   TestJsonArrayDecode
		{Nathan Garzya Santoso 0 false [Reading Swimming Coding]}
		[Reading Swimming Coding]
		--- PASS: TestJsonArrayDecode (0.00s)
		PASS
	*/
}

func TestJsonArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName:  "Nathan",
		MiddleName: "Garzya",
		LastName:   "Santoso",
		Hobbies:    []string{"Reading", "Swimming", "Coding"},
		Addresses: []Address{
			{
				Street:     "Jalani saja",
				Country:    "Indonesia",
				PostalCode: "123",
			},
			{
				Street:     "Jalani saja 2",
				Country:    "Singapore",
				PostalCode: "321",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
	/*
		=== RUN   TestJsonArrayComplex
		{"FirstName":"Nathan","MiddleName":"Garzya","LastName":"Santoso","Age":0,"Married":false,"Hobbies":["Reading","Swimming","Coding"],"Addresses":[{"Street":"Jalani saja","Country":"Indonesia","PostalCode":"123"},{"Street":"Jalani saja 2","Country":"Singapore","PostalCode":"321"}]}
		--- PASS: TestJsonArrayComplex (0.00s)
		PASS
	*/
}

func TestJsonArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Nathan","MiddleName":"Garzya","LastName":"Santoso","Age":0,"Married":false,"Hobbies":["Reading","Swimming","Coding"],"Addresses":[{"Street":"Jalani saja","Country":"Indonesia","PostalCode":"123"},{"Street":"Jalani saja 2","Country":"Singapore","PostalCode":"321"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(*customer)
	fmt.Println(customer.Hobbies)
	fmt.Println(customer.Addresses)
	/*
		=== RUN   TestJsonArrayComplexDecode
		{Nathan Garzya Santoso 0 false [Reading Swimming Coding] [{Jalani saja Indonesia 123} {Jalani saja 2 Singapore 321}]}
		[Reading Swimming Coding]
		[{Jalani saja Indonesia 123} {Jalani saja 2 Singapore 321}]
		--- PASS: TestJsonArrayComplexDecode (0.00s)
		PASS

	*/
}

func TestOnlyJsonArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Jalani saja","Country":"Indonesia","PostalCode":"123"},{"Street":"Jalani saja 2","Country":"Singapore","PostalCode":"321"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}

	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(*addresses)
	/*
		=== RUN   TestOnlyJsonArrayComplexDecode
		[{Jalani saja Indonesia 123} {Jalani saja 2 Singapore 321}]
		--- PASS: TestOnlyJsonArrayComplexDecode (0.00s)
		PASS
	*/
}
