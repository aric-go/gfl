package cmd

import (
	"fmt"
	"github-flow/utils"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var (
	localFlag  bool
	remoteFlag bool
)

var sweepCmd = &cobra.Command{
	Use:     "sweep [keyword]",
	Aliases: []string{"clean", "rm"},
	Short:   "清理包含特定关键词的分支(alias: clean, rm)",
	Args:    cobra.ExactArgs(1), // 需要一个关键词参数
	Run: func(cmd *cobra.Command, args []string) {
		keyword := args[0]

		// 如果没有设置本地或远程标志，打印错误并返回
		if !localFlag && !remoteFlag {
			fmt.Println("请至少指定一个 --local 或 --remote 标志")
			return
		}

		if localFlag {
			// 清理本地分支
			cleanLocalBranches(keyword)
		}
		if remoteFlag {
			// 清理远程分支
			cleanRemoteBranches(keyword)
		}
	},
}

func cleanLocalBranches(keyword string) {
	// 获取本地分支列表
	branches, err := exec.Command("git", "branch").Output()
	if err != nil {
		fmt.Println("获取本地分支列表失败:", err)
		return
	}

	// 遍历本地分支列表并删除包含关键词的分支
	for _, branch := range strings.Split(string(branches), "\n") {
		branch = strings.TrimSpace(branch) // 去除空格
		if branch == "" {
			continue // 跳过空行
		}

		if strings.Contains(branch, keyword) {
			// 执行命令: git branch -d branch-name
			command := fmt.Sprintf("git branch -d %s", branch)
			if err := utils.RunCommandWithSpin(command, "🚗 正在删除本地分支\n"); err != nil {
				fmt.Printf("💔 删除本地分支 %s 失败: %s\n", branch, err)
			} else {
				fmt.Printf("✅ 本地分支 %s 删除成功\n", branch)
			}
		}
	}
}

func cleanRemoteBranches(keyword string) {
	// 获取远程分支列表
	branches, err := exec.Command("git", "branch", "-r").Output()
	if err != nil {
		fmt.Println("💔 获取远程分支列表失败:", err)
		return
	}

	// 遍历远程分支列表并删除包含关键词的分支
	for _, branch := range strings.Split(string(branches), "\n") {
		branch = strings.TrimSpace(branch) // 去除空格
		if branch == "" {
			continue // 跳过空行
		}

		if strings.Contains(branch, keyword) {
			// 提取分支名称（去掉远程名）
			remoteBranch := strings.TrimPrefix(branch, "origin/")
			command := fmt.Sprintf("git push origin --delete %s", remoteBranch)
			if err := utils.RunCommandWithSpin(command, "🚗 正在删除远程分支\n"); err != nil {
				fmt.Printf("💔 删除远程分支 %s 失败: %s\n", branch, err)
			} else {
				fmt.Printf("✅ 远程分支 %s 删除成功\n", branch)
			}
		}
	}
}

func init() {
	sweepCmd.Flags().BoolVarP(&localFlag, "local", "l", false, "清理本地分支")
	sweepCmd.Flags().BoolVarP(&remoteFlag, "remote", "r", false, "清理远程分支")
	rootCmd.AddCommand(sweepCmd)
}
