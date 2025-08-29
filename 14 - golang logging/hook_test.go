package _4___golang_logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"testing"
)

type SampleHook struct {
}

func (s SampleHook) Levels() []logrus.Level {
	// Kalau ada event log dengan level di bawah, maka hook-nya akan dieksekusi
	return []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel}
}

func (s SampleHook) Fire(entry *logrus.Entry) error {
	fmt.Println("Sample Hook Fire", entry.Level, entry.Message)
	return nil
}

func TestHook(t *testing.T) {
	logger := logrus.New()
	logger.AddHook(&SampleHook{})
	logger.SetLevel(logrus.DebugLevel)

	logger.Info("Hello Info")
	logger.Warn("Hello Warn")
	logger.Error("Hello Error")
	logger.Debug("Hello Debug")
	/*
		=== RUN   TestHook
		time="2025-08-29T13:54:12+07:00" level=info msg="Hello Info"
		Sample Hook Fire warning Hello Warn
		time="2025-08-29T13:54:12+07:00" level=warning msg="Hello Warn"
		Sample Hook Fire error Hello Error
		time="2025-08-29T13:54:12+07:00" level=error msg="Hello Error"
		time="2025-08-29T13:54:12+07:00" level=debug msg="Hello Debug"
		--- PASS: TestHook (0.00s)
		PASS
	*/
}
