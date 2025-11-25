package cmd

import (
	"fmt"
	"gfl/utils"
	"gfl/utils/strings"

	"github.com/spf13/cobra"
)

// hotfixCmd represents the hotfix command
var hotfixCmd = &cobra.Command{
	Use:     "hotfix [hotfix-name]",
	Aliases: []string{"hf", "hot"},
	Short:   "Start a hotfix branch", // Will be updated after strings load
	Args:    cobra.ExactArgs(1),      // 要求提供一个参数
	Run: func(cmd *cobra.Command, args []string) {
		config := utils.ReadConfig()
		featureName := args[0] // 从参数中获取Hotfix名称
		branchName := utils.GenerateBranchName(config, "hotfix", featureName)

		//baseRemoteBranch := fmt.Sprintf("origin/%s", utils.GetLatestReleaseBranch())

		// 执行命令: git fetch origin
		command1 := "git fetch origin"
		if err := utils.RunCommandWithSpin(command1, strings.GetString("hotfix", "syncing")); err != nil {
			return
		}

		// 执行命令: git checkout -b hotfix/aric/new-feature origin/develop
		command2 := fmt.Sprintf("git checkout -b %s origin/%s", branchName, config.ProductionBranch)
		utils.Infof(strings.GetString("shell", "executing_command"), command2)
		if err := utils.RunCommandWithSpin(command2, strings.GetString("hotfix", "creating")); err != nil {
			return
		}
		utils.Successf(strings.GetString("hotfix", "success"), branchName)
	},
}

func init() {
	rootCmd.AddCommand(hotfixCmd)
}
