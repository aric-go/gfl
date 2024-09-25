package cmd

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os/exec"

	"github.com/spf13/cobra"
	"io/ioutil"
)

var featureName string

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "开始一个新功能",
	Run: func(cmd *cobra.Command, args []string) {
		config := readConfig()
		if config == nil {
			return
		}

		branchName := fmt.Sprintf("feature/%s/%s", config.Nickname, featureName)

		// 执行命令: git fetch origin develop
		if err := exec.Command("git", "fetch", "origin", config.BaseBranch).Run(); err != nil {
			fmt.Println("拉取分支失败:", err)
			return
		}

		// 执行命令: git checkout -b feature/aric/new-feature origin/develop
		if err := exec.Command("git", "checkout", "-b", branchName, config.BaseBranch).Run(); err != nil {
			fmt.Println("创建分支失败:", err)
		} else {
			fmt.Printf("已创建功能分支: %s\n", branchName)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVarP(&featureName, "feature", "f", "", "新功能的名称")
	startCmd.MarkFlagRequired("feature")
}

// 读取配置文件
func readConfig() *YamlConfig {
	data, err := ioutil.ReadFile(".gflow.config.yml")
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
