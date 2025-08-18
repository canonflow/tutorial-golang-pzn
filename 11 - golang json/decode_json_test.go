package _1___golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJson(t *testing.T) {
	jsongString := `{"FirstName":"Nathan","MiddleName":"Garzya","LastName":"Santoso","Age":22,"Married":false}`
	jsonBytes := []byte(jsongString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(*customer)
	/*
		=== RUN   TestDecodeJson
		{Nathan Garzya Santoso 22 false}
		--- PASS: TestDecodeJson (0.00s)
		PASS
	*/
}
