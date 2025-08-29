package logging

import (
	"github.com/sirupsen/logrus"
	"os"
)

func CustomLogger() *logrus.Logger {
	logger := logrus.New()
	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)

	return logger
}
