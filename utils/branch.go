package utils

import (
	"fmt"
	"strings"

	"github.com/ettle/strcase"
)

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

	var branchName string
	// Include nickname if configured, otherwise use simpler format
	if config.Nickname != "" {
		branchName = fmt.Sprintf("%s/%s/%s", prefix, config.Nickname, name)
	} else {
		branchName = fmt.Sprintf("%s/%s", prefix, name)
	}

	// Apply case formatting if configured
	if config.BranchCaseFormat != "" && config.BranchCaseFormat != "original" {
		branchName = FormatBranchName(branchName, config.BranchCaseFormat)
	}

	return branchName
}

// FormatBranchName applies case formatting to a branch name based on configuration.
// It converts only the feature name part (last segment) to the specified format,
// preserving the prefix and nickname.
//
// Parameters:
//   - branchName: The full branch name (e.g., "feature/test/KAT-123" or "feature/KAT-123")
//   - format: The target format ("lower", "upper", "snake", "camel", "pascal", "kebab", "original")
//
// Returns:
//   - string: The formatted branch name
//
// Examples:
//   - "feature/KAT-123" with "lower" -> "feature/kat-123"
//   - "feature/test/KAT-123" with "lower" -> "feature/test/kat-123"
//   - "feature/KAT-123" with "snake" -> "feature/kat_123"
//   - "feature/kat-123" with "camel" -> "feature/kat123"
func FormatBranchName(branchName, format string) string {
	// If format is original or empty, return as-is
	if format == "original" || format == "" {
		return branchName
	}

	// Split branch name into all parts
	parts := strings.Split(branchName, "/")
	if len(parts) < 2 {
		// No prefix found, format the entire string
		return formatString(branchName, format)
	}

	// Format only the last part (feature name), preserve prefix and nickname
	lastIndex := len(parts) - 1
	parts[lastIndex] = formatString(parts[lastIndex], format)

	return strings.Join(parts, "/")
}

// formatString converts a string to the specified case format.
//
// Parameters:
//   - s: The string to format
//   - format: The target format
//
// Returns:
//   - string: The formatted string
func formatString(s, format string) string {
	switch format {
	case "lower":
		return strings.ToLower(s)
	case "upper":
		return strings.ToUpper(s)
	case "snake":
		// Convert kebab-case or other formats to snake_case
		return strcase.ToSnake(s)
	case "camel":
		// Convert to camelCase
		return strcase.ToCamel(s)
	case "pascal":
		// Convert to PascalCase
		return strcase.ToPascal(s)
	case "kebab":
		// Convert to kebab-case
		return strcase.ToKebab(s)
	default:
		// Unknown format, return as-is
		return s
	}
}