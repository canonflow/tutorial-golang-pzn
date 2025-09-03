package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Data[T any] struct {
	First  T
	Second T
}

func (d *Data[_]) SayHello(name string) string {
	// Gak butuh generic type, pakai _ (underscore)
	return "Hello " + name
}

func (d *Data[T]) ChangeFirst(first T) T {
	d.First = first
	return d.First
}

func TestData(t *testing.T) {
	data := Data[string]{
		First:  "Nathan",
		Second: "Garzya",
	}
	fmt.Println(data)

	/*
		=== RUN   TestData
		{Nathan Garzya}
		--- PASS: TestData (0.00s)
		PASS
	*/
}

func TestGenericMethod(t *testing.T) {
	data := Data[string]{
		First:  "Nathan",
		Second: "Garzya",
	}

	assert.Equal(t, "Canonflow", data.ChangeFirst("Canonflow"))
	assert.Equal(t, "Hello Nathan", data.SayHello("Nathan"))
	fmt.Println(data)

	/*
		=== RUN   TestGenericMethod
		{Canonflow Garzya}
		--- PASS: TestGenericMethod (0.00s)
		PASS
	*/
}
