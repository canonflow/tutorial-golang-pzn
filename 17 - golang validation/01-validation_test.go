package main

import (
	"fmt"
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

func TestValidationVariable(t *testing.T) {
	validate := validator.New()

	user := "Nathan"

	err := validate.Var(user, "required")
	if err != nil {
		fmt.Println(err.Error())
		/*
			=== RUN   TestValidationVariable
			--- PASS: TestValidationVariable (0.00s)
		*/
	}
}
