package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"time"
)

var Logger *logrus.Logger

// Function that initializes the logger
func init() {
	Logger = logrus.New()
	Logger.SetFormatter(&logrus.JSONFormatter{})

	logLevel, exists := viper.Get("LOG_LEVEL").(string)
	if !exists {
		logLevel = "info"
	}

	switch logLevel {
	case "debug":
		Logger.SetLevel(logrus.DebugLevel)
	case "info":
		Logger.SetLevel(logrus.InfoLevel)
	case "warn":
		Logger.SetLevel(logrus.WarnLevel)
	case "error":
		Logger.SetLevel(logrus.ErrorLevel)
	default:
		Logger.SetLevel(logrus.InfoLevel)
	}

	currentDate := time.Now().Format("2006-01-02")
	logFileName := "logs/" + currentDate + "_" + logLevel + ".log"

	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		err := os.Mkdir("logs", 0755)
		if err != nil {
			Logger.Fatalf("Error creating the logs directory: %v", err)
		}
	}

	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		Logger.Fatalf("Error opening the log file: %v", err)
	}

	Logger.SetOutput(logFile)
}

// Function to retrieve a logger
func GetLogger() *logrus.Logger {
	return Logger
}
