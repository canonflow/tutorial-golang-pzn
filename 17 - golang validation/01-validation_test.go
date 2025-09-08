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

func TestValidationTwoVariables(t *testing.T) {
	validate := validator.New()
	password := "rahasia"
	confirmPassword := "rahasia"

	err := validate.VarWithValue(password, confirmPassword, "eqfield")
	if err != nil {
		fmt.Println(err.Error())
	}

	/*
		=== RUN   TestValidationTwoVariables
		--- PASS: TestValidationTwoVariables (0.00s)
		PASS
	*/
}

func TestMultipleTag(t *testing.T) {
	validate := validator.New()

	user := "1234"

	err := validate.Var(user, "required,numeric")
	if err != nil {
		fmt.Println(err.Error())
	}

	/*
		=== RUN   TestMultipleTag
		--- PASS: TestMultipleTag (0.00s)
	*/
}

func TestTagParameter(t *testing.T) {
	validate := validator.New()

	user := "99999"

	err := validate.Var(user, "numeric,min=5,max=10")
	if err != nil {
		fmt.Println(err.Error())
	}

	/*
		=== RUN   TestTagParameter
		--- PASS: TestTagParameter (0.00s)
		PASS
	*/
}
