package utils

import (
	"fmt"
	"os"
	str "strings"

	"gfl/utils/strings"
	"gopkg.in/yaml.v3"
)

// CreateGflConfigOptions contains options for creating GFL configuration files.
// This struct provides fine-grained control over the configuration file creation process.
type CreateGflConfigOptions struct {
	// Filename is the name of the configuration file to create
	Filename string

	// Force indicates whether to overwrite an existing configuration file
	Force bool

	// AddGitIgnore determines whether to add the config file to .gitignore
	AddGitIgnore bool
}

// AddGitIgnore adds GFL configuration files to the .gitignore.
// This is a legacy function that adds the hardcoded .gflow.config.yml to .gitignore.
// It silently fails if the .gitignore file doesn't exist or if there are errors.
//
// Note: This function is deprecated. Use CreateGflConfig with AddGitIgnore option instead.
func AddGitIgnore() {
	// Check if .gitignore file exists
	if _, err := os.Stat(".gitignore"); os.IsNotExist(err) {
		return
	}

	// Add `.gflow.config.yml` to `.gitignore`
	f, err := os.OpenFile(".gitignore", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return
	}
	defer f.Close()

	_, err = f.WriteString("\n.gflow.config.yml\n")
	if err != nil {
		return
	}
}

// CreateGflConfig creates a GFL configuration file with the specified options.
// This function handles YAML serialization, file creation, and .gitignore management.
//
// Parameters:
//   - config: The YamlConfig structure to serialize
//   - opts: Options controlling file creation behavior
//
// Returns:
//   - error: Error if file creation fails, nil on success
//
// Process:
//   1. Check if file exists and validate Force option
//   2. Create/open the configuration file
//   3. Serialize configuration to YAML format
//   4. Write YAML data to file
//   5. Optionally add file to .gitignore if requested
//
// Safety Features:
//   - Validates file existence before overwriting (unless Force is true)
//   - Proper file handle management with defer
//   - Checks .gitignore content before adding to avoid duplicates
//   - Uses standardized error messages from internationalization strings
func CreateGflConfig(config YamlConfig, opts CreateGflConfigOptions) error {
	// Step 1: Validate file existence and Force option
	if _, err := os.Stat(opts.Filename); err == nil && !opts.Force {
		return fmt.Errorf(strings.GetPath("init.config_exists_error"), opts.Filename)
	}

	// Step 2: Create or overwrite the configuration file
	file, err := os.Create(opts.Filename)
	if err != nil {
		return fmt.Errorf(strings.GetPath("init.create_config_error"), err)
	}
	defer file.Close()

	// Step 3: Serialize configuration to YAML format
	data, err := yaml.Marshal(&config)
	if err != nil {
		return fmt.Errorf(strings.GetPath("init.generate_yaml_error"), err)
	}

	// Step 4: Write YAML data to file
	if _, err := file.Write(data); err != nil {
		return fmt.Errorf(strings.GetPath("init.write_config_error"), err)
	}

	// Step 5: Check if file is already in .gitignore
	content, _ := os.ReadFile(".gitignore")
	contentString := string(content)
	if str.Contains(contentString, opts.Filename) {
		Info(strings.GetPath("init.gitignore_skip"))
		return nil
	}

	// Step 6: Add to .gitignore if requested
	if opts.AddGitIgnore {
		if f, err := os.OpenFile(".gitignore", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600); err == nil {
			defer f.Close()
			f.WriteString(fmt.Sprintf("\n%s\n", opts.Filename))
		}
		// Silently ignore .gitignore errors as they shouldn't prevent config creation
	}

	return nil
}

// CreateGflConfigFromBytes creates a GFL configuration file from raw bytes.
// This function directly writes the provided byte data to the file without YAML parsing,
// which preserves comments and formatting in the original content.
//
// Parameters:
//   - data: Raw YAML file content as bytes
//   - opts: Options controlling file creation behavior
//
// Returns:
//   - error: Error if file creation fails, nil on success
//
// Use Cases:
//   - Copying template configuration files with comments
//   - Preserving YAML formatting and comments
//   - Creating config files from embedded assets
func CreateGflConfigFromBytes(data []byte, opts CreateGflConfigOptions) error {
	// Step 1: Validate file existence and Force option
	if _, err := os.Stat(opts.Filename); err == nil && !opts.Force {
		return fmt.Errorf(strings.GetPath("init.config_exists_error"), opts.Filename)
	}

	// Step 2: Create or overwrite the configuration file
	file, err := os.Create(opts.Filename)
	if err != nil {
		return fmt.Errorf(strings.GetPath("init.create_config_error"), err)
	}
	defer file.Close()

	// Step 3: Write raw data to file
	if _, err := file.Write(data); err != nil {
		return fmt.Errorf(strings.GetPath("init.write_config_error"), err)
	}

	// Step 4: Check if file is already in .gitignore
	content, _ := os.ReadFile(".gitignore")
	contentString := string(content)
	if str.Contains(contentString, opts.Filename) {
		Info(strings.GetPath("init.gitignore_skip"))
		return nil
	}

	// Step 5: Add to .gitignore if requested
	if opts.AddGitIgnore {
		if f, err := os.OpenFile(".gitignore", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600); err == nil {
			defer f.Close()
			f.WriteString(fmt.Sprintf("\n%s\n", opts.Filename))
		}
	}

	return nil
}

// RemoveEmptyFields creates a clean configuration object by removing empty fields.
// This function filters out zero-value fields to create a minimal configuration
// representation that contains only explicitly set values.
//
// Parameters:
//   - config: The source configuration to clean
//
// Returns:
//   - *YamlConfig: A new configuration object with only non-empty fields
//   - nil: If the input config is nil
//
// Fields that are preserved:
//   - Debug: Only if true
//   - String fields: Only if non-empty
//   - All other fields are filtered out if they are zero values
//
// Use Cases:
//   - Creating minimal configuration files
//   - Filtering out default values during configuration merging
//   - Reducing configuration file size
func RemoveEmptyFields(config *YamlConfig) *YamlConfig {
	if config == nil {
		return nil
	}

	// Create a new clean configuration object
	cleanConfig := &YamlConfig{}

	// Only preserve non-zero values
	if config.Debug {
		cleanConfig.Debug = config.Debug
	}
	if config.DevBaseBranch != "" {
		cleanConfig.DevBaseBranch = config.DevBaseBranch
	}
	if config.ProductionBranch != "" {
		cleanConfig.ProductionBranch = config.ProductionBranch
	}
	if config.Nickname != "" {
		cleanConfig.Nickname = config.Nickname
	}
	if config.FeaturePrefix != "" {
		cleanConfig.FeaturePrefix = config.FeaturePrefix
	}
	if config.FixPrefix != "" {
		cleanConfig.FixPrefix = config.FixPrefix
	}
	if config.HotfixPrefix != "" {
		cleanConfig.HotfixPrefix = config.HotfixPrefix
	}

	return cleanConfig
}
