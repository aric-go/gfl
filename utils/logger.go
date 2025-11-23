package utils

import (
	"log"
	"os"
)

var (
	// gflLogger is the global logger instance for GFL CLI.
	// It's configured with a standardized prefix and standard log flags
	// to provide consistent logging throughout the application.
	gflLogger = log.New(os.Stdout, "[gfl] ", log.LstdFlags)
)

// Info logs informational messages to the console.
// This function is used for general information that helps users understand
// what the application is doing, such as configuration loading status,
// command execution progress, or general operational information.
//
// Parameters:
//   - msg: The message to log
//
// Example:
//   - Info("Configuration file loaded successfully")
//   - Info("Starting branch creation process")
func Info(msg string) {
	gflLogger.Println(msg)
}

// Infof logs formatted informational messages to the console.
// This function provides a convenient way to format log messages with
// variable arguments, similar to fmt.Printf.
//
// Parameters:
//   - format: The format string following printf conventions
//   - args: Arguments to be formatted into the message
//
// Example:
//   - Infof("Created branch: %s", branchName)
//   - Infof("Processing %d files...", count)
func Infof(format string, args ...interface{}) {
	gflLogger.Printf(format, args...)
}

// Error logs error messages to the console with an ERROR prefix.
// This function is used for non-fatal errors that occur during execution
// but don't necessarily require the application to exit immediately.
//
// Parameters:
//   - msg: The error message to log
//
// Example:
//   - Error("Failed to read configuration file")
//   - Error("Remote repository not found")
func Error(msg string) {
	gflLogger.Printf("ERROR: %s", msg)
}

// Errorf logs formatted error messages to the console with an ERROR prefix.
// This function provides formatted error logging with variable arguments.
// Use this for detailed error information that includes context or variables.
//
// Parameters:
//   - format: The format string following printf conventions
//   - args: Arguments to be formatted into the error message
//
// Example:
//   - Errorf("Failed to create branch: %v", err)
//   - Errorf("Configuration not found at: %s", configPath)
func Errorf(format string, args ...interface{}) {
	gflLogger.Printf("ERROR: "+format, args...)
}

// Warning logs warning messages to the console with a WARNING prefix.
// This function is used for situations that might require user attention
// but don't prevent the application from continuing to execute.
//
// Parameters:
//   - msg: The warning message to log
//
// Example:
//   - Warning("Working directory is not clean")
//   - Warning("Using default configuration values")
func Warning(msg string) {
	gflLogger.Printf("WARNING: %s", msg)
}

// Warningf logs formatted warning messages to the console with a WARNING prefix.
// This function provides formatted warning logging with variable arguments.
// Use this for detailed warning information that includes context or variables.
//
// Parameters:
//   - format: The format string following printf conventions
//   - args: Arguments to be formatted into the warning message
//
// Example:
//   - Warningf("Configuration file %s not found, using defaults", filename)
//   - Warningf("Branch %s may not be fully merged", branchName)
func Warningf(format string, args ...interface{}) {
	gflLogger.Printf("WARNING: "+format, args...)
}

// Success logs success messages to the console with a checkmark emoji prefix.
// This function is used for successful operations and provides positive feedback
// to users when commands complete successfully.
//
// Parameters:
//   - msg: The success message to log
//
// Example:
//   - Success("Configuration initialized successfully")
//   - Success("Branch created and checked out")
func Success(msg string) {
	gflLogger.Printf("✅ %s", msg)
}

// Successf logs formatted success messages to the console with a checkmark emoji prefix.
// This function provides formatted success logging with variable arguments.
// Use this for detailed success information that includes context or variables.
//
// Parameters:
//   - format: The format string following printf conventions
//   - args: Arguments to be formatted into the success message
//
// Example:
//   - Successf("Branch %s created successfully", branchName)
//   - Successf("Processed %d files successfully", count)
func Successf(format string, args ...interface{}) {
	gflLogger.Printf("✅ "+format, args...)
}