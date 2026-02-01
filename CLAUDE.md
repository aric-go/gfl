# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build and Development Commands

This project uses npm scripts to manage the Go build process:

```bash
# Build the CLI (outputs to dist/gfl)
npm run build

# Run directly without building
npm start [arguments]

# Link binary to /usr/local/bin/gfl
npm run bin:link

# Unlink binary
npm run bin:unlink

# Test the built CLI
./dist/gfl [command]

# Go dependency management
go mod tidy
```

Note: This project uses Go 1.24.0. There are no automated tests currently - testing is done manually.

## Project Architecture

GFL (GitHub Flow CLI) is a command-line tool built with Cobra that implements GitHub Flow workflows. The architecture follows a clean separation between command definitions (in `cmd/`) and business logic utilities (in `utils/`).

### Command Organization

- **Framework**: All commands use [Cobra](https://github.com/spf13/cobra)
- **Command Registration**: Each file in `cmd/` defines a command variable and registers it in `init()`
- **Pattern**: Commands have aliases (short and medium), flags defined in `init()`, and logic in the `Run` field

Example command pattern:
```go
var startCmd = &cobra.Command{
    Use:     "start [feature-name]",
    Short:   strings.GetPath("start.short"),  // Loaded from strings.yml
    Aliases: []string{"s"},
    Args:    cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        // Command implementation
    },
}

func init() {
    rootCmd.AddCommand(startCmd)
    startCmd.Flags().StringVarP(&startBaseBranch, "base", "b", "", "base branch")
}
```

### Configuration System

The configuration system uses **Viper** and supports multiple sources with explicit priority:

**Priority Order** (highest to lowest):
1. Custom config (`GFL_CONFIG_FILE` environment variable)
2. Local config (`.gfl.config.local.yml`)
3. Global config (`.gfl.config.yml`)
4. Default values (hardcoded in `ReadConfigWithSources`)

**Set Flag Pattern**: Each configuration field has a corresponding `*Set` boolean flag to track whether the value was explicitly set. This is critical for distinguishing between "not set" (use default/parent) and "set to false/empty" (explicit value).

Example from `utils/config.go`:
```go
type YamlConfig struct {
    Debug bool `yaml:"debug"`
    DebugSet bool `yaml:"-"`  // Tracks if debug was explicitly set

    DevBaseBranch string `yaml:"devBaseBranch,omitempty"`
    DevBaseBranchSet bool `yaml:"-"`
    // ... other fields follow same pattern
}
```

When adding new configuration fields:
1. Add the field to `YamlConfig` struct with `yaml` tag
2. Add corresponding `*Set bool` field with `yaml:"-"` tag
3. Add `v.IsSet("fieldName")` check in `loadConfigFile`
4. Add merge logic using the Set flag in `mergeConfig`
5. Update `getSource` in `cmd/config.go` for display
6. Update `strings.yml` with translation keys

### Internationalization (i18n)

The project uses a custom string system embedded in `utils/strings/strings.yml`:

- **Dot Notation Access**: `strings.GetPath("category.key")` (uses go-yaml-path library)
- **Languages**: Chinese (zh-CN) and English (en-US)
- **Fallback**: Falls back to Chinese when English translation is missing
- **Loading**: Strings are loaded in `main()` before command execution

Always use `strings.GetPath()` for user-facing text, never hardcode strings.

### Branch Naming System

Branch names are generated using `utils.GenerateBranchName()` with the following patterns:

- Feature: `feature/[nickname]/[name]` or `feature/[name]`
- Bug fix: `fix/[nickname]/[name]` or `fix/[name]`
- Hotfix: `hotfix/[nickname]/[name]` or `hotfix/[name]`

The `nickname` is optional - if not configured, branches are created without it.

**Case Format Support**: The `branchCaseFormat` config option allows automatic case conversion:
- `original` (default) - Keep as-is
- `lower`, `upper`, `snake`, `camel`, `pascal`, `kebab`

Prefix and nickname are always preserved in their original case - only the issue/description part is transformed.

### Command Execution Patterns

Use `utils.RunCommandWithSpin()` for user-facing operations that may take time:
```go
err := utils.RunCommandWithSpin("git push -u origin HEAD", "Pushing branch...")
```

Use `utils.RunCommandWithArgs()` when you have pre-split arguments (safer for user input):
```go
err := utils.RunCommandWithArgs("git", []string{"fetch", "origin"}, "Fetching...")
```

**Debug Mode**: When `debug: true` is set in config, commands display detailed execution info using `utils.debugBox()` which outputs ASCII box formatting via go-box library.

### Git Integration Patterns

The codebase relies on these key utilities:

- `utils.GetLocalBranches()` - Get all local branches
- `utils.IsCleanWorkingDirectory()` - Check if working directory is clean before operations
- `utils.GetCurrentBranch()` - Get current branch name
- `utils.GetRepoOwnerAndName()` - Parse owner/repo from git remote
- `utils.CheckRemoteBranchExists()` - Verify remote branch existence

### Semantic Versioning

Version handling uses `golang.org/x/mod/semver`:
- `utils.GetLatestTag()` - Get latest semver tag from git
- `utils.IncrementVersion()` - Increment version (major/minor/patch)
- Fetches latest tags from remote before version operations

## Common Development Patterns

### Adding a New Command

1. Create `cmd/yourcommand.go`
2. Define command with Use, Short, Aliases, Run
3. Register in `init()` with `rootCmd.AddCommand()`
4. Add string keys to `utils/strings/strings.yml`
5. Use `strings.GetPath()` for all user-facing text
6. Add aliases in Short description: `"功能描述(alias: 短, 长)"`

### Configuration Display

The `gfl config` command displays:
- Final configuration values in a table with color coding
- Source of each value (Global Config, Local Config, Custom Config, Default)
- List of existing config files
- Priority explanation

Color coding: Red (Custom), Yellow (Local), Blue (Global), Cyan (Default)

### Error Handling

Use utility functions for consistent error messaging:
```go
utils.Success("Operation completed")    // ✅ Green
utils.Errorf("Failed: %v", err)        // ERROR: Red
utils.Info("Processing...")             // [gfl] Blue
utils.Warning("Using default")          // WARNING: Yellow
```

## Key Dependencies

- `github.com/spf13/cobra` - CLI framework
- `github.com/spf13/viper` - Configuration management
- `github.com/afeiship/go-yaml-path` - YAML path navigation for i18n
- `github.com/afeiship/go-box` - ASCII box formatting for debug output
- `github.com/briandowns/spinner` - Loading spinners
- `github.com/ettle/strcase` - Case conversion utilities
- `golang.org/x/mod/semver` - Semantic versioning
- `github.com/AlecAivazis/survey/v2` - Interactive prompts

## Configuration Field Reference

Current configuration fields (utils/config.go:16-64):

| Field | Type | Default | Description |
|-------|------|---------|-------------|
| `debug` | bool | false | Enable verbose logging and debug output |
| `devBaseBranch` | string | "dev" | Base branch for feature development |
| `productionBranch` | string | "main" | Main production branch |
| `nickname` | string | "" | Developer identifier for branch naming |
| `featurePrefix` | string | "feature" | Prefix for feature branches |
| `fixPrefix` | string | "fix" | Prefix for bug fix branches |
| `hotfixPrefix` | string | "hotfix" | Prefix for hotfix branches |
| `branchCaseFormat` | string | "original" | Case format for branch names |

All fields must have corresponding `*Set` boolean flags for proper config merging.
