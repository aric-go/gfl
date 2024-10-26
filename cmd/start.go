package cmd

import (
	"fmt"
	"github-flow/utils"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

var (
	skipFetch bool
)

var startCmd = &cobra.Command{
	Use:     "start [feature-name]",
	Short:   "开始一个新功能(alias: s)",
	Aliases: []string{"s"},
	Args:    cobra.ExactArgs(1), // 要求提供一个参数
	Run: func(cmd *cobra.Command, args []string) {
		config := readConfig()
		if config == nil {
			return
		}

		startName := parseStartName(args[0])
		branchName := fmt.Sprintf("%s/%s/%s", startName.ActionName, config.Nickname, startName.FeatureName)
		baseRemoteBranch := fmt.Sprintf("origin/%s", config.DevBaseBranch)

		// 执行命令: git fetch origin develop
		if !skipFetch {
			command1 := fmt.Sprintf("git fetch origin")
			if err := utils.RunCommandWithSpin(command1, " 正在同步远程分支...\n"); err != nil {
				return
			}
		}

		// 执行命令: git checkout -b feature/aric/new-feature origin/develop
		command2 := fmt.Sprintf("git checkout -b %s %s", branchName, baseRemoteBranch)
		if err := utils.RunCommandWithSpin(command2, " 正在创建功能分支...\n"); err != nil {
			return
		}
		fmt.Printf("✅ 已创建功能分支: %s\n", branchName)
	},
}

func init() {
	// add --skip-fetch flag to start command
	startCmd.Flags().BoolVarP(&skipFetch, "skip-fetch", "s", false, "跳过拉取远程分支步骤")
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

type StartName struct {
	ActionName  string
	FeatureName string
}

func parseStartName(name string) *StartName {
	hasColon := strings.Contains(name, ":")
	if hasColon {
		parts := strings.Split(name, ":")
		return &StartName{
			ActionName:  parts[0],
			FeatureName: parts[1],
		}
	} else {
		return &StartName{
			ActionName:  "feature",
			FeatureName: name,
		}
	}
}
