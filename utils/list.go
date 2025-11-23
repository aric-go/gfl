package utils

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"strings"
)

// BuildCommandList creates an interactive branch selection interface for Git checkout.
// This function presents the user with a selection list of available branches and
// automatically executes the git checkout command for the selected branch.
//
// The function uses the survey library to create a terminal-based interactive prompt
// that provides a user-friendly way to switch between branches without needing
// to remember or type the full branch names.
//
// Parameters:
//   - branches: A slice of branch names available for selection
//
// Side Effects:
//   - Displays an interactive selection prompt in the terminal
//   - Executes 'git checkout' command for the selected branch
//   - Logs success or error messages
//
// User Experience:
//   - Shows a numbered or arrow-key navigable list of branches
//   - Supports keyboard navigation and selection
//   - Provides immediate feedback on selection
//   - Automatically switches to the selected branch
//
// Error Handling:
//   - Logs survey interaction errors
//   - Logs git checkout command execution errors
//   - Continues execution even if the operation fails
//
// Example Output:
//   ? Choose a branch:
//     ▸ main
//       develop
//       feature/aric/user-auth
//       hotfix/security-fix
func BuildCommandList(branches []string) {
	// Define the structure to hold the survey answer
	answers := struct {
		Module string `survey:"branch"` // The survey field name and corresponding struct field
	}{}

	// Create the survey question with a select prompt
	var qs = []*survey.Question{
		{
			Name: "module",
			Prompt: &survey.Select{
				Message: "Choose a branch:",
				Options: branches,
				// Additional options could include:
				// - Help text for the selection
				// - Default selection
				// - Filtering options
			},
		},
	}

	// Execute the interactive survey
	err := survey.Ask(qs, &answers)
	if err != nil {
		Error(fmt.Sprintf("Survey interaction failed: %v", err))
		return
	}

	// Execute git checkout command for the selected branch
	selectedBranch := answers.Module
	command := fmt.Sprintf("git checkout %s", selectedBranch)

	// Run the checkout command and handle any errors
	output, err := RunShell(command)
	if err != nil {
		Error(fmt.Sprintf("Failed to checkout branch '%s': %v", selectedBranch, err))
	} else {
		// Provide feedback on successful checkout
		Successf("Successfully checked out branch: %s", selectedBranch)

		// Show any output from the git command if it exists
		if len(strings.TrimSpace(output)) > 0 {
			Infof("Git output: %s", strings.TrimSpace(output))
		}
	}
}
