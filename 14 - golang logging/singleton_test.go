package _4___golang_logging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestSingleton(t *testing.T) {
	logrus.Info("Hello Info")
	logrus.Error("Hello Error")

	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Info("Hello Info")
	logrus.Error("Hello Error")

	/*
		=== RUN   TestSingleton
		time="2025-08-29T13:59:14+07:00" level=info msg="Hello Info"
		time="2025-08-29T13:59:14+07:00" level=error msg="Hello Error"
		{"level":"info","msg":"Hello Info","time":"2025-08-29T13:59:14+07:00"}
		{"level":"error","msg":"Hello Error","time":"2025-08-29T13:59:14+07:00"}
		--- PASS: TestSingleton (0.00s)
		PASS
	*/
}
