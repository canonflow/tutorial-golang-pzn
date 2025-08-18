package _1___golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestEncode(t *testing.T) {
	logJson("Nathan")
	logJson(1)
	logJson(true)
	logJson([]string{"nathan", "garzya", "Santoso"})
	/*
		=== RUN   TestEncode
		"Nathan"
		1
		true
		["nathan","garzya","Santoso"]
		--- PASS: TestEncode (0.00s)
		PASS
	*/
}
