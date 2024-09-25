package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "发布功能分支",
	Run: func(cmd *cobra.Command, args []string) {
		config := readConfig()
		if config == nil {
			return
		}

		branchName := fmt.Sprintf("feature/%s/%s", config.Nickname, featureName)

		// 执行命令: git push origin feature/aric/new-feature
		if err := exec.Command("git", "push", "origin", branchName).Run(); err != nil {
			fmt.Println("推送分支失败:", err)
		} else {
			fmt.Printf("已推送分支: %s 到远程仓库\n", branchName)
		}
	},
}

func init() {
	rootCmd.AddCommand(publishCmd)
}
