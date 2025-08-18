package _1___golang_json

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("customer-out.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Nathan",
		MiddleName: "Garzya",
		LastName:   "Santoso",
	}

	encoder.Encode(customer)
}
