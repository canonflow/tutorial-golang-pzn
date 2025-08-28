package test

import (
	"github.com/stretchr/testify/assert"
	"golang-dependency-injection/simple"
	"testing"
)

func TestConnection(t *testing.T) {
	connection, cleanup := simple.InitializedConnection("Database")
	assert.NotNil(t, connection)

	cleanup()
	/*
		=== RUN   TestConnection
		Close connection Database
		Close file Database
		--- PASS: TestConnection (0.00s)
		PASS
	*/
}
