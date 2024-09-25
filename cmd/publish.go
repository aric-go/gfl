package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "发布当前分支",
	Run: func(cmd *cobra.Command, args []string) {
		// 执行命令: git push -u origin HEAD
		if err := exec.Command("git", "push", "-u", "origin", "HEAD").Run(); err != nil {
			fmt.Println("推送当前分支失败:", err)
		} else {
			fmt.Println("已推送当前分支到远程仓库，并设置上游分支")
		}
	},
}

func init() {
	rootCmd.AddCommand(publishCmd)
}
