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
