package cmd

import (
	"fmt"
	"gfl/utils"
	"gfl/utils/strings"
	"os/exec"
	str "strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	localFlag  bool
	remoteFlag bool
)

var sweepCmd = &cobra.Command{
	Use:     "sweep [keyword]",
	Aliases: []string{"clean", "rm"},
	Short:   "Clean branches containing specific keywords (alias: clean, rm)",
	Args:    cobra.ExactArgs(1), // 需要一个关键词参数
	Run: func(cmd *cobra.Command, args []string) {
		keyword := args[0]
		// get flag confirm
		confirm, _ := cmd.Flags().GetBool("confirm")

		// 如果没有设置本地或远程标志，打印错误并返回
		if !localFlag && !remoteFlag {
			utils.Error(strings.GetString("sweep", "local_remote_required"))
			return
		}

		if localFlag {
			// 清理本地分支
			cleanLocalBranches(keyword, confirm)
		}

		if remoteFlag {
			// 清理远程分支
			cleanRemoteBranches(keyword, confirm)
		}

		if !confirm {
			utils.Info(strings.GetString("sweep", "skip_confirm"))
		}
	},
}

func cleanLocalBranches(keyword string, confirm bool) {
	// 获取本地分支列表
	branches, err := exec.Command("git", "branch").Output()
	if err != nil {
		utils.Errorf(strings.GetString("sweep", "local_branches_error"), err)
		return
	}

	// 遍历本地分支列表并删除包含关键词的分支
	for _, branch := range str.Split(string(branches), "\n") {
		branch = str.TrimSpace(branch) // 去除空格
		if branch == "" {
			continue // 跳过空行
		}

		if str.Contains(branch, keyword) {
			// 执行命令: git branch -d branch-name
			command := fmt.Sprintf("git branch -d %s", branch)
			if confirm {
				if err := utils.RunCommandWithSpin(command, strings.GetString("sweep", "deleting_local")); err != nil {
					utils.Errorf(strings.GetString("sweep", "delete_local_error"), branch, err)
				} else {
					utils.Successf(strings.GetString("sweep", "delete_local_success"), branch)
				}
			} else {
				logRemove(branch, keyword)
			}
		}
	}
}

func cleanRemoteBranches(keyword string, confirm bool) {
	// 获取远程分支列表
	branches, err := exec.Command("git", "branch", "-r").Output()
	if err != nil {
		utils.Errorf(strings.GetString("sweep", "remote_branches_error"), err)
		return
	}

	// 遍历远程分支列表并删除包含关键词的分支
	for _, branch := range str.Split(string(branches), "\n") {
		branch = str.TrimSpace(branch) // 去除空格
		if branch == "" {
			continue // 跳过空行
		}

		if str.Contains(branch, keyword) {
			// 提取分支名称（去掉远程名）
			remoteBranch := str.TrimPrefix(branch, "origin/")
			command := fmt.Sprintf("git push origin --delete %s", remoteBranch)
			if confirm {
				if err := utils.RunCommandWithSpin(command, strings.GetString("sweep", "deleting_remote")); err != nil {
					utils.Errorf(strings.GetString("sweep", "delete_remote_error"), branch, err)
				} else {
					utils.Successf(strings.GetString("sweep", "delete_remote_success"), branch)
				}
			} else {
				logRemove(branch, keyword)
			}
		}
	}
}

func logRemove(branch string, keyword string) {
	colorBranch := color.GreenString(branch)
	colorKeyword := color.RedString(keyword)
	// list branches without confirm
	utils.Infof(strings.GetString("sweep", "manual_delete"), colorBranch, colorKeyword)
}

func init() {
	sweepCmd.Flags().BoolVarP(&localFlag, "local", "l", false, strings.GetString("sweep", "local_flag"))
	sweepCmd.Flags().BoolVarP(&remoteFlag, "remote", "r", false, strings.GetString("sweep", "remote_flag"))
	rootCmd.AddCommand(sweepCmd)
}
