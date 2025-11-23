package strings

import (
	_ "embed"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

//go:embed strings.yml
var stringsData string

// Language represents the supported languages
type Language string

const (
	LanguageZHCN Language = "zh-CN"
	LanguageENUS Language = "en-US"
)

// StringData represents the entire strings structure
type StringData struct {
	Root         RootStrings         `yaml:"root"`
	Init         InitStrings         `yaml:"init"`
	Start        StartStrings        `yaml:"start"`
	Publish      PublishStrings      `yaml:"publish"`
	Hotfix       HotfixStrings       `yaml:"hotfix"`
	Tag          TagStrings          `yaml:"tag"`
	PR           PRStrings           `yaml:"pr"`
	Checkout     CheckoutStrings     `yaml:"checkout"`
	Sweep        SweepStrings        `yaml:"sweep"`
	Sync         SyncStrings         `yaml:"sync"`
	Release      ReleaseStrings      `yaml:"release"`
	Config       ConfigStrings       `yaml:"config"`
	Logger       LoggerStrings       `yaml:"logger"`
	Shell        ShellStrings        `yaml:"shell"`
	UtilsConfig  ConfigUtilsStrings  `yaml:"utils_config"`
	PRUtils      PRUtilsStrings      `yaml:"pr_utils"`
	Semver       SemverStrings       `yaml:"semver"`
	Git          GitStrings          `yaml:"git"`
}

// Root strings
type RootStrings struct {
	Short       string `yaml:"short"`
	Welcome     string `yaml:"welcome"`
	ConfirmFlag string `yaml:"confirm_flag"`
}

// Init strings
type InitStrings struct {
	Short              string `yaml:"short"`
	ForceFlag          string `yaml:"force_flag"`
	NicknameFlag       string `yaml:"nickname_flag"`
	ConfigExistsError  string `yaml:"config_exists_error"`
	CreateConfigError  string `yaml:"create_config_error"`
	GenerateYAMLError  string `yaml:"generate_yaml_error"`
	WriteConfigError   string `yaml:"write_config_error"`
	GitignoreSkip      string `yaml:"gitignore_skip"`
}

// Start strings
type StartStrings struct {
	Short   string `yaml:"short"`
	Syncing string `yaml:"syncing"`
	Creating string `yaml:"creating"`
	Success string `yaml:"success"`
}

// Publish strings
type PublishStrings struct {
	Short    string `yaml:"short"`
	Pushing  string `yaml:"pushing"`
	Success  string `yaml:"success"`
}

// Hotfix strings
type HotfixStrings struct {
	Short    string `yaml:"short"`
	Syncing  string `yaml:"syncing"`
	Creating string `yaml:"creating"`
	Success  string `yaml:"success"`
}

// Tag strings
type TagStrings struct {
	Short             string `yaml:"short"`
	PreviousVersion   string `yaml:"previous_version"`
	NewVersion        string `yaml:"new_version"`
	Step1             string `yaml:"step1"`
	Step2             string `yaml:"step2"`
	Step3             string `yaml:"step3"`
	Step4             string `yaml:"step4"`
	Step5             string `yaml:"step5"`
	ReleaseSuccess    string `yaml:"release_success"`
	GHNotInstalled    string `yaml:"gh_not_installed"`
	TypeFlag          string `yaml:"type_flag"`
}

// PR strings
type PRStrings struct {
	Short               string `yaml:"short"`
	SyncFailed          string `yaml:"sync_failed"`
	BrowserError        string `yaml:"browser_error"`
	CurrentBranchError  string `yaml:"current_branch_error"`
	SyncFlag            string `yaml:"sync_flag"`
	OpenFlag            string `yaml:"open_flag"`
}

// Checkout strings
type CheckoutStrings struct {
	Short        string `yaml:"short"`
	ChooseBranch string `yaml:"choose_branch"`
}

// Sweep strings
type SweepStrings struct {
	Short                   string `yaml:"short"`
	LocalRemoteRequired     string `yaml:"local_remote_required"`
	SkipConfirm             string `yaml:"skip_confirm"`
	LocalBranchesError      string `yaml:"local_branches_error"`
	DeletingLocal           string `yaml:"deleting_local"`
	DeleteLocalError        string `yaml:"delete_local_error"`
	DeleteLocalSuccess      string `yaml:"delete_local_success"`
	RemoteBranchesError     string `yaml:"remote_branches_error"`
	DeletingRemote          string `yaml:"deleting_remote"`
	DeleteRemoteError       string `yaml:"delete_remote_error"`
	DeleteRemoteSuccess     string `yaml:"delete_remote_success"`
	ManualDelete            string `yaml:"manual_delete"`
	LocalFlag               string `yaml:"local_flag"`
	RemoteFlag              string `yaml:"remote_flag"`
}

// Sync strings
type SyncStrings struct {
	Short        string `yaml:"short"`
	Fetching     string `yaml:"fetching"`
	FetchSuccess string `yaml:"fetch_success"`
	Updating     string `yaml:"updating"`
	SyncSuccess  string `yaml:"sync_success"`
}

// Release strings
type ReleaseStrings struct {
	Short          string `yaml:"short"`
	PreviousVersion string `yaml:"previous_version"`
	NewVersion     string `yaml:"new_version"`
	Step1          string `yaml:"step1"`
	Step2          string `yaml:"step2"`
	Step3          string `yaml:"step3"`
	ReleaseSuccess string `yaml:"release_success"`
	TypeFlag       string `yaml:"type_flag"`
	HotfixFlag     string `yaml:"hotfix_flag"`
}

// Config strings
type ConfigStrings struct {
	Short                   string `yaml:"short"`
	Long                    string `yaml:"long"`
	Title                   string `yaml:"title"`
	ConfigKey               string `yaml:"config_key"`
	FinalValue              string `yaml:"final_value"`
	Source                  string `yaml:"source"`
	DefaultValue            string `yaml:"default_value"`
	CustomConfig            string `yaml:"custom_config"`
	LocalConfig             string `yaml:"local_config"`
	GlobalConfig            string `yaml:"global_config"`
	DebugMode               string `yaml:"debug_mode"`
	DevelopBaseBranch       string `yaml:"develop_base_branch"`
	ProductionBranch        string `yaml:"production_branch"`
	Nickname                string `yaml:"nickname"`
	FeaturePrefix           string `yaml:"feature_prefix"`
	FixPrefix               string `yaml:"fix_prefix"`
	HotfixPrefix            string `yaml:"hotfix_prefix"`
	ExampleFeatureBranch    string `yaml:"example_feature_branch"`
	ConfigSourcesTitle      string `yaml:"config_sources_title"`
	CustomConfigFile        string `yaml:"custom_config_file"`
	PriorityTitle           string `yaml:"priority_title"`
	PriorityCustom          string `yaml:"priority_custom"`
	PriorityLocal           string `yaml:"priority_local"`
	PriorityGlobal          string `yaml:"priority_global"`
	PriorityDefault         string `yaml:"priority_default"`
}

// Logger strings
type LoggerStrings struct {
	Error   string `yaml:"error"`
	Warning string `yaml:"warning"`
	Success string `yaml:"success"`
}

// Shell strings
type ShellStrings struct {
	ConfigReadError  string `yaml:"config_read_error"`
	ExecutingCommand string `yaml:"executing_command"`
	CommandFailed    string `yaml:"command_failed"`
	ExecuteFailed    string `yaml:"execute_failed"`
}

// Config strings (utils/config)
type ConfigUtilsStrings struct {
	GlobalName string `yaml:"global_name"`
	LocalName  string `yaml:"local_name"`
	CustomName string `yaml:"custom_name"`
}

// PR utils strings
type PRUtilsStrings struct {
	BrowserError     string `yaml:"browser_error"`
	PRPageOpened     string `yaml:"pr_page_opened"`
	SyncingBranches  string `yaml:"syncing_branches"`
	DirtyWorkingDir  string `yaml:"dirty_working_dir"`
	CurrentBranchError string `yaml:"current_branch_error"`
	Executing        string `yaml:"executing"`
	CommandFailed    string `yaml:"command_failed"`
	ErrorOutput      string `yaml:"error_output"`
	RollingBack      string `yaml:"rolling_back"`
	Output           string `yaml:"output"`
	SyncSuccess      string `yaml:"sync_success"`
}

// Semver strings
type SemverStrings struct {
	InvalidVersion       string `yaml:"invalid_version"`
	VersionFormatError   string `yaml:"version_format_error"`
	MajorConvertError    string `yaml:"major_convert_error"`
	MinorConvertError    string `yaml:"minor_convert_error"`
	PatchConvertError    string `yaml:"patch_convert_error"`
	InvalidVersionPart   string `yaml:"invalid_version_part"`
	CommandFailed        string `yaml:"command_failed"`
}

// Git strings
type GitStrings struct {
	InvalidURLFormat    string `yaml:"invalid_url_format"`
	UnsupportedURLFormat string `yaml:"unsupported_url_format"`
}

var (
	// Global strings data map
	stringsDataMap map[Language]*StringData

	// Current language (can be set via environment variable or config)
	currentLanguage Language = LanguageZHCN // Default to Chinese
)

// LoadStrings initializes the strings package by loading the embedded strings.yml data
func LoadStrings() error {
	// Get the current language from environment variable or default to zh-CN
	if lang := os.Getenv("GFL_LANG"); lang != "" {
		currentLanguage = Language(lang)
	}

	stringsDataMap = make(map[Language]*StringData)

	// Parse the embedded YAML data
	var allStrings map[string]*StringData
	if err := yaml.Unmarshal([]byte(stringsData), &allStrings); err != nil {
		return fmt.Errorf("failed to parse embedded strings data: %w", err)
	}

	// Store strings for each language
	for langKey, strings := range allStrings {
		stringsDataMap[Language(langKey)] = strings
	}

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

// GetStrings returns the strings data for the current language
func GetStrings() *StringData {
	if strings, ok := stringsDataMap[currentLanguage]; ok {
		return strings
	}

	// Fallback to zh-CN if current language is not available
	if strings, ok := stringsDataMap[LanguageZHCN]; ok {
		return strings
	}

	// Return empty strings if nothing is available
	return &StringData{}
}

// GetString returns a specific string by key (for generic access)
func GetString(category, key string, args ...interface{}) string {
	strings := GetStrings()

	// Use reflection or a map-based approach to access the specific string
	// For simplicity, we'll use a switch statement
	switch category {
	case "root":
		switch key {
		case "short":
			return strings.Root.Short
		case "welcome":
			return fmt.Sprintf(strings.Root.Welcome, args...)
		case "confirm_flag":
			return strings.Root.ConfirmFlag
		}
	case "init":
		switch key {
		case "short":
			return strings.Init.Short
		case "force_flag":
			return strings.Init.ForceFlag
		case "nickname_flag":
			return strings.Init.NicknameFlag
		case "config_exists_error":
			return fmt.Sprintf(strings.Init.ConfigExistsError, args...)
		case "create_config_error":
			return fmt.Sprintf(strings.Init.CreateConfigError, args...)
		case "generate_yaml_error":
			return fmt.Sprintf(strings.Init.GenerateYAMLError, args...)
		case "write_config_error":
			return fmt.Sprintf(strings.Init.WriteConfigError, args...)
		case "gitignore_skip":
			return strings.Init.GitignoreSkip
		}
	case "start":
		switch key {
		case "short":
			return strings.Start.Short
		case "syncing":
			return strings.Start.Syncing
		case "creating":
			return strings.Start.Creating
		case "success":
			return fmt.Sprintf(strings.Start.Success, args...)
		}
	case "publish":
		switch key {
		case "short":
			return strings.Publish.Short
		case "pushing":
			return strings.Publish.Pushing
		case "success":
			return strings.Publish.Success
		}
	case "hotfix":
		switch key {
		case "short":
			return strings.Hotfix.Short
		case "syncing":
			return strings.Hotfix.Syncing
		case "creating":
			return strings.Hotfix.Creating
		case "success":
			return fmt.Sprintf(strings.Hotfix.Success, args...)
		}
	case "tag":
		switch key {
		case "short":
			return strings.Tag.Short
		case "previous_version":
			return fmt.Sprintf(strings.Tag.PreviousVersion, args...)
		case "new_version":
			return fmt.Sprintf(strings.Tag.NewVersion, args...)
		case "step1":
			return strings.Tag.Step1
		case "step2":
			return strings.Tag.Step2
		case "step3":
			return strings.Tag.Step3
		case "step4":
			return strings.Tag.Step4
		case "step5":
			return strings.Tag.Step5
		case "release_success":
			return fmt.Sprintf(strings.Tag.ReleaseSuccess, args...)
		case "gh_not_installed":
			return strings.Tag.GHNotInstalled
		case "type_flag":
			return strings.Tag.TypeFlag
		}
	case "pr":
		switch key {
		case "short":
			return strings.PR.Short
		case "sync_failed":
			return strings.PR.SyncFailed
		case "browser_error":
			return fmt.Sprintf(strings.PR.BrowserError, args...)
		case "current_branch_error":
			return fmt.Sprintf(strings.PR.CurrentBranchError, args...)
		case "sync_flag":
			return strings.PR.SyncFlag
		case "open_flag":
			return strings.PR.OpenFlag
		}
	case "checkout":
		switch key {
		case "short":
			return strings.Checkout.Short
		case "choose_branch":
			return strings.Checkout.ChooseBranch
		}
	case "sweep":
		switch key {
		case "short":
			return strings.Sweep.Short
		case "local_remote_required":
			return strings.Sweep.LocalRemoteRequired
		case "skip_confirm":
			return strings.Sweep.SkipConfirm
		case "local_branches_error":
			return fmt.Sprintf(strings.Sweep.LocalBranchesError, args...)
		case "deleting_local":
			return strings.Sweep.DeletingLocal
		case "delete_local_error":
			return fmt.Sprintf(strings.Sweep.DeleteLocalError, args...)
		case "delete_local_success":
			return fmt.Sprintf(strings.Sweep.DeleteLocalSuccess, args...)
		case "remote_branches_error":
			return fmt.Sprintf(strings.Sweep.RemoteBranchesError, args...)
		case "deleting_remote":
			return strings.Sweep.DeletingRemote
		case "delete_remote_error":
			return fmt.Sprintf(strings.Sweep.DeleteRemoteError, args...)
		case "delete_remote_success":
			return fmt.Sprintf(strings.Sweep.DeleteRemoteSuccess, args...)
		case "manual_delete":
			return fmt.Sprintf(strings.Sweep.ManualDelete, args...)
		case "local_flag":
			return strings.Sweep.LocalFlag
		case "remote_flag":
			return strings.Sweep.RemoteFlag
		}
	case "sync":
		switch key {
		case "short":
			return strings.Sync.Short
		case "fetching":
			return strings.Sync.Fetching
		case "fetch_success":
			return strings.Sync.FetchSuccess
		case "updating":
			return strings.Sync.Updating
		case "sync_success":
			return strings.Sync.SyncSuccess
		}
	case "release":
		switch key {
		case "short":
			return strings.Release.Short
		case "previous_version":
			return fmt.Sprintf(strings.Release.PreviousVersion, args...)
		case "new_version":
			return fmt.Sprintf(strings.Release.NewVersion, args...)
		case "step1":
			return strings.Release.Step1
		case "step2":
			return strings.Release.Step2
		case "step3":
			return strings.Release.Step3
		case "release_success":
			return fmt.Sprintf(strings.Release.ReleaseSuccess, args...)
		case "type_flag":
			return strings.Release.TypeFlag
		case "hotfix_flag":
			return strings.Release.HotfixFlag
		}
	case "config":
		switch key {
		case "short":
			return strings.Config.Short
		case "long":
			return strings.Config.Long
		case "title":
			return strings.Config.Title
		case "config_key":
			return strings.Config.ConfigKey
		case "final_value":
			return strings.Config.FinalValue
		case "source":
			return strings.Config.Source
		case "default_value":
			return strings.Config.DefaultValue
		case "custom_config":
			return strings.Config.CustomConfig
		case "local_config":
			return strings.Config.LocalConfig
		case "global_config":
			return strings.Config.GlobalConfig
		case "debug_mode":
			return strings.Config.DebugMode
		case "develop_base_branch":
			return strings.Config.DevelopBaseBranch
		case "production_branch":
			return strings.Config.ProductionBranch
		case "nickname":
			return strings.Config.Nickname
		case "feature_prefix":
			return strings.Config.FeaturePrefix
		case "fix_prefix":
			return strings.Config.FixPrefix
		case "hotfix_prefix":
			return strings.Config.HotfixPrefix
		case "example_feature_branch":
			return strings.Config.ExampleFeatureBranch
		case "config_sources_title":
			return strings.Config.ConfigSourcesTitle
		case "custom_config_file":
			return fmt.Sprintf(strings.Config.CustomConfigFile, args...)
		case "priority_title":
			return strings.Config.PriorityTitle
		case "priority_custom":
			return strings.Config.PriorityCustom
		case "priority_local":
			return strings.Config.PriorityLocal
		case "priority_global":
			return strings.Config.PriorityGlobal
		case "priority_default":
			return strings.Config.PriorityDefault
		}
	case "utils_config":
		switch key {
		case "global_name":
			return strings.UtilsConfig.GlobalName
		case "local_name":
			return strings.UtilsConfig.LocalName
		case "custom_name":
			return strings.UtilsConfig.CustomName
		}
	case "logger":
		switch key {
		case "error":
			return strings.Logger.Error
		case "warning":
			return strings.Logger.Warning
		case "success":
			return strings.Logger.Success
		}
	case "shell":
		switch key {
		case "config_read_error":
			return strings.Shell.ConfigReadError
		case "executing_command":
			return fmt.Sprintf(strings.Shell.ExecutingCommand, args...)
		case "command_failed":
			return fmt.Sprintf(strings.Shell.CommandFailed, args...)
		case "execute_failed":
			return fmt.Sprintf(strings.Shell.ExecuteFailed, args...)
		}
	case "pr_utils":
		switch key {
		case "browser_error":
			return fmt.Sprintf(strings.PRUtils.BrowserError, args...)
		case "pr_page_opened":
			return fmt.Sprintf(strings.PRUtils.PRPageOpened, args...)
		case "syncing_branches":
			return fmt.Sprintf(strings.PRUtils.SyncingBranches, args...)
		case "dirty_working_dir":
			return strings.PRUtils.DirtyWorkingDir
		case "current_branch_error":
			return fmt.Sprintf(strings.PRUtils.CurrentBranchError, args...)
		case "executing":
			return fmt.Sprintf(strings.PRUtils.Executing, args...)
		case "command_failed":
			return fmt.Sprintf(strings.PRUtils.CommandFailed, args...)
		case "error_output":
			return fmt.Sprintf(strings.PRUtils.ErrorOutput, args...)
		case "rolling_back":
			return fmt.Sprintf(strings.PRUtils.RollingBack, args...)
		case "output":
			return fmt.Sprintf(strings.PRUtils.Output, args...)
		case "sync_success":
			return fmt.Sprintf(strings.PRUtils.SyncSuccess, args...)
		}
	case "semver":
		switch key {
		case "invalid_version":
			return fmt.Sprintf(strings.Semver.InvalidVersion, args...)
		case "version_format_error":
			return fmt.Sprintf(strings.Semver.VersionFormatError, args...)
		case "major_convert_error":
			return fmt.Sprintf(strings.Semver.MajorConvertError, args...)
		case "minor_convert_error":
			return fmt.Sprintf(strings.Semver.MinorConvertError, args...)
		case "patch_convert_error":
			return fmt.Sprintf(strings.Semver.PatchConvertError, args...)
		case "invalid_version_part":
			return fmt.Sprintf(strings.Semver.InvalidVersionPart, args...)
		case "command_failed":
			return fmt.Sprintf(strings.Semver.CommandFailed, args...)
		}
	case "git":
		switch key {
		case "invalid_url_format":
			return fmt.Sprintf(strings.Git.InvalidURLFormat, args...)
		case "unsupported_url_format":
			return fmt.Sprintf(strings.Git.UnsupportedURLFormat, args...)
		}
	}

	// Return empty string if not found
	return ""
}

