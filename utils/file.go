package utils

import (
	"fmt"
	"os"
	str "strings"

	"gfl/utils/strings"
	"gopkg.in/yaml.v3"
)

type CreateGflConfigOptions struct {
	Filename     string
	Force        bool
	AddGitIgnore bool
}

func AddGitIgnore() {
	// test has .gitignore file
	if _, err := os.Stat(".gitignore"); os.IsNotExist(err) {
		return
	}

	// add `.gflow.config.yml` to `.gitignore`
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

// CreateGflConfig 创建配置文件
func CreateGflConfig(config YamlConfig, opts CreateGflConfigOptions) error {
	// 检查文件是否存在
	if _, err := os.Stat(opts.Filename); err == nil && !opts.Force {
		return fmt.Errorf(strings.GetString("init", "config_exists_error"), opts.Filename)
	}

	// 创建或覆盖配置文件
	file, err := os.Create(opts.Filename)
	if err != nil {
		return fmt.Errorf(strings.GetString("init", "create_config_error"), err)
	}
	defer file.Close()

	// 序列化配置
	data, err := yaml.Marshal(&config)
	if err != nil {
		return fmt.Errorf(strings.GetString("init", "generate_yaml_error"), err)
	}

	// 写入配置
	if _, err := file.Write(data); err != nil {
		return fmt.Errorf(strings.GetString("init", "write_config_error"), err)
	}

	// 检测 .gitignore 中是否已经存在 Filename 配置
	content, _ := os.ReadFile(".gitignore")
	contentString := string(content)
	if str.Contains(contentString, opts.Filename) {
		Info(strings.GetString("init", "gitignore_skip"))
		return nil
	}

	// 如果需要，添加到 .gitignore
	if opts.AddGitIgnore {
		if f, err := os.OpenFile(".gitignore", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600); err == nil {
			defer f.Close()
			f.WriteString(fmt.Sprintf("\n%s\n", opts.Filename))
		}
	}

	return nil
}

// RemoveEmptyFields 移除配置中的空值字段
func RemoveEmptyFields(config *YamlConfig) *YamlConfig {
	if config == nil {
		return nil
	}

	// 创建新的配置对象
	cleanConfig := &YamlConfig{}

	// 只保留非空值
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
