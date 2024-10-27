package utils

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type YamlConfig struct {
	DevBaseBranch     string   `yaml:"devBaseBranch"`
	HotfixBaseBranch  string   `yaml:"hotfixBaseBranch"`
	ProductionBranch  string   `yaml:"productionBranch"`
	Nickname          string   `yaml:"nickname"`
	Repository        string   `yaml:"repository"`
	ProtectedBranches []string `yaml:"protectedBranches"`
}

func ReadConfig() *YamlConfig {
	data, err := os.ReadFile(".gflow.config.yml")
	if err != nil {
		fmt.Println("读取配置文件失败:", err)
		return nil
	}

	var config YamlConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		fmt.Println("解析配置文件失败:", err)
		return nil
	}

	return &config
}
