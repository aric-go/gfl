package cmd

import (
	"gfl/utils"
	"gfl/utils/strings"
	"github.com/spf13/cobra"
)

var publishCmd = &cobra.Command{
	Use:     "publish",
	Aliases: []string{"p"},
	Short:   "Publish current branch (alias: p)", // Will be updated after strings load
	Run: func(cmd *cobra.Command, args []string) {
		// 执行命令: git push -u origin HEAD
		if err := utils.RunCommandWithSpin("git push -u origin HEAD", strings.GetPath("publish.pushing")); err != nil {
			return
		}
		utils.Success(strings.GetPath("publish.success"))
	},
}

func init() {
	rootCmd.AddCommand(publishCmd)
}
