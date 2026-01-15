package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

// extractOwnerAndRepo extracts the owner and repository name from a Git URL.
// It supports both SSH and HTTPS URL formats commonly used with Git repositories.
//
// Supported formats:
//   - SSH: git@github.com:owner/repo.git
//   - HTTPS: https://github.com/owner/repo.git
//
// Parameters:
//   - url: The Git repository URL
//
// Returns:
//   - string: The "owner/repo" format
//   - error: Error if URL format is unsupported or invalid
//
// Examples:
//   - "git@github.com:user/project.git" -> "user/project"
//   - "https://github.com/user/project.git" -> "user/project"
func extractOwnerAndRepo(url string) (string, error) {
	if strings.HasPrefix(url, "git@") {
		// Handle SSH format: git@github.com:owner/repo.git
		parts := strings.Split(url, ":")
		if len(parts) != 2 {
			return "", fmt.Errorf("invalid SSH git URL format: %s", url)
		}
		// Remove .git suffix and return the path part
		path := strings.TrimSuffix(parts[1], ".git")
		return path, nil
	} else if strings.HasPrefix(url, "https://") {
		// Handle HTTPS format: https://github.com/owner/repo.git
		parts := strings.Split(url, "/")
		if len(parts) < 2 {
			return "", fmt.Errorf("invalid HTTPS git URL format: %s", url)
		}
		// Join the last two parts (owner and repo) and remove .git suffix
		orgRepo := strings.Join(parts[len(parts)-2:], "/")
		orgRepo = strings.TrimSuffix(orgRepo, ".git")
		return orgRepo, nil
	}

	return "", fmt.Errorf("unsupported git URL format: %s", url)
}

// GetRepository retrieves the current Git repository's owner and name.
// It queries Git configuration to get the remote origin URL and extracts
// the repository identifier in "owner/repo" format.
//
// This function is typically used for:
//   - Creating GitHub Pull Request URLs
//   - Generating repository-specific operations
//   - Repository identification in logging
//
// Returns:
//   - string: The repository identifier in "owner/repo" format
//   - error: Error if repository URL cannot be retrieved or parsed
//
// Example:
//   - Returns "myorg/myproject" for a repository at git@github.com:myorg/myproject.git
func GetRepository() (string, error) {
	// Execute git command to get the remote origin URL
	cmd := exec.Command("git", "config", "--get", "remote.origin.url")
	url, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get repository URL: %w", err)
	}

	// Clean up the URL and extract owner/repo
	cleanURL := strings.TrimSpace(string(url))
	return extractOwnerAndRepo(cleanURL)
}

// GetCurrentBranch retrieves the name of the current Git branch.
// It uses 'git rev-parse --abbrev-ref HEAD' to get the branch name,
// which is a reliable method that works in all Git scenarios.
//
// This function is commonly used for:
//   - Determining the source branch for Pull Requests
//   - Validating branch operations
//   - Branch-specific workflow logic
//
// Returns:
//   - string: The current branch name (e.g., "main", "develop", "feature/user-auth")
//   - error: Error if not in a Git repository or command execution fails
//
// Special cases:
//   - Returns "HEAD" when in detached HEAD state
//   - Returns empty string if not a Git repository
func GetCurrentBranch() (string, error) {
	// Execute git command to get the current branch name
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get current branch: %w", err)
	}

	// Remove trailing newline and return the branch name
	branchName := strings.TrimSpace(string(output))
	return branchName, nil
}

// GetRemoteBranches retrieves a list of all remote branches from the origin.
// It executes 'git branch -r' and parses the output into a clean slice of branch names.
//
// Returns:
//   - []string: List of remote branch names (e.g., ["origin/main", "origin/develop"])
//   - error: Error if the command execution fails
//
// Note:
//   - The branch names include the "origin/" prefix
//   - Empty lines and whitespace are trimmed
func GetRemoteBranches() ([]string, error) {
	cmd := exec.Command("git", "branch", "-r")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get remote branches: %w", err)
	}

	// Parse the output into a slice of branch names
	var branches []string
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		branch := strings.TrimSpace(line)
		// Skip empty lines and HEAD pointer
		if branch == "" || strings.Contains(branch, "->") {
			continue
		}
		branches = append(branches, branch)
	}

	return branches, nil
}
