package cmd

import (
	"fmt"
	"gfl/utils"
	"gfl/utils/strings"
	str "strings"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:     "start [feature-name]",
	Short:   "Start a new feature (alias: s)", // Will be updated after strings load
	Aliases: []string{"s"},
	Args:    cobra.ExactArgs(1), // 要求提供一个参数
	Run: func(cmd *cobra.Command, args []string) {
		config := utils.ReadConfig()

		if config == nil {
			return
		}

		startName := parseStartName(args[0])
		branchName := utils.GenerateBranchName(config, startName.ActionName, startName.FeatureName)
		baseRemoteBranch := fmt.Sprintf("origin/%s", config.DevBaseBranch)

		// 执行命令: git fetch origin
		fetchCmd := "git fetch origin"
		if err := utils.RunCommandWithSpin(fetchCmd, strings.GetString("start", "syncing")); err != nil {
			return
		}

		// 执行命令: git checkout -b feature/aric/new-feature origin/develop
		checkoutCmd := fmt.Sprintf("git checkout -b %s %s", branchName, baseRemoteBranch)
		if err := utils.RunCommandWithSpin(checkoutCmd, strings.GetString("start", "creating")); err != nil {
			return
		}
		utils.Successf(strings.GetString("start", "success"), startName.ActionName, branchName)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}

type StartName struct {
	ActionName  string
	FeatureName string
}

func parseStartName(name string) *StartName {
	hasColon := str.Contains(name, ":")
	if hasColon {
		parts := str.Split(name, ":")
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
