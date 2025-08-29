package _4___golang_logging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "Nathan").Info("Hello World")

	logger.WithField("username", "garzya").
		WithField("name", "Garzya").
		Info("Hello Garzya")

	/*
		{"level":"info","msg":"Hello World","time":"2025-08-29T13:27:42+07:00","username":"Nathan"}
		{"level":"info","msg":"Hello Garzya","name":"Garzya","time":"2025-08-29T13:27:42+07:00","username":"garzya"}
	*/
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "nathan",
		"name":     "Nathan Garzya",
	}).Info("Hello Garzya")

	/*
		=== RUN   TestFields
		{"level":"info","msg":"Hello Garzya","name":"Nathan Garzya","time":"2025-08-29T13:31:33+07:00","username":"nathan"}
		--- PASS: TestFields (0.00s)
		PASS
	*/
}
