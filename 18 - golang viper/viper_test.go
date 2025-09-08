package main

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestViper(t *testing.T) {
	var viper *viper.Viper = viper.New()

	assert.NotNil(t, viper)

	/*
		=== RUN   TestViper
		--- PASS: TestViper (0.00s)
		PASS
	*/
}
