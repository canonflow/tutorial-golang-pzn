package _4___golang_logging

import (
	"github.com/sirupsen/logrus"
	"golang-logging/logging"
	"testing"
)

func TestLogger(t *testing.T) {
	var logger *logrus.Logger = logrus.New()

	logger.Println("Hello Logger!")
}

func TestLevel(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel) // Biar log trace dan debuh bisa muncul

	logger.Trace("This is a Trace")
	logger.Debug("This is a Debug")
	logger.Info("This is a Info")
	logger.Warn("This is a Warn")
	logger.Error("This is a Error")
	/*
		=== RUN   TestLevel
		time="2025-08-29T13:08:42+07:00" level=trace msg="This is a Trace"
		time="2025-08-29T13:08:42+07:00" level=debug msg="This is a Debug"
		time="2025-08-29T13:08:42+07:00" level=info msg="This is a Info"
		time="2025-08-29T13:08:42+07:00" level=warning msg="This is a Warn"
		time="2025-08-29T13:08:42+07:00" level=error msg="This is a Error"
		--- PASS: TestLevel (0.00s)
		PASS
	*/
}

func TestOutput(t *testing.T) {
	/*
		logger := logrus.New()
		file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		logger.SetOutput(file)
	*/
	logger := logging.CustomLogger()
	logger.Info("This is a Info")
	logger.Warn("This is a Warn")
	logger.Error("This is a Error")
}

func TestFormatter(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(new(logrus.JSONFormatter)) // atau pakai &logrus.JSONFormatter

	logger.Info("Hello Formatter!")
	logger.Warn("Hello Formatter!")
	logger.Error("Hello Formatter!")
	/*
		=== RUN   TestFormatter
		{"level":"info","msg":"Hello Formatter!","time":"2025-08-29T13:22:35+07:00"}
		{"level":"warning","msg":"Hello Formatter!","time":"2025-08-29T13:22:35+07:00"}
		{"level":"error","msg":"Hello Formatter!","time":"2025-08-29T13:22:35+07:00"}
		--- PASS: TestFormatter (0.00s)
		PASS
	*/
}
