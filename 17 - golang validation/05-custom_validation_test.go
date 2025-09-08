package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
)

func MustValidUsername(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string) // .Field().String() juga bisa

	if ok {
		if value != strings.ToUpper(value) {
			return false
		}

		if len(value) < 5 {
			return false
		}
	}

	return true
}

func TestCustomValidation(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("username", MustValidUsername)

	type LoginRequest struct {
		Username string `validate:"required,username"`
		Password string `validate:"required"`
	}

	loginRequest := LoginRequest{
		Username: "canonflow",
		Password: "rahasia",
	}

	err := validate.Struct(loginRequest)
	if err != nil {
		fmt.Println(err.Error())
		// Key: 'LoginRequest.Username' Error:Field validation for 'Username' failed on the 'username' tag
	}
}
