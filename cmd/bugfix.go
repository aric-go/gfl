package cmd

import (
	"fmt"
	"gfl/utils"
	"gfl/utils/strings"

	"github.com/spf13/cobra"
)

var bugfixCmd = &cobra.Command{
	Use:     "bugfix [bug-name]",
	Short:   strings.GetPath("bugfix.short"),
	Aliases: []string{"b", "fix"},
	Args:    cobra.ExactArgs(1), // 要求提供一个参数
	Run: func(cmd *cobra.Command, args []string) {
		config := utils.ReadConfig()

		if config == nil {
			return
		}

		bugName := args[0]
		branchName := utils.GenerateBranchName(config, "fix", bugName)
		baseRemoteBranch := fmt.Sprintf("origin/%s", config.DevBaseBranch)

		// 执行命令: git fetch origin
		fetchCmd := "git fetch origin"
		if err := utils.RunCommandWithSpin(fetchCmd, strings.GetPath("bugfix.syncing")); err != nil {
			return
		}

		// 执行命令: git checkout -b fix/aric/bug-name origin/develop
		checkoutCmd := fmt.Sprintf("git checkout -b %s %s", branchName, baseRemoteBranch)
		if err := utils.RunCommandWithSpin(checkoutCmd, strings.GetPath("bugfix.creating")); err != nil {
			return
		}

		utils.Successf(strings.GetPath("bugfix.success", "bugfix", branchName))
	},
}

func init() {
	rootCmd.AddCommand(bugfixCmd)
}
