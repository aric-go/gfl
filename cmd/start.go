package cmd

import (
	"fmt"
	"github-flow/utils"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"os"
)

var startCmd = &cobra.Command{
	Use:   "start [feature-name]",
	Short: "开始一个新功能",
	Args:  cobra.ExactArgs(1), // 要求提供一个参数
	Run: func(cmd *cobra.Command, args []string) {
		config := readConfig()
		if config == nil {
			return
		}

		featureName := args[0] // 从参数中获取功能名称
		branchName := fmt.Sprintf("feature/%s/%s", config.Nickname, featureName)
		baseRemoteBranch := fmt.Sprintf("origin/%s", config.BaseBranch)

		// 执行命令: git fetch origin develop
		if err := utils.RunCommandWithSpin("git fetch origin", " 拉取远程分支...\n"); err != nil {
			fmt.Println("拉取远程分支失败:", err)
			return
		}

		// 执行命令: git checkout -b feature/aric/new-feature origin/develop
		command := fmt.Sprintf("git checkout -b %s %s", branchName, baseRemoteBranch)
		if err := utils.RunCommandWithSpin(command, " 正在创建功能分支...\n"); err != nil {
			fmt.Println("创建分支失败:", err)
			return
		} else {
			fmt.Printf("已创建功能分支: %s\n", branchName)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}

// 读取配置文件
func readConfig() *YamlConfig {
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
