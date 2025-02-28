package cmd

import (
	"fmt"
	"github-flow/utils"
	"github.com/spf13/cobra"
)

var publishCmd = &cobra.Command{
	Use:     "publish",
	Aliases: []string{"p"},
	Short:   "发布当前分支(alias: p)",
	Run: func(cmd *cobra.Command, args []string) {
		// 执行命令: git push -u origin HEAD
		if err := utils.RunCommandWithSpin("git push -u origin HEAD", " 正在推送当前分支到远程仓库 \n"); err != nil {
			return
		}
		fmt.Println("🚗 已推送当前分支到远程仓库，并设置上游分支")
	},
}

func init() {
	rootCmd.AddCommand(publishCmd)
}
