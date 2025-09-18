package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var db = OpenConnection()

func TestConnection(t *testing.T) {
	assert.NotNil(t, db)

	/*
		=== RUN   TestConnection
		--- PASS: TestConnection (0.00s)
		PASS
	*/
}
