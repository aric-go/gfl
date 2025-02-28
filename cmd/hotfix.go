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
		config := utils.ReadConfig()
		skipFetch, _ := cmd.Flags().GetBool("skip-fetch")
		featureName := args[0] // 从参数中获取Hotfix名称
		branchName := fmt.Sprintf("hotfix/%s/%s", config.Nickname, featureName)

		//baseRemoteBranch := fmt.Sprintf("origin/%s", utils.GetLatestReleaseBranch())

		// 执行命令: git fetch origin
		if !skipFetch {
			command1 := "git fetch origin"
			if err := utils.RunCommandWithSpin(command1, " 正在同步远程分支...\n"); err != nil {
				return
			}
		}

		// 执行命令: git checkout -b hotfix/aric/new-feature origin/develop
		command2 := fmt.Sprintf("git checkout -b %s origin/%s", branchName, config.ProductionBranch)
		fmt.Println("command2:", command2)
		if err := utils.RunCommandWithSpin(command2, " 正在创建Hotfix分支...\n"); err != nil {
			return
		}
		fmt.Printf("✅ 已创建Hotfix分支: %s\n", branchName)
	},
}

func init() {
	rootCmd.AddCommand(hotfixCmd)
}
