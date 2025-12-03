package cmd

import (
	"gfl/utils"
	"gfl/utils/strings"

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
	Short:   "Rename a branch (local and/or remote)", // Will be updated after strings load
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
	renameCmd.Flags().BoolVarP(&renameLocalFlag, "local", "l", false, strings.GetString("rename", "local_flag"))
	renameCmd.Flags().BoolVarP(&renameRemoteFlag, "remote", "r", false, strings.GetString("rename", "remote_flag"))
	renameCmd.Flags().BoolVarP(&renameDeleteFlag, "delete", "d", false, strings.GetString("rename", "delete_flag"))
	rootCmd.AddCommand(renameCmd)
}
