package cmd

import (
	"gfl/utils"
	"gfl/utils/strings"

	"github.com/spf13/cobra"
)

var restoreCmd = &cobra.Command{
	Use:   "restore [path...]",
	Short: "恢复文件到未修改之前",
	Long:  "恢复指定的文件或目录到 HEAD 提交的状态，丢弃所有本地变更（包括已暂存的）",
	Args:  cobra.MinimumNArgs(0), // 可以接受 0 个或多个参数
	Run: func(cmd *cobra.Command, args []string) {
		// get flag confirm
		confirm, _ := cmd.Flags().GetBool("confirm")

		if len(args) == 0 {
			// 没有参数时，作用于当前目录
			utils.RestorePath(".", confirm)
		} else {
			// 有参数时，作用于指定的路径
			for _, path := range args {
				utils.RestorePath(path, confirm)
			}
		}

		if !confirm {
			utils.Info(strings.GetString("restore", "skip_confirm"))
		}
	},
}

func init() {
	// 不需要额外标志，只使用全局的 confirm 标志
	rootCmd.AddCommand(restoreCmd)
}