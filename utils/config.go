package utils

import (
	"os"
	"github.com/spf13/viper"
)

// YamlConfig represents the configuration structure for GFL (GitHub Flow CLI).
// It defines all available configuration options with their YAML tags for serialization.
//
// Configuration priority (highest to lowest):
//   1. Custom config file (GFL_CONFIG_FILE environment variable)
//   2. Local config file (.gfl.config.local.yml)
//   3. Global config file (.gfl.config.yml)
//   4. Default values
type YamlConfig struct {
	// Debug enables verbose logging and debugging output
	Debug bool `yaml:"debug"`

	// DebugSet indicates whether debug was explicitly set in config
	DebugSet bool `yaml:"-"`

	// DevBaseBranch specifies the base branch for feature development (default: "develop")
	DevBaseBranch string `yaml:"devBaseBranch,omitempty"`

	// DevBaseBranchSet indicates whether devBaseBranch was explicitly set
	DevBaseBranchSet bool `yaml:"-"`

	// ProductionBranch specifies the main production branch (default: "main")
	ProductionBranch string `yaml:"productionBranch,omitempty"`

	// ProductionBranchSet indicates whether productionBranch was explicitly set
	ProductionBranchSet bool `yaml:"-"`

	// Nickname is the developer's identifier used in branch naming
	Nickname string `yaml:"nickname,omitempty"`

	// NicknameSet indicates whether nickname was explicitly set in config
	NicknameSet bool `yaml:"-"`

	// FeaturePrefix defines the prefix for feature branches (default: "feature")
	FeaturePrefix string `yaml:"featurePrefix,omitempty"`

	// FeaturePrefixSet indicates whether featurePrefix was explicitly set
	FeaturePrefixSet bool `yaml:"-"`

	// FixPrefix defines the prefix for bug fix branches (default: "fix")
	FixPrefix string `yaml:"fixPrefix,omitempty"`

	// FixPrefixSet indicates whether fixPrefix was explicitly set
	FixPrefixSet bool `yaml:"-"`

	// HotfixPrefix defines the prefix for hotfix branches (default: "hotfix")
	HotfixPrefix string `yaml:"hotfixPrefix,omitempty"`

	// HotfixPrefixSet indicates whether hotfixPrefix was explicitly set
	HotfixPrefixSet bool `yaml:"-"`

	// BranchCaseFormat defines the case format for branch names (default: "original")
	// Supported values: "lower", "upper", "snake", "camel", "pascal", "kebab", "original"
	BranchCaseFormat string `yaml:"branchCaseFormat,omitempty"`

	// BranchCaseFormatSet indicates whether branchCaseFormat was explicitly set
	BranchCaseFormatSet bool `yaml:"-"`
}

// ConfigSource represents a single configuration source with metadata.
// It tracks where each configuration setting originated from.
type ConfigSource struct {
	// Name is a human-readable name for the configuration source
	Name string

	// Path is the file system path to the configuration file
	Path string

	// Config contains the parsed configuration from this source
	Config YamlConfig

	// Exists indicates whether the configuration file actually exists on disk
	Exists bool
}

// ConfigInfo contains comprehensive configuration information including
// the final merged configuration and all individual sources.
type ConfigInfo struct {
	// FinalConfig is the merged configuration after applying all sources
	FinalConfig YamlConfig

	// Sources contains all configuration sources in order of priority
	Sources []ConfigSource
}

// ReadConfig reads and merges configuration from all sources.
// This function maintains backward compatibility by returning only the final config.
//
// Returns:
//   - *YamlConfig: The merged configuration or nil if no config found
func ReadConfig() *YamlConfig {
	info := ReadConfigWithSources()
	return &info.FinalConfig
}

