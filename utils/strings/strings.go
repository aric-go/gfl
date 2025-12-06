package strings

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/afeiship/go-yaml-path"
)

//go:embed strings.yml
var stringsData string

// Language represents the supported languages
type Language string

const (
	LanguageZHCN Language = "zh-CN"
	LanguageENUS Language = "en-US"
)

var (
	// Global YPath instance for the entire strings data
	globalYPath *ypath.YPath

	// Current language (can be set via environment variable or config)
	currentLanguage Language = LanguageZHCN // Default to Chinese
)

// LoadStrings initializes the strings package by loading the embedded strings.yml data using go-yaml-path
func LoadStrings() error {
	// Get the current language from environment variable or default to zh-CN
	if lang := os.Getenv("GFL_LANG"); lang != "" {
		currentLanguage = Language(lang)
	}

	// Parse the embedded YAML data using go-yaml-path
	yp, err := ypath.New([]byte(stringsData))
	if err != nil {
		return fmt.Errorf("failed to parse embedded strings data: %w", err)
	}

	globalYPath = yp
	return nil
}

// SetLanguage sets the current language
func SetLanguage(lang Language) {
	currentLanguage = lang
}

// GetLanguage returns the current language
func GetLanguage() Language {
	return currentLanguage
}

// GetPath returns a specific string by direct dot-path notation (e.g., "rename.local_flag")
// Uses language.path format with automatic fallback to zh-CN
func GetPath(path string, args ...any) string {
	if globalYPath == nil {
		return ""
	}

	// Construct the full path using dot notation: language.path
	fullPath := fmt.Sprintf("%s.%s", currentLanguage, path)

	// Get the string value
	value := globalYPath.GetString(fullPath)

	// If the value is not found, try fallback to zh-CN
	if value == "" && currentLanguage != LanguageZHCN {
		fallbackPath := fmt.Sprintf("%s.%s", LanguageZHCN, path)
		value = globalYPath.GetString(fallbackPath)
	}

	// If we have arguments and the value contains formatting placeholders, use fmt.Sprintf
	if len(args) > 0 && value != "" {
		return fmt.Sprintf(value, args...)
	}

	return value
}
