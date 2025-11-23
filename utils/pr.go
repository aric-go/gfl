package utils

import (
	"fmt"
	"os/exec"
	"strings"
	"github.com/pkg/browser"
)

// CreatePr creates a GitHub Pull Request by opening the PR creation page in the default browser.
// This function constructs the GitHub comparison URL and opens it for manual completion.
//
// The function follows the GitHub URL pattern:
//   https://github.com/{owner}/{repo}/compare/{base}...{head}?expand=1
//
// Parameters:
//   - base: The target branch to merge into (e.g., "main", "develop")
//   - head: The source branch containing changes (e.g., "feature/aric/user-auth")
//
// URL Query Parameters:
//   - expand=1: Automatically expands the comparison view
//   - The URL format handles cross-branch comparisons
//
// Side Effects:
//   - Opens the default browser with the PR creation URL
//   - Logs success or error messages
//   - Does not create the PR automatically (user must complete manually)
//
// Example URL:
//   https://github.com/myorg/myproject/compare/develop...feature/aric/user-auth?expand=1
func CreatePr(base string, head string) {
	repo, err := GetRepository()
	if err != nil {
		Errorf("Failed to get repository information: %v", err)
		return
	}

	// Generate GitHub PR URL for branch comparison
	// Example: https://github.com/owner/repo/compare/base...head?expand=1
	url := fmt.Sprintf("https://github.com/%s/compare/%s...%s?expand=1", repo, base, head)

	// Open the URL in the default browser
	err = browser.OpenURL(url)
	if err != nil {
		Errorf("Failed to open browser: %v", err)
		Infof("Please manually open this URL: %s", url)
	} else {
		Infof("Opened PR creation page: %s", url)
	}
}

// SyncProductionToDev synchronizes the production branch with the development branch.
// This function performs a complete sync operation to ensure the development branch
// contains all changes from the production branch, typically done before starting
// new development cycles.
//
// Process Overview:
//   1. Verify working directory is clean
//   2. Save current branch for later restoration
//   3. Fetch latest changes from remote
//   4. Switch to development branch
//   5. Update development branch with remote changes
//   6. Merge production branch into development branch
//   7. Push synchronized development branch to remote
//   8. Restore original branch
//
// Parameters:
//   - productionBranch: The production branch name (e.g., "main", "master")
//   - devBranch: The development branch name (e.g., "develop", "dev")
//
// Returns:
//   - bool: true if synchronization succeeded, false otherwise
//
// Safety Features:
//   - Checks for clean working directory before starting
//   - Automatically rolls back to original branch on failure
//   - Provides detailed error reporting
//   - Only shows important command output to reduce noise
func SyncProductionToDev(productionBranch, devBranch string) bool {
	Infof("Syncing production branch '%s' to development branch '%s'...", productionBranch, devBranch)

	// Step 1: Ensure working directory is clean before starting sync
	if !isWorkingDirectoryClean() {
		Errorf("Working directory is not clean. Please commit or stash your changes first.")
		return false
	}

	// Step 2: Save current branch for restoration after sync
	currentBranch, err := GetCurrentBranch()
	if err != nil {
		Errorf("Failed to get current branch: %v", err)
		return false
	}

	// Step 3: Define the sequence of Git commands for synchronization
	commands := [][]string{
		{"git", "fetch", "origin"},                                 // Fetch latest remote changes
		{"git", "checkout", devBranch},                             // Switch to development branch
		{"git", "pull", "origin", devBranch},                       // Update development branch
		{"git", "merge", "origin/" + productionBranch},             // Merge production into development
		{"git", "push", "origin", devBranch},                       // Push synchronized development branch
	}

	// Step 4: Execute each command with error handling and rollback capability
	for i, cmd := range commands {
		Infof("Executing: %s", strings.Join(cmd, " "))
		output, err := exec.Command(cmd[0], cmd[1:]...).CombinedOutput()

		if err != nil {
			Errorf("Command execution failed: %v", err)
			Errorf("Error output: %s", string(output))

			// Attempt to rollback to original branch on failure
			Warningf("Attempting to rollback to original branch '%s'...", currentBranch)
			if i > 0 { // Only rollback if we've already switched branches
				rollbackCmd := exec.Command("git", "checkout", currentBranch)
				rollbackCmd.CombinedOutput() // Ignore rollback errors to avoid masking original error
			}
			return false
		}

		// Display important output only (filter out noise)
		if len(output) > 0 {
			outputStr := string(output)
			if containsImportantOutput(outputStr) {
				Infof("Output: %s", outputStr)
			}
		}
	}

	Successf("Successfully synchronized %s to %s", productionBranch, devBranch)
	return true
}

// isWorkingDirectoryClean checks if the Git working directory has no uncommitted changes.
// This function uses 'git status --porcelain' to determine if there are any pending changes.
//
// Returns:
//   - bool: true if working directory is clean, false otherwise
//
// Method:
//   - Uses --porcelain format for machine-readable output
//   - Empty output indicates no pending changes
//   - Any output (staged or unstaged) indicates a dirty working directory
//
// This check ensures safe branch switching and synchronization operations.
func isWorkingDirectoryClean() bool {
	output, err := exec.Command("git", "status", "--porcelain").CombinedOutput()
	if err != nil {
		// If git command fails, assume working directory is not clean
		return false
	}

	// Trim whitespace and check if output is empty
	return len(strings.TrimSpace(string(output))) == 0
}

// containsImportantOutput filters Git command output to show only relevant information.
// This function reduces noise by filtering out verbose Git messages that don't
// require user attention.
//
// Parameters:
//   - output: The Git command output string
//
// Returns:
//   - bool: true if output contains important information worth displaying
//
// Important output includes:
//   - Merge messages
//   - File change notifications
//   - Fast-forward updates
//   - "Already up to date" messages
func containsImportantOutput(output string) bool {
	return strings.Contains(output, "Already up to date") ||
		   strings.Contains(output, "Fast-forward") ||
		   strings.Contains(output, "Merge") ||
		   strings.Contains(output, "file changed") ||
		   strings.Contains(output, "insertion") ||
		   strings.Contains(output, "deletion")
}
