package utils

import (
	"fmt"
	"strings"
)

// BranchInfo contains information about the current branch
type BranchInfo struct {
	CurrentBranch    string
	TrackingBranch   string
	AheadCommits     int
	BehindCommits    int
	WorkingDirClean  bool
	RemoteURL        string
}

// GetTrackingBranch returns the remote tracking branch
func GetTrackingBranch() (string, error) {
	output, err := RunShell("git rev-parse --abbrev-ref --symbolic-full-name @{u}")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(output), nil
}

// GetAheadBehind returns the number of commits ahead and behind the remote
func GetAheadBehind() (ahead, behind int, err error) {
	output, err := RunShell("git rev-list --left-right --count @{u}...HEAD")
	if err != nil {
		return 0, 0, err
	}

	parts := strings.Split(strings.TrimSpace(output), "\t")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("unexpected rev-list output format")
	}

	fmt.Sscanf(parts[0], "%d", &behind)
	fmt.Sscanf(parts[1], "%d", &ahead)

	return ahead, behind, nil
}

// IsWorkingDirectoryClean checks if there are any uncommitted changes
func IsWorkingDirectoryClean() bool {
	output, err := RunShell("git status --porcelain")
	if err != nil {
		return false
	}
	return strings.TrimSpace(output) == ""
}

// GetRemoteURL returns the remote repository URL in a clean format
func GetRemoteURL() string {
	output, err := RunShell("git config --get remote.origin.url")
	if err != nil {
		return ""
	}
	url := strings.TrimSpace(output)

	// Parse URL to get a clean format (e.g., github.com/owner/repo)
	if strings.Contains(url, "github.com") {
		// SSH: git@github.com:owner/repo.git
		if strings.HasPrefix(url, "git@") {
			url = strings.TrimPrefix(url, "git@")
			url = strings.Replace(url, ":", "/", 1)
			url = strings.TrimSuffix(url, ".git")
		}
		// HTTPS: https://github.com/owner/repo.git
		if strings.HasPrefix(url, "https://") {
			url = strings.TrimPrefix(url, "https://")
			url = strings.TrimSuffix(url, ".git")
		}
	}

	return url
}

// GetBranchInfo collects all branch information
func GetBranchInfo() (*BranchInfo, error) {
	info := &BranchInfo{}

	// Get current branch (uses existing function from utils/git.go)
	branch, err := GetCurrentBranch()
	if err != nil {
		return nil, err
	}
	info.CurrentBranch = branch

	// Get tracking branch (may be empty if not tracking)
	tracking, _ := GetTrackingBranch()
	info.TrackingBranch = tracking

	// Get ahead/behind if tracking
	if tracking != "" {
		ahead, behind, err := GetAheadBehind()
		if err == nil {
			info.AheadCommits = ahead
			info.BehindCommits = behind
		}
	}

	// Check working directory
	info.WorkingDirClean = IsWorkingDirectoryClean()

	// Get remote URL
	info.RemoteURL = GetRemoteURL()

	return info, nil
}
