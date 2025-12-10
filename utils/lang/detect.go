package lang

import (
	"os"
	"strings"
)

// Language represents the supported languages
type Language string

const (
	LanguageZHCN Language = "zh-CN"
	LanguageENUS Language = "en-US"
)

// DetectSystemLanguage attempts to detect the system language
// Checks common environment variables: LANGUAGE, LC_ALL, LC_MESSAGES, LANG
// Returns the language code if supported, otherwise returns empty string
func DetectSystemLanguage() Language {
	// Check common environment variables for language settings
	envVars := []string{
		"LANGUAGE",   // GNU/Linux
		"LC_ALL",     // POSIX
		"LC_MESSAGES", // POSIX
		"LANG",       // Unix/Linux
	}

	for _, envVar := range envVars {
		if value := os.Getenv(envVar); value != "" {
			// Extract the language part (before any . or @)
			lang := strings.Split(value, ".")[0]
			lang = strings.Split(lang, "@")[0]
			lang = strings.TrimSpace(lang)

			// Map common language codes to our supported languages
			switch strings.ToLower(lang) {
			case "en", "en_us", "en-us", "english":
				return LanguageENUS
			case "zh", "zh_cn", "zh-cn", "chinese":
				return LanguageZHCN
			}
		}
	}

	return "" // No supported system language detected
}

// GetLanguagePriority determines the language to use based on priority:
// 1. GFL_LANG environment variable (if set and valid)
// 2. System detected language (if supported)
// 3. Default to Chinese (zh-CN)
func GetLanguagePriority() Language {
	// 1. Check if GFL_LANG environment variable is set
	if lang := os.Getenv("GFL_LANG"); lang != "" {
		// Validate that it's a supported language
		langType := Language(lang)
		if langType == LanguageZHCN || langType == LanguageENUS {
			return langType
		}
	}

	// 2. If no valid GFL_LANG, try to detect system language
	if systemLang := DetectSystemLanguage(); systemLang != "" {
		return systemLang
	}

	// 3. If no system language detected, default to Chinese
	return LanguageZHCN
}