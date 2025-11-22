package utils

import (
	"fmt"
	"github.com/spf13/viper"
)

type YamlConfig struct {
	Debug            bool   `yaml:"debug"`
	DevBaseBranch    string `yaml:"devBaseBranch,omitempty"`
	ProductionBranch string `yaml:"productionBranch,omitempty"`
	Nickname         string `yaml:"nickname,omitempty"`
}


func ReadConfig() *YamlConfig {
	// 设置配置文件名（不带扩展名）
	viper.SetConfigName(".gfl.config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".") // 配置文件路径

	// 加载 gfl.config.yml
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading .gfl.config.yml: %v\n", err)
	}

	// 加载 gfl.local.config.yml（如果存在）
	viper.SetConfigName(".gfl.local.config")
	if err := viper.MergeInConfig(); err != nil {
		fmt.Printf("No .gfl.local.config.yml found, using only .gfl.config.yml\n")
	}

	// 最终配置
	var config YamlConfig
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println("解析配置文件失败:", err)
		return nil
	}
	return &config
}

