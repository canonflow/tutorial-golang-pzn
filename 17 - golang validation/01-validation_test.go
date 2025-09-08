package main

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidation(t *testing.T) {
	var validate *validator.Validate = validator.New()

	if validate == nil {
		t.Error("Expected validator to be initialized")
	}
	/*
		=== RUN   TestValidation
		--- PASS: TestValidation (0.00s)
		PASS
	*/
}
