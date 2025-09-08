package main

import (
	"fmt"
	"regexp"
	"strconv"
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

var regexNumber = regexp.MustCompile("^[0-9]+$")

func MustValidPin(field validator.FieldLevel) bool {
	length, err := strconv.Atoi(field.Param())
	if err != nil {
		panic(err)
	}

	value := field.Field().String()
	// Memastikan hanya angka
	if !regexNumber.MatchString(value) {
		return false
	}

	// Memastikan panjang sesuai
	return len(value) == length
}

func TestCustomValidationWithParameter(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("pin", MustValidPin)

	type LoginRequest struct {
		Phone string `validate:"required,number"`
		Pin   string `validate:"required,pin=6"`
	}

	loginRequest := LoginRequest{
		Phone: "081234567890",
		Pin:   "12345",
	}

	err := validate.Struct(loginRequest)
	if err != nil {
		fmt.Println(err.Error())
		// Key: 'LoginRequest.Pin' Error:Field validation for 'Pin' failed on the 'pin' tag
	}
}

func TestOrRule(t *testing.T) {
	validate := validator.New()

	type Login struct {
		Username string `validate:"email|numeric"`
		Password string `validate:"required"`
	}

	login := Login{
		Username: "canonflow",
		Password: "rahasia",
	}

	err := validate.Struct(login)
	if err != nil {
		fmt.Println(err.Error())
		// Key: 'Login.Username' Error:Field validation for 'Username' failed on the 'email|numeric' tag
	}
}

func MustEqualIgnoreCase(field validator.FieldLevel) bool {
	value, _, _, ok := field.GetStructFieldOK2()

	if !ok {
		panic("field not found")
	}

	firstValue := strings.ToUpper(field.Field().String())
	secondValue := strings.ToUpper(value.String())

	return firstValue == secondValue
}

func TestCrossValidation(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("field_equals_ignore_case", MustEqualIgnoreCase)

	type User struct {
		Username string `validate:"required,field_equals_ignore_case=Email|field_equals_ignore_case=Phone"`
		Email    string `validate:"required,email"`
		Phone    string `validate:"required,numeric"`
		Name     string `validate:"required"`
	}

	user := User{
		Username: "canonflow1@gmail.com",
		Email:    "CanonFlow@gmail.com",
		Phone:    "081234567890",
		Name:     "Canonflow",
	}

	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err.Error())
		// Key: 'User.Username' Error:Field validation for 'Username' failed on the 'field_equals_ignore_case=Email|field_equals_ignore_case=Phone' tag
	}
}
