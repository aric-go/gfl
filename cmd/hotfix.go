/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github-flow/utils"

	"github.com/spf13/cobra"
)

// hotfixCmd represents the hotfix command
var hotfixCmd = &cobra.Command{
	Use:   "hotfix",
	Short: "开始一个hotfix分支",
	Run: func(cmd *cobra.Command, args []string) {
		config := readConfig()
		featureName := args[0] // 从参数中获取Hotfix名称
		branchName := fmt.Sprintf("hotfix/%s/%s", config.Nickname, featureName)
		baseRemoteBranch := fmt.Sprintf("origin/%s", config.HotfixBaseBranch)

		// 执行命令: git fetch origin develop
		if !skipFetch {
			command1 := fmt.Sprintf("git fetch origin")
			if err := utils.RunCommandWithSpin(command1, " 正在同步远程分支...\n"); err != nil {
				return
			}
		}

		// 执行命令: git checkout -b feature/aric/new-feature origin/develop
		command2 := fmt.Sprintf("git checkout -b %s %s", branchName, baseRemoteBranch)
		if err := utils.RunCommandWithSpin(command2, " 正在创建Hotfix分支...\n"); err != nil {
			return
		}
		fmt.Printf("✅ 已创建Hotfix分支: %s\n", branchName)
	},
}

func init() {
	rootCmd.AddCommand(hotfixCmd)
}
