package _4___golang_logging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestEntry(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	entry := logrus.NewEntry(logger)

	entry.WithField("username", "canonflow")
	entry.Info("Hello Entry")

	/*
		=== RUN   TestEntry
		{"level":"info","msg":"Hello Entry","time":"2025-08-29T13:45:44+07:00"}
		--- PASS: TestEntry (0.00s)
		PASS

	*/
}
