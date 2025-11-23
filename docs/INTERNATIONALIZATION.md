# GFL CLI Internationalization Guide

This document explains how to use the internationalization system implemented in GFL CLI.

## Overview

The GFL CLI now supports multiple languages through a strings-based internationalization system. The system currently supports:

- **zh-CN**: Chinese (Simplified) - Default language
- **en-US**: English

## How it Works

### 1. Strings Storage

All user-facing strings are stored in `utils/strings.yml` with separate sections for each language:

```yaml
zh-CN:
  root:
    short: "GitHub Flow CLI"
    welcome: "🌈 Welcome to GitHub Flow CLI!"
    confirm_flag: "确认操作"

en-US:
  root:
    short: "GitHub Flow CLI"
    welcome: "🌈 Welcome to GitHub Flow CLI!"
    confirm_flag: "Confirm operation"
```

### 2. Strings Package

The `gfl/utils/strings` package provides functions to load and access strings:

- `LoadStrings()`: Initializes the strings system by loading the YAML file
- `GetString(category, key, args...interface{})`: Retrieves a string with optional formatting
- `SetLanguage(lang Language)`: Sets the current language
- `GetLanguage()`: Gets the current language

### 3. Language Selection

Language can be set via environment variable:

```bash
# Use Chinese (default)
./gfl --help

# Use English
GFL_LANG=en-US ./gfl --help
```

## Adding New Strings

### 1. Add to YAML

Add new strings to both language sections in `utils/strings.yml`:

```yaml
zh-CN:
  new_command:
    short: "新命令描述"
    success: "操作成功: %s"

en-US:
  new_command:
    short: "New command description"
    success: "Operation successful: %s"
```

### 2. Update Go Code

Import the strings package and use `GetString()`:

```go
import "gfl/utils/strings"

// In your command
strings.GetString("new_command", "short")
strings.GetString("new_command", "success", "arg1")
```

### 3. Update Command Descriptions

Add the command to `updateCommandDescriptions()` in `cmd/root.go`:

```go
func updateCommandDescriptions() {
    // ... existing code ...

    if newCommand != nil {
        newCommand.Short = strings.GetString("new_command", "short")
    }
}
```

## Supported String Categories

- `root`: Root command strings
- `init`: Initialize command strings
- `start`: Start feature command strings
- `publish`: Publish branch command strings
- `hotfix`: Hotfix command strings
- `tag`: Tag command strings
- `pr`: Pull request command strings
- `checkout`: Interactive checkout command strings
- `sweep`: Branch cleanup command strings
- `sync`: Repository sync command strings
- `release`: Release command strings
- `config`: Configuration command strings
- `logger`: Logger utility strings
- `shell`: Shell utility strings
- `utils_config`: Configuration utility strings
- `pr_utils`: Pull request utility strings
- `semver`: Semantic versioning strings
- `git`: Git utility strings

## String Format Support

The system supports formatted strings with placeholders:

```yaml
# In utils/strings.yml
success: "Created %s branch: %s"

# In Go code
strings.GetString("start", "success", "feature", "feature/new-login")
```

## Best Practices

1. **Consistent Keys**: Use the same key names across both languages
2. **Descriptive Keys**: Use clear, descriptive category and key names
3. **Parameter Order**: Keep parameter order consistent between languages
4. **Testing**: Test both languages to ensure strings display correctly
5. **Context**: Keep strings contextually appropriate for each language

## Adding New Languages

To add support for a new language:

1. Add a new language section to `utils/strings.yml` with all strings translated
2. Add the language constant to `utils/strings/strings.go`:
   ```go
   const (
       LanguageZHCN Language = "zh-CN"
       LanguageENUS Language = "en-US"
       LanguageJP   Language = "ja-JP"  // New language
   )
   ```

3. Use the new language code:
   ```bash
   GFL_LANG=ja-JP ./gfl --help
   ```

## Migration Progress

The internationalization system is implemented and functional. The following components have been migrated:

- ✅ Root command and flags
- ✅ Start command (complete example)
- ⏳ Other commands (framework ready, needs implementation)

To complete the migration for remaining commands:

1. Import `"gfl/utils/strings"` in each command file
2. Replace hardcoded strings with `strings.GetString()` calls
3. Update command Short/Long descriptions to use placeholder text
4. Add commands to `updateCommandDescriptions()` in `cmd/root.go`

## Testing

Test the internationalization:

```bash
# Test Chinese (default)
./gfl --help
./gfl start test

# Test English
GFL_LANG=en-US ./gfl --help
GFL_LANG=en-US ./gfl start test
```

This system provides a solid foundation for supporting multiple languages in the GFL CLI while maintaining clean, maintainable code.