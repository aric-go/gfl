package utils

import (
	"log"
	"os"
)

var (
	// GFL logger with prefix
	gflLogger = log.New(os.Stdout, "[gfl] ", log.LstdFlags)
)

// Info logs informational messages
func Info(msg string) {
	gflLogger.Println(msg)
}

// Infof logs formatted informational messages
func Infof(format string, args ...interface{}) {
	gflLogger.Printf(format, args...)
}

// Error logs error messages
func Error(msg string) {
	gflLogger.Printf("ERROR: %s", msg)
}

// Errorf logs formatted error messages
func Errorf(format string, args ...interface{}) {
	gflLogger.Printf("ERROR: "+format, args...)
}

// Warning logs warning messages
func Warning(msg string) {
	gflLogger.Printf("WARNING: %s", msg)
}

// Warningf logs formatted warning messages
func Warningf(format string, args ...interface{}) {
	gflLogger.Printf("WARNING: "+format, args...)
}

// Success logs success messages
func Success(msg string) {
	gflLogger.Printf("✅ %s", msg)
}

// Successf logs formatted success messages
func Successf(format string, args ...interface{}) {
	gflLogger.Printf("✅ "+format, args...)
}