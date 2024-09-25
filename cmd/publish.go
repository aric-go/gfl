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
		// 获取当前分支名称
		currentBranch, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
		if err != nil {
			fmt.Println("无法获取当前分支:", err)
			return
		}
		branchName := string(currentBranch)
		branchName = branchName[:len(branchName)-1] // 移除末尾换行符

		// 执行命令: git push origin -u feature/aric/new-feature
		_, err = exec.Command("git", "push", "origin", "-u", branchName).Output()
		if err != nil {
			fmt.Println("推送失败:", err)
		} else {
			fmt.Printf("已推送当前分支: %s 到远程仓库\n", branchName)
		}
	},
}

func init() {
	rootCmd.AddCommand(publishCmd)
}
