package test

import (
	"github.com/stretchr/testify/assert"
	"golang-dependency-injection/simple"
	"testing"
)

func TestSimpleServiceSuccess(t *testing.T) {
	simpleService, err := simple.InitializeService(false)
	assert.NotNil(t, simpleService)
	assert.Nil(t, err)
	/*
		=== RUN   TestSimpleServiceSuccess
		--- PASS: TestSimpleServiceSuccess (0.00s)
		PASS
	*/
}

func TestSimpleServiceError(t *testing.T) {
	simpleService, err := simple.InitializeService(true)
	assert.Nil(t, simpleService)
	assert.NotNil(t, err)
	/*
		=== RUN   TestSimpleServiceError
		--- PASS: TestSimpleServiceError (0.00s)
		PASS
	*/
}
