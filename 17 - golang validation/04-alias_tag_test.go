package main

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestAliasTag(t *testing.T) {
	validate := validator.New()

	validate.RegisterAlias("varchar", "required,max=255")

	type Seller struct {
		Id     string `validate:"varchar"`
		Name   string `validate:"varchar"`
		Owner  string `validate:"varchar"`
		Slogan string `validate:"varchar"`
	}

	seller := Seller{
		Id:     "1",
		Name:   "Gopher",
		Owner:  "Canonflow",
		Slogan: "",
	}

	err := validate.Struct(seller)
	if err != nil {
		fmt.Println(err.Error())
		// Key: 'Seller.Slogan' Error:Field validation for 'Slogan' failed on the 'varchar' tag
	}
}
