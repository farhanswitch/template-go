package utilities

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// LogLevel defines the level of logging
type LogLevel int

const (
	// DEBUG level for detailed debugging information
	DEBUG LogLevel = iota
	// INFO level for general information
	INFO
	// WARN level for warnings
	WARN
	// ERROR level for errors
	ERROR
)

// levelStrings maps LogLevel to string representation
var levelStrings = [...]string{
	"DEBUG",
	"INFO",
	"WARN",
	"ERROR",
}

// String returns the string representation of a LogLevel
func (l LogLevel) String() string {
	if l < DEBUG || l > ERROR {
		return "UNKNOWN"
	}
	return levelStrings[l]
}

// Log writes a log message with a specific level and context.
func Log(level LogLevel, path string, functionName string, payload interface{}, message string, data interface{}) {
	logLevelStr := os.Getenv("LOG_LEVEL")
	if logLevelStr == "" {
		logLevelStr = "INFO" // Default log level
	}

	currentLevel := getLogLevelFromString(logLevelStr)

	if level >= currentLevel {
		logData := map[string]interface{}{
			"level":    level.String(),
			"path":     path,
			"function": functionName,
			"payload":  payload,
			"message":  message,
			"data":     data,
			"time":     FormatTimeMillis(time.Now()),
		}

		logOutput, err := json.Marshal(logData)
		if err != nil {
			log.Printf("Error creating log message: %v", err)
			return
		}

		fmt.Println(string(logOutput))
	}
}

// getLogLevelFromString converts a string to a LogLevel
func getLogLevelFromString(levelStr string) LogLevel {
	switch strings.ToUpper(levelStr) {
	case "DEBUG":
		return DEBUG
	case "INFO":
		return INFO
	case "WARN":
		return WARN
	case "ERROR":
		return ERROR
	default:
		return INFO
	}
}
