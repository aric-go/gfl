/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github-flow/utils"

	"github.com/spf13/cobra"
)

var branchFlag string

// hotfixCmd represents the hotfix command
var hotfixCmd = &cobra.Command{
	Use:   "hotfix",
	Short: "开始一个hotfix分支",
	Run: func(cmd *cobra.Command, args []string) {
		config := utils.ReadConfig()
		featureName := args[0] // 从参数中获取Hotfix名称
		branchName := fmt.Sprintf("hotfix/%s/%s", config.Nickname, featureName)

		// 如果用户传入了 -b 则使用用户的分支，否则使用配置的默认分支
		baseBranch := config.HotfixBaseBranch
		if branchFlag != "" {
			baseBranch = branchFlag
		}
		baseRemoteBranch := fmt.Sprintf("origin/%s", baseBranch)

		// 执行命令: git fetch origin
		if !skipFetch {
			command1 := "git fetch origin"
			if err := utils.RunCommandWithSpin(command1, " 正在同步远程分支...\n"); err != nil {
				return
			}
		}

		// 执行命令: git checkout -b hotfix/aric/new-feature origin/develop
		command2 := fmt.Sprintf("git checkout -b %s %s", branchName, baseRemoteBranch)
		if err := utils.RunCommandWithSpin(command2, " 正在创建Hotfix分支...\n"); err != nil {
			return
		}
		fmt.Printf("✅ 已创建Hotfix分支: %s\n", branchName)
	},
}

func init() {
	hotfixCmd.Flags().StringVarP(&branchFlag, "branch", "b", "", "指定一个自定义的 base 分支")
	hotfixCmd.Flags().BoolVarP(&skipFetch, "skip-fetch", "s", false, "跳过 git fetch 步骤")
	rootCmd.AddCommand(hotfixCmd)
}
