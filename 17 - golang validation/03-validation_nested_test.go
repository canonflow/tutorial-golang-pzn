package main

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

type Address struct {
	City    string `validate:"required"`
	Country string `validate:"required"`
}

type User struct {
	Id      string  `validate:"required"`
	Name    string  `validate:"required"`
	Address Address `validate:"required`
}

func TestNestedStruct(t *testing.T) {
	validate := validator.New()

	request := User{
		Id:   "",
		Name: "",
		Address: Address{
			City:    "",
			Country: "",
		},
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
		/*
			Key: 'User.Id' Error:Field validation for 'Id' failed on the 'required' tag
			Key: 'User.Name' Error:Field validation for 'Name' failed on the 'required' tag
			Key: 'User.Address.City' Error:Field validation for 'City' failed on the 'required' tag
			Key: 'User.Address.Country' Error:Field validation for 'Country' failed on the 'required' tag
		*/
	}
}