// ReadConfigWithSources reads configuration from all sources and returns detailed information.
// This function provides visibility into where each configuration value originated from,
// which is useful for debugging and the 'gfl config' command.
//
// Configuration sources are loaded in priority order:
//   1. Default values (lowest priority)
//   2. Global config file (.gfl.config.yml)
//   3. Local config file (.gfl.config.local.yml)
//   4. Custom config file (GFL_CONFIG_FILE environment variable)
//
// Returns:
//   - ConfigInfo: Complete configuration information including all sources
func ReadConfigWithSources() ConfigInfo {
	var info ConfigInfo

	// 1. Load default configuration (lowest priority)
	defaultConfig := YamlConfig{
		Debug:            false,
		DevBaseBranch:    "dev",
		ProductionBranch: "main",
		FeaturePrefix:    "feature",
		FixPrefix:        "fix",
		HotfixPrefix:     "hotfix",
		BranchCaseFormat: "original",
	}

	// 2. Load global configuration file
	globalConfigFile := ".gfl.config.yml"
	globalConfig := loadConfigFile(globalConfigFile)
	info.Sources = append(info.Sources, ConfigSource{
		Name:   "Global Config",
		Path:   globalConfigFile,
		Config: globalConfig,
		Exists: fileExists(globalConfigFile),
	})

	// 3. Load local configuration file
	localConfigFile := ".gfl.config.local.yml"
	localConfig := loadConfigFile(localConfigFile)
	info.Sources = append(info.Sources, ConfigSource{
		Name:   "Local Config",
		Path:   localConfigFile,
		Config: localConfig,
		Exists: fileExists(localConfigFile),
	})

	// 4. Load custom configuration file from environment variable
	var customConfig YamlConfig
	customConfigFile := os.Getenv("GFL_CONFIG_FILE")
	if customConfigFile != "" &&
	   customConfigFile != globalConfigFile &&
	   customConfigFile != localConfigFile {
		customConfig = loadConfigFile(customConfigFile)
		info.Sources = append(info.Sources, ConfigSource{
			Name:   "Custom Config",
			Path:   customConfigFile,
			Config: customConfig,
			Exists: fileExists(customConfigFile),
		})
	}

	// 5. Merge configurations in priority order:
	// Default -> Global -> Local -> Custom
	info.FinalConfig = defaultConfig
	mergeConfig(&info.FinalConfig, globalConfig)
	mergeConfig(&info.FinalConfig, localConfig)
	mergeConfig(&info.FinalConfig, customConfig)

	return info
}

// loadConfigFile loads and parses a YAML configuration file.
// It returns an empty config if the file doesn't exist or on parsing errors.
//
// Parameters:
//   - filename: Path to the configuration file
//
// Returns:
//   - YamlConfig: Parsed configuration or empty config on error
func loadConfigFile(filename string) YamlConfig {
	if !fileExists(filename) {
		return YamlConfig{}
	}

	v := viper.New()
	v.SetConfigFile(filename)

	if err := v.ReadInConfig(); err != nil {
		Errorf("Error reading config file %s: %v", filename, err)
		return YamlConfig{}
	}

	var config YamlConfig
	if err := v.Unmarshal(&config); err != nil {
		Errorf("Error parsing config file %s: %v", filename, err)
		return YamlConfig{}
	}

	// Check if fields were explicitly set in the config file
	// This includes cases where fields are set to empty string or false
	if v.IsSet("debug") {
		config.DebugSet = true
	}
	if v.IsSet("devBaseBranch") {
		config.DevBaseBranchSet = true
	}
	if v.IsSet("productionBranch") {
		config.ProductionBranchSet = true
	}
	if v.IsSet("nickname") {
		config.NicknameSet = true
	}
	if v.IsSet("featurePrefix") {
		config.FeaturePrefixSet = true
	}
	if v.IsSet("fixPrefix") {
		config.FixPrefixSet = true
	}
	if v.IsSet("hotfixPrefix") {
		config.HotfixPrefixSet = true
	}
	if v.IsSet("branchCaseFormat") {
		config.BranchCaseFormatSet = true
	}

	return config
}

// mergeConfig merges override configuration into base configuration.
// Only non-empty values from override are applied to base, preserving
// existing values for fields that are not explicitly set.
//
// Parameters:
//   - base: Pointer to the base configuration to be modified
//   - override: Configuration containing values to override
func mergeConfig(base *YamlConfig, override YamlConfig) {
	// Only override fields if they were explicitly set (including empty string or false)
	if override.DebugSet {
		base.Debug = override.Debug
		base.DebugSet = true
	}
	if override.DevBaseBranchSet {
		base.DevBaseBranch = override.DevBaseBranch
		base.DevBaseBranchSet = true
	}
	if override.ProductionBranchSet {
		base.ProductionBranch = override.ProductionBranch
		base.ProductionBranchSet = true
	}
	if override.NicknameSet {
		base.Nickname = override.Nickname
		base.NicknameSet = true
	}
	if override.FeaturePrefixSet {
		base.FeaturePrefix = override.FeaturePrefix
		base.FeaturePrefixSet = true
	}
	if override.FixPrefixSet {
		base.FixPrefix = override.FixPrefix
		base.FixPrefixSet = true
	}
	if override.HotfixPrefixSet {
		base.HotfixPrefix = override.HotfixPrefix
		base.HotfixPrefixSet = true
	}
	if override.BranchCaseFormatSet {
		base.BranchCaseFormat = override.BranchCaseFormat
		base.BranchCaseFormatSet = true
	}
}

// fileExists checks if a file exists at the specified path.
//
// Parameters:
//   - filename: Path to the file to check
//
// Returns:
//   - bool: true if the file exists, false otherwise
func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

