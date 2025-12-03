package cmd

import (
	"gfl/utils"
	"gfl/utils/strings"
	"os"

	"github.com/spf13/cobra"
)

var (
	renameLocalFlag  bool
	renameRemoteFlag bool
	renameDeleteFlag bool
)

var renameCmd = &cobra.Command{
	Use:     "rename [old-branch] [new-branch]",
	Aliases: []string{"mv"},
	Short:   "重命名分支 (支持本地和远程)",
	Args:    cobra.ExactArgs(2), // 需要两个参数：旧分支名和新分支名
	Run: func(cmd *cobra.Command, args []string) {
		oldBranch := args[0]
		newBranch := args[1]
		// get flag confirm
		confirm, _ := cmd.Flags().GetBool("confirm")

		// 如果没有设置本地或远程标志，默认操作本地分支
		if !renameLocalFlag && !renameRemoteFlag {
			renameLocalFlag = true
		}

		if renameLocalFlag {
			// 重命名本地分支
			if err := utils.RenameLocalBranch(oldBranch, newBranch, confirm); err != nil {
				utils.Errorf(err.Error())
				return
			}
		}

		if renameRemoteFlag {
			// 处理远程分支
			if err := utils.HandleRemoteBranch(oldBranch, newBranch, renameDeleteFlag, confirm); err != nil {
				utils.Errorf(err.Error())
				return
			}
		}

		if !confirm {
			utils.Info(strings.GetString("rename", "skip_confirm"))
		}
	},
}

func init() {
	// Simple internationalization based on environment
	lang := os.Getenv("GFL_LANG")
	if lang == "" || lang == "zh-CN" {
		renameCmd.Short = "重命名分支 (支持本地和远程)"
		renameCmd.Flags().BoolVarP(&renameLocalFlag, "local", "l", false, "重命名本地分支")
		renameCmd.Flags().BoolVarP(&renameRemoteFlag, "remote", "r", false, "重命名远程分支")
		renameCmd.Flags().BoolVarP(&renameDeleteFlag, "delete", "d", false, "删除远程旧分支")
	} else {
		renameCmd.Short = "Rename a branch (supports local and remote)"
		renameCmd.Flags().BoolVarP(&renameLocalFlag, "local", "l", false, "Rename local branch")
		renameCmd.Flags().BoolVarP(&renameRemoteFlag, "remote", "r", false, "Rename remote branch")
		renameCmd.Flags().BoolVarP(&renameDeleteFlag, "delete", "d", false, "Delete old remote branch")
	}
	rootCmd.AddCommand(renameCmd)
}
