package utils

import (
	"fmt"
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


func ReadConfig() *YamlConfig {
	// 从环境变量获取配置文件路径，如果未设置则使用默认值
	configFile := os.Getenv("GFL_CONFIG_FILE")
	if configFile == "" {
		configFile = ".gfl.config.yml"
	}

	// 设置配置文件完整路径
	viper.SetConfigFile(configFile)

	// 加载配置文件
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file %s: %v\n", configFile, err)
	}

	// 最终配置
	var config YamlConfig
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println("解析配置文件失败:", err)
		return nil
	}
	return &config
}

