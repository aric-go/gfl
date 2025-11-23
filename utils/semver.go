package utils

import (
	"bytes"
	"fmt"
	"golang.org/x/mod/semver"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

// IncrementVersion increments a semantic version based on the specified version type.
// This function follows Semantic Versioning 2.0.0 (SemVer) conventions where
// MAJOR version changes indicate incompatible API changes,
// MINOR version changes add functionality in a backward compatible manner,
// and PATCH version changes indicate backward compatible bug fixes.
//
// Version increment rules:
//   - MAJOR: Resets MINOR and PATCH to 0, increments MAJOR (e.g., v1.2.3 → v2.0.0)
//   - MINOR: Resets PATCH to 0, increments MINOR (e.g., v1.2.3 → v1.3.0)
//   - PATCH: Increments PATCH only (e.g., v1.2.3 → v1.2.4)
//
// Parameters:
//   - currentVersion: Current version string in format "vX.Y.Z"
//   - versionType: Type of increment ("major", "minor", "patch" - case insensitive)
//
// Returns:
//   - string: New version string in format "vX.Y.Z"
//   - error: Error if version format is invalid or versionType is unsupported
//
// Examples:
//   - IncrementVersion("v1.2.3", "patch") → "v1.2.4"
//   - IncrementVersion("v1.2.3", "minor") → "v1.3.0"
//   - IncrementVersion("v1.2.3", "major") → "v2.0.0"
func IncrementVersion(currentVersion string, versionType string) (string, error) {
	// Validate input version format
	if !semver.IsValid(currentVersion) {
		return "", fmt.Errorf("invalid semantic version: %s", currentVersion)
	}

	// Remove 'v' prefix for parsing
	if len(currentVersion) < 2 || currentVersion[0] != 'v' {
		return "", fmt.Errorf("version must start with 'v': %s", currentVersion)
	}

	version := currentVersion[1:]
	parts := strings.Split(version, ".")
	if len(parts) != 3 {
		return "", fmt.Errorf("invalid version format, expected X.Y.Z: %s", currentVersion)
	}

	// Parse MAJOR, MINOR, PATCH as integers
	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return "", fmt.Errorf("failed to parse MAJOR version: %v", err)
	}
	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", fmt.Errorf("failed to parse MINOR version: %v", err)
	}
	patch, err := strconv.Atoi(parts[2])
	if err != nil {
		return "", fmt.Errorf("failed to parse PATCH version: %v", err)
	}

	// Increment the appropriate version component
	switch strings.ToUpper(versionType) {
	case "MAJOR":
		major++
		minor = 0  // Reset MINOR for MAOR version change
		patch = 0  // Reset PATCH for MAJOR version change
	case "MINOR":
		minor++
		patch = 0  // Reset PATCH for MINOR version change
	case "PATCH":
		patch++
	default:
		return "", fmt.Errorf("unsupported version type: %s (must be 'major', 'minor', or 'patch')", versionType)
	}

	// Construct new version string with 'v' prefix
	newVersion := fmt.Sprintf("v%d.%d.%d", major, minor, patch)
	return newVersion, nil
}

// GetLatestVersion retrieves the latest semantic version from Git tags.
// This function fetches all tags from remote repository and returns the highest
// semantic version tag found. If no semantic version tags exist, it falls back
// to local tags.
//
// Returns:
//   - string: Latest version in format "vX.Y.Z", or empty string on error
//   - Error handling: Logs errors but doesn't return them to maintain backward compatibility
//
// Process:
//   1. Fetch all tags from remote repository
//   2. Get latest local version using GetLatestLocalVersion()
//   3. Return the version or log errors
func GetLatestVersion() string {
	// Fetch all tags from remote repository to ensure we have the latest versions
	command := "git fetch --tags"
	_, err := RunShell(command)
	if err != nil {
		Errorf("Failed to fetch tags: %v", err)
	}

	// Get the latest version from local tags
	if result, err := GetLatestLocalVersion(); err == nil {
		return result
	} else {
		Errorf("Failed to get latest version: %v", err)
	}

	// Return empty string if all attempts failed
	return ""
}

// GetLatestLocalVersion finds the highest semantic version tag in the local repository.
// It scans all Git tags, filters for valid semantic versions, sorts them, and returns
// the highest version. If no semantic version tags exist, it returns "v1.0.0" as default.
//
// Returns:
//   - string: Latest semantic version tag in format "vX.Y.Z"
//   - error: Error if Git command execution fails
//
// Algorithm:
//   1. Execute 'git tag' to get all local tags
//   2. Filter tags for valid semantic versions using golang.org/x/mod/semver
//   3. Sort versions semantically (not lexicographically)
//   4. Return the highest version or "v1.0.0" if none found
//
// Examples:
//   - If tags are ["v1.0.0", "v1.1.0", "v1.2.0"] → returns "v1.2.0"
//   - If tags are ["alpha", "beta", "v1.0.0"] → returns "v1.0.0"
//   - If tags are ["release-1", "v2.0"] → returns "v1.0.0" (default)
func GetLatestLocalVersion() (string, error) {
	// Execute 'git tag' command to get all local tags
	cmd := exec.Command("git", "tag")
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to execute git tag command: %w", err)
	}

	// Parse output and filter for valid semantic versions
	var versions []string
	lines := strings.Split(out.String(), "\n")

	for _, line := range lines {
		tag := strings.TrimSpace(line)
		if tag != "" && semver.IsValid(tag) {
			versions = append(versions, tag)
		}
	}

	// If no semantic version tags found, return default initial version
	if len(versions) == 0 {
		return "v1.0.0", nil
	}

	// Sort versions semantically using Go's semver package
	// This ensures proper version comparison (e.g., v1.10.0 > v1.2.0)
	sort.Slice(versions, func(i, j int) bool {
		return semver.Compare(versions[i], versions[j]) < 0
	})

	// Return the highest version (last element in sorted array)
	latestVersion := versions[len(versions)-1]
	return latestVersion, nil
}
