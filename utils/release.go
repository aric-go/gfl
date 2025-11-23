package utils

// GetLatestReleaseBranch generates the latest release branch name based on the current version.
// This function follows the GFL convention where release branches are named in the format
// "releases/release-{version}" where {version} is the latest semantic version tag.
//
// Returns:
//   - string: The release branch name in format "releases/release-vX.Y.Z"
//
// Examples:
//   - If latest version is "v1.2.3" → returns "releases/release-v1.2.3"
//   - If latest version is "v2.0.0" → returns "releases/release-v2.0.0"
//
// Usage:
//   - Used by release command to determine the target release branch
//   - Used by tag command to locate the appropriate release branch
func GetLatestReleaseBranch() string {
	version := GetLatestVersion()
	return "releases/release-" + version
}
