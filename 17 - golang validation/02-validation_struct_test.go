package main

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

type LoginRequest struct {
	Username string `validate:"required,email"`
	Password string `validate:"required,min=5"`
}

func TestValidationStruct(t *testing.T) {
	validate := validator.New()

	loginRequest := LoginRequest{
		Username: "nathan@gmail.com",
		Password: "rahasia",
	}

	err := validate.Struct(loginRequest)
	if err != nil {
		fmt.Println(err.Error())
		// Key: 'LoginRequest.Username' Error:Field validation for 'Username' failed on the 'email' tag
	}

	/*
		=== RUN   TestValidationStruct
		--- PASS: TestValidationStruct (0.00s)
		PASS
	*/
}

func TestValidationError(t *testing.T) {
	validate := validator.New()

	loginRequest := LoginRequest{
		Username: "nathan",
		Password: "abc",
	}

	err := validate.Struct(loginRequest)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		for _, fieldError := range validationErrors {
			fmt.Println("Field:", fieldError.Field(), "Tag:", fieldError.Tag(), "Value:", fieldError.Value(), "Error:", fieldError.Error())

			/*
				Field: Username Tag: email Value: nathan Error: Key: 'LoginRequest.Username' Error:Field validation for 'Username' failed on the 'email' tag
				Field: Password Tag: min Value: abc Error: Key: 'LoginRequest.Password' Error:Field validation for 'Password' failed on the 'min' tag
			*/
		}
	}
}

type RegisterRequest struct {
	Username        string `validate:"required,email"`
	Password        string `validate:"required,min=5"`
	ConfirmPassword string `validate:"required,eqfield=Password"`
}

func TestValidationCrossField(t *testing.T) {
	validate := validator.New()

	registerRequest := RegisterRequest{
		Username:        "Nathan@gmail.com",
		Password:        "12345",
		ConfirmPassword: "12345",
	}

	err := validate.Struct(registerRequest)
	if err != nil {
		fmt.Println(err.Error())
		// Key: 'RegisterRequest.ConfirmPassword' Error:Field validation for 'ConfirmPassword' failed on the 'eqfield' tag
	}

	/*
		=== RUN   TestValidationCrossField
		--- PASS: TestValidationCrossField (0.00s)
		PASS
	*/
}
