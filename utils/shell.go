package utils

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/afeiship/go-box"
	"github.com/briandowns/spinner"
)

// spin is a global spinner instance used for command execution visualization.
// It provides visual feedback during long-running operations.
//
// Configuration:
//   - Character set: spinner.CharSets[35] (dots animation)
//   - Interval: 200ms between animation frames
//   - Color: Green (set in RunCommandWithSpin)
var spin = spinner.New(spinner.CharSets[35], 200*time.Millisecond)

// debugBox prints debug information in a compact ASCII box.
// Uses go-box library for visual output with borders and colors.
func debugBox(title, command string) {
	lines := []string{
		"\x1b[33m⚙ DEBUG\x1b[0m " + "\x1b[36m" + title + "\x1b[0m",
		command,
	}
	box.PrintASCIIBox(lines)
	fmt.Println()
}

// RunShell executes a shell command and returns its output.
// This function provides a simple interface for executing shell commands
// with bash shell interpretation (for pipelines, redirects, etc.).
//
// Parameters:
//   - cmd: The shell command to execute
//
// Returns:
//   - string: The command's standard output
//   - error: Error if command execution fails
//
// Note:
//   - Uses bash for shell interpretation
//   - Returns raw output including newlines
//   - For user-facing operations, use RunCommandWithSpin instead
func RunShell(cmd string) (string, error) {
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return "", fmt.Errorf("shell command failed: %w", err)
	}
	return string(out), nil
}

// RunCommandWithSpin executes a command with a loading spinner for user feedback.
// This is the preferred method for user-facing operations that may take time.
// It provides visual feedback and debug output when configured.
//
// Parameters:
//   - command: The command to execute (can include arguments)
//   - message: The message to display with the spinner
//
// Returns:
//   - error: Error if command execution fails
//
// Features:
//   - Animated spinner with customizable message
//   - Debug mode support with colored output
//   - Automatic spinner cleanup on success/failure
//   - Smart command parsing (splits command and arguments)
//
// Example:
//   - RunCommandWithSpin("git fetch origin", "Fetching remote changes...")
//
// Debug output (when enabled):
//   - Shows the actual command being executed
//   - Provides rainbow emoji indicator for debug visibility
func RunCommandWithSpin(command string, message string) error {
	// Load configuration for debug mode
	config := ReadConfig()
	if config == nil {
		log.Fatal("Failed to read configuration file")
	}

	// Show debug information if enabled (before spinner starts)
	if config.Debug {
		debugBox("Executing command", command)
	}

	// Configure spinner appearance
	_ = spin.Color("green")
	spin.Start()
	spin.Suffix = message

	// Parse command into executable and arguments
	// This handles commands with spaces and quotes properly
	cmdArgs := strings.Fields(command)
	if len(cmdArgs) == 0 {
		spin.Stop()
		return fmt.Errorf("empty command provided")
	}

	// Execute the command with parsed arguments
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)

	// Run the command and handle errors
	if err := cmd.Run(); err != nil {
		spin.Stop()
		return fmt.Errorf("command execution failed: %w", err)
	}

	spin.Stop()
	return nil
}

// RunCommandWithArgs executes a command with explicit arguments and a spinner.
// This is an alternative to RunCommandWithSpin when you have pre-split arguments.
//
// Parameters:
//   - executable: The command to execute (e.g., "git", "gh", "npm")
//   - args: Slice of command arguments
//   - message: The message to display with the spinner
//
// Returns:
//   - error: Error if command execution fails
//
// Features:
//   - Same spinner behavior as RunCommandWithSpin
//   - No shell interpretation, args are passed directly
//   - Safer for commands with user-provided arguments
//   - Debug mode support
//
// Example:
//   - RunCommandWithArgs("git", []string{"fetch", "origin"}, "Fetching remote changes...")
//   - RunCommandWithArgs("gh", []string{"pr", "create", "--title", "My PR"}, "Creating PR...")
func RunCommandWithArgs(executable string, args []string, message string) error {
	config := ReadConfig()
	if config == nil {
		log.Fatal("Failed to read configuration file")
	}

	if config.Debug {
		// Quote arguments that contain spaces to make debug output clear
		var quotedArgs []string
		for _, arg := range args {
			if strings.ContainsAny(arg, " \t\n") {
				quotedArgs = append(quotedArgs, `"`+arg+`"`)
			} else {
				quotedArgs = append(quotedArgs, arg)
			}
		}
		fullCmd := executable + " " + strings.Join(quotedArgs, " ")
		debugBox("Executing command with args", fullCmd)
	}

	_ = spin.Color("green")
	spin.Start()
	spin.Suffix = message

	cmd := exec.Command(executable, args...)
	if err := cmd.Run(); err != nil {
		spin.Stop()
		return fmt.Errorf("command execution failed: %w", err)
	}

	spin.Stop()
	return nil
}

// GetLocalBranches retrieves a list of all local Git branches.
// It executes 'git branch' and parses the output into a slice of branch names.
//
// Returns:
//   - []string: List of local branch names
//   - nil: If command execution fails (error is logged)
//
// Output format:
//   - Each string represents a line from 'git branch' output
//   - Includes the asterisk (*) for the current branch
//   - May contain whitespace formatting
//
// Example output:
//   - ["* main", "  develop", "  feature/user-auth", "  hotfix/security-fix"]
func GetLocalBranches() []string {
	output, err := RunShell("git branch")
	if err != nil {
		Errorf("Failed to get local branches: %v", err)
		return nil
	}

	// Convert output to string and split by lines
	outputStr := strings.TrimSpace(string(output))
	if outputStr == "" {
		return []string{}
	}

	branches := strings.Split(outputStr, "\n")
	return branches
}

// IsCommandAvailable checks if a command-line tool is available in the system PATH.
// This is useful for verifying dependencies before attempting to use them.
//
// Parameters:
//   - name: The name of the command to check (e.g., "gh", "git", "node")
//
// Returns:
//   - bool: true if the command is available, false otherwise
//
// Use cases:
//   - Checking for GitHub CLI availability before creating PRs
//   - Verifying Git installation
//   - Dependency validation for optional features
//
// Examples:
//   - IsCommandAvailable("gh") -> true if GitHub CLI is installed
//   - IsCommandAvailable("nonexistent") -> false
func IsCommandAvailable(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}
