package _1___golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func TestJsonTag(t *testing.T) {
	product := Product{
		Id:       "P00011",
		Name:     "Apple MacBook Pro",
		ImageUrl: "http://example.com/image.png",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
	/*
		=== RUN   TestJsonTag
		{"id":"P00011","name":"Apple MacBook Pro","image_url":"http://example.com/image.png"}
		--- PASS: TestJsonTag (0.00s)
		PASS
	*/
}

func TestJsonTagDecode(t *testing.T) {
	jsonString := `{"id":"P00011","name":"Apple MacBook Pro","image_url":"http://example.com/image.png"}`
	jsonBytes := []byte(jsonString)

	product := Product{}
	json.Unmarshal(jsonBytes, &product)
	fmt.Println(product)
	/*
		=== RUN   TestJsonTagDecode
		{P00011 Apple MacBook Pro http://example.com/image.png}
		--- PASS: TestJsonTagDecode (0.00s)
		PASS
	*/
}
