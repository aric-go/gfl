package cmd

import (
	"fmt"
	"github-flow/utils"
	"github.com/spf13/cobra"
)

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "同步远程仓库到本地",
	Run: func(cmd *cobra.Command, args []string) {
		if err := utils.RunCommandWithSpin("git fetch origin", " 获取远程仓库中...\n"); err == nil {
			fmt.Printf("✅ 获取远程仓库成功。\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
}
