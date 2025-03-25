package cmd

import (
	"fmt"
	"gfl/utils"
	"github.com/spf13/cobra"
)

// syncCmd represents the fetch command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "同步远程仓库到本地仓库/更新所有远程仓库的引用",
	Run: func(cmd *cobra.Command, args []string) {
		if err := utils.RunCommandWithSpin("git fetch origin", " 获取远程仓库中...\n"); err == nil {
			fmt.Printf("✅ 获取远程仓库成功。\n")
		}

		if err := utils.RunCommandWithSpin("git remote update origin --prune", " 获取远程仓库中...\n"); err == nil {
			fmt.Printf("✅ 成功同步远程仓库到本地。\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
