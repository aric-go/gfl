package utils

import (
	"fmt"
	"github.com/afeiship/go-ipt"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"os"
)

type PublishItem struct {
	Name   string `yaml:"name"`
	Source string `yaml:"source"`
	Target string `yaml:"target"`
}

type YamlConfig struct {
	Debug            bool          `yaml:"debug"`
	DevBaseBranch    string        `yaml:"devBaseBranch,omitempty"`
	ProductionBranch string        `yaml:"productionBranch,omitempty"`
	Nickname         string        `yaml:"nickname,omitempty"`
	PublishList      []PublishItem `yaml:"publishList,omitempty"`
}

type PublishOption struct {
	Label string
	Value PublishItem
}

func ReadConfig1() *YamlConfig {
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

func IptPublishList(config *YamlConfig) {
	publishList := config.PublishList

	var opts []ipt.Option[PublishOption]
	for _, item := range publishList {
		opts = append(opts, ipt.Option[PublishOption]{
			Label: item.Name,
			Value: PublishOption{
				Label: item.Name,
				Value: item,
			},
		})
	}

	selected, err := ipt.Ipt("What is your favorite color?", opts)
	if err != nil {
		fmt.Println("选择发布项目终止:", err)
		return
	}

	var currentBranch string

	if selected.Value.Source == "current" {
		currentBranch, _ = GetCurrentBranch()
	} else {
		currentBranch = selected.Value.Source
	}

	CreatePr(selected.Value.Target, currentBranch)
}
