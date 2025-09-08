package main

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestNestedStruct(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id      string  `validate:"required"`
		Name    string  `validate:"required"`
		Address Address `validate:"required"`
	}

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

func TestCollection(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id        string    `validate:"required"`
		Name      string    `validate:"required"`
		Addresses []Address `validate:"required,dive"`
	}

	validate := validator.New()

	request := User{
		Id:   "",
		Name: "",
		Addresses: []Address{
			{
				City:    "",
				Country: "",
			},
			{
				City:    "",
				Country: "",
			},
		},
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
		/*
			Key: 'User.Id' Error:Field validation for 'Id' failed on the 'required' tag
			Key: 'User.Name' Error:Field validation for 'Name' failed on the 'required' tag
			Key: 'User.Addresses[0].City' Error:Field validation for 'City' failed on the 'required' tag
			Key: 'User.Addresses[0].Country' Error:Field validation for 'Country' failed on the 'required' tag
			Key: 'User.Addresses[1].City' Error:Field validation for 'City' failed on the 'required' tag
			Key: 'User.Addresses[1].Country' Error:Field validation for 'Country' failed on the 'required' tag
		*/
	}
}

func TestBasicCollection(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id        string    `validate:"required"`
		Name      string    `validate:"required"`
		Addresses []Address `validate:"required,dive"`
		Hobbies   []string  `validate:"required,dive,required,min=3"`
	}

	validate := validator.New()

	request := User{
		Id:   "",
		Name: "",
		Addresses: []Address{
			{
				City:    "",
				Country: "",
			},
			{
				City:    "",
				Country: "",
			},
		},
		Hobbies: []string{"Swimming", "Eating", "X"},
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
		/*
			Key: 'User.Id' Error:Field validation for 'Id' failed on the 'required' tag
			Key: 'User.Name' Error:Field validation for 'Name' failed on the 'required' tag
			Key: 'User.Addresses[0].City' Error:Field validation for 'City' failed on the 'required' tag
			Key: 'User.Addresses[0].Country' Error:Field validation for 'Country' failed on the 'required' tag
			Key: 'User.Addresses[1].City' Error:Field validation for 'City' failed on the 'required' tag
			Key: 'User.Addresses[1].Country' Error:Field validation for 'Country' failed on the 'required' tag
			Key: 'User.Hobbies[2]' Error:Field validation for 'Hobbies[2]' failed on the 'min' tag
		*/
	}
}
