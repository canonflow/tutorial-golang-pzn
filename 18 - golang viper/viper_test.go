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

func TestJson(t *testing.T) {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath(".")

	// Membaca
	err := config.ReadInConfig()
	assert.Nil(t, err)
	assert.Equal(t, "golang-viper", config.GetString("app.name"))
	assert.Equal(t, "Nathan Garzya", config.GetString("app.author"))
	assert.Equal(t, 3306, config.GetInt("database.port"))
	assert.True(t, config.GetBool("database.show_sql"))

	/*
		=== RUN   TestJson
		--- PASS: TestJson (0.00s)
		PASS
	*/
}

func TestYaml(t *testing.T) {
	config := viper.New()
	// config.SetConfigName("config")
	// config.SetConfigType("yaml")
	config.SetConfigFile("config.yaml")
	config.AddConfigPath(".")

	// Membaca
	err := config.ReadInConfig()
	assert.Nil(t, err)
	assert.Equal(t, "golang-viper", config.GetString("app.name"))
	assert.Equal(t, "Nathan Garzya", config.GetString("app.author"))
	assert.Equal(t, 3306, config.GetInt("database.port"))
	assert.True(t, config.GetBool("database.show_sql"))

	/*
		=== RUN   TestYaml
		--- PASS: TestYaml (0.00s)
		PASS
	*/
}

func TestEnv(t *testing.T) {
	config := viper.New()
	// config.SetConfigName("config")
	// config.SetConfigType("yaml")
	config.SetConfigFile("config.env")
	config.AddConfigPath(".")

	// Membaca
	err := config.ReadInConfig()
	assert.Nil(t, err)
	assert.Equal(t, "golang-viper", config.GetString("APP_NAME"))
	assert.Equal(t, "Nathan Garzya", config.GetString("APP_AUTHOR"))
	assert.Equal(t, 3306, config.GetInt("DATABASE_PORT"))
	assert.True(t, config.GetBool("DATABASE_SHOW_SQL"))

	/*
		=== RUN   TestEnv
		--- PASS: TestEnv (0.00s)
		PASS
	*/
}
