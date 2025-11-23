package utils

import (
	"os"
	"github.com/spf13/viper"
)

type YamlConfig struct {
	Debug            bool   `yaml:"debug"`
	DevBaseBranch    string `yaml:"devBaseBranch,omitempty"`
	ProductionBranch string `yaml:"productionBranch,omitempty"`
	Nickname         string `yaml:"nickname,omitempty"`
	FeaturePrefix    string `yaml:"featurePrefix,omitempty"`
	FixPrefix        string `yaml:"fixPrefix,omitempty"`
	HotfixPrefix     string `yaml:"hotfixPrefix,omitempty"`
}

type ConfigSource struct {
	Name   string
	Path   string
	Config YamlConfig
	Exists bool
}

type ConfigInfo struct {
	FinalConfig YamlConfig
	Sources     []ConfigSource
}

// ReadConfig 读取配置（保持向后兼容）
func ReadConfig() *YamlConfig {
	info := ReadConfigWithSources()
	return &info.FinalConfig
}

// ReadConfigWithSources 读取配置并返回所有来源信息
func ReadConfigWithSources() ConfigInfo {
	var info ConfigInfo

	// 1. 默认配置（最低优先级）
	defaultConfig := YamlConfig{
		Debug:            false,
		DevBaseBranch:    "dev",
		ProductionBranch: "main",
		FeaturePrefix:    "feature",
		FixPrefix:        "fix",
		HotfixPrefix:     "hotfix",
	}

	// 2. 全局配置文件
	globalConfigFile := ".gfl.config.yml"
	globalConfig := loadConfigFile(globalConfigFile)
	info.Sources = append(info.Sources, ConfigSource{
		Name:   "全局配置",
		Path:   globalConfigFile,
		Config: globalConfig,
		Exists: fileExists(globalConfigFile),
	})

	// 3. 本地配置文件
	localConfigFile := ".gfl.config.local.yml"
	localConfig := loadConfigFile(localConfigFile)
	info.Sources = append(info.Sources, ConfigSource{
		Name:   "本地配置",
		Path:   localConfigFile,
		Config: localConfig,
		Exists: fileExists(localConfigFile),
	})

	// 4. 环境变量配置（仅 GFL_CONFIG_FILE，已在上面处理）
	var envConfig YamlConfig

	// 5. 自定义配置文件（通过 GFL_CONFIG_FILE 环境变量）
	customConfigFile := os.Getenv("GFL_CONFIG_FILE")
	var customConfig YamlConfig
	if customConfigFile != "" && customConfigFile != globalConfigFile && customConfigFile != localConfigFile {
		customConfig = loadConfigFile(customConfigFile)
		info.Sources = append(info.Sources, ConfigSource{
			Name:   "自定义配置",
			Path:   customConfigFile,
			Config: customConfig,
			Exists: fileExists(customConfigFile),
		})
	}

	// 合并配置：默认值 -> 全局配置 -> 本地配置 -> 自定义配置 -> 环境变量
	info.FinalConfig = defaultConfig
	mergeConfig(&info.FinalConfig, globalConfig)
	mergeConfig(&info.FinalConfig, localConfig)
	mergeConfig(&info.FinalConfig, customConfig)
	mergeConfig(&info.FinalConfig, envConfig)

	return info
}

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

	return config
}


func mergeConfig(base *YamlConfig, override YamlConfig) {
	if override.Debug {
		base.Debug = override.Debug
	}
	if override.DevBaseBranch != "" {
		base.DevBaseBranch = override.DevBaseBranch
	}
	if override.ProductionBranch != "" {
		base.ProductionBranch = override.ProductionBranch
	}
	if override.Nickname != "" {
		base.Nickname = override.Nickname
	}
	if override.FeaturePrefix != "" {
		base.FeaturePrefix = override.FeaturePrefix
	}
	if override.FixPrefix != "" {
		base.FixPrefix = override.FixPrefix
	}
	if override.HotfixPrefix != "" {
		base.HotfixPrefix = override.HotfixPrefix
	}
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

