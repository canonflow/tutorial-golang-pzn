package _1___golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T) {
	jsonString := `{"id": "P0002", "name": "Apple Mac Book Pro", "price": 20000000}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}

	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
	/*
		=== RUN   TestMapDecoded
		P0002
		Apple Mac Book Pro
		2e+07
		--- PASS: TestMapDecode (0.00s)
		PASS
	*/
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":    "P0002",
		"name":  "Apple Mac Book Pro",
		"price": 2000000000,
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
	/*
		=== RUN   TestMapEncode
		{"id":"P0002","name":"Apple Mac Book Pro","price":2000000000}
		--- PASS: TestMapEncode (0.00s)
		PASS
	*/
}
