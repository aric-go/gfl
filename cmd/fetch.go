package cmd

import (
	"fmt"
	"github-flow/utils"
	"github.com/spf13/cobra"
)

const (
	CommandShort   = "同步远程仓库到本地"
	CommandString  = "git fetch origin"
	IngRemote      = " 获取远程仓库中...\n"
	SuccessMessage = "✅ 获取远程仓库成功。\n"
)

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: CommandShort,
	Run: func(cmd *cobra.Command, args []string) {
		if err := utils.RunCommandWithSpin(CommandString, IngRemote); err == nil {
			fmt.Printf(SuccessMessage)
		}
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
}
