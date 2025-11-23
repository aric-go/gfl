package utils

import "fmt"

// GetBranchTypePrefix retrieves the branch type prefix from configuration.
// It returns the custom prefix if configured, otherwise falls back to default values.
//
// Parameters:
//   - config: The YAML configuration containing branch type settings
//   - branchType: The type of branch (e.g., "feature", "fix", "hotfix")
//
// Returns:
//   - string: The prefix to be used for the specified branch type
//
// Default prefixes:
//   - feature: "feature"
//   - fix: "fix"
//   - hotfix: "hotfix"
//   - custom: returns the branchType as-is for unknown types
func GetBranchTypePrefix(config *YamlConfig, branchType string) string {
	switch branchType {
	case "feature":
		if config.FeaturePrefix != "" {
			return config.FeaturePrefix
		}
		return "feature"
	case "fix":
		if config.FixPrefix != "" {
			return config.FixPrefix
		}
		return "fix"
	case "hotfix":
		if config.HotfixPrefix != "" {
			return config.HotfixPrefix
		}
		return "hotfix"
	default:
		// For unknown branch types, return the branchType as-is
		return branchType
	}
}

// GenerateBranchName generates a standardized branch name based on configuration.
// The branch naming follows GitHub Flow conventions with optional developer nickname.
//
// Branch naming patterns:
//   - With nickname: {prefix}/{nickname}/{name}
//   - Without nickname: {prefix}/{name}
//
// Parameters:
//   - config: The YAML configuration containing branch naming settings
//   - branchType: The type of branch (e.g., "feature", "fix", "hotfix", "release")
//   - name: The descriptive name for the branch (e.g., "user-authentication")
//
// Returns:
//   - string: The fully formatted branch name
//
// Examples:
//   - With nickname "aric": "feature/aric/user-authentication"
//   - Without nickname: "feature/user-authentication"
//   - Hotfix with nickname: "hotfix/aric/security-fix"
//   - Custom prefix: "feat/bob/payment-integration"
func GenerateBranchName(config *YamlConfig, branchType, name string) string {
	// Get the appropriate prefix for the branch type
	prefix := GetBranchTypePrefix(config, branchType)

	// Include nickname if configured, otherwise use simpler format
	if config.Nickname != "" {
		return fmt.Sprintf("%s/%s/%s", prefix, config.Nickname, name)
	}

	return fmt.Sprintf("%s/%s", prefix, name)
}