package cmd

import (
	"fmt"
	"github-flow/utils"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// prCmd represents the pr command
var prCmd = &cobra.Command{
	Use:   "pr",
	Short: "打开 GitHub 的 PR 页面",
	Run: func(cmd *cobra.Command, args []string) {
		// 读取配置文件获取默认分支和仓库
		config := utils.ReadConfig()
		if config == nil {
			return
		}

		// 获取当前的分支名称
		currentBranch, err := getCurrentBranch()
		if err != nil {
			fmt.Println("无法获取当前分支:", err)
			return
		}

		utils.CreatePr(config.DevBaseBranch, currentBranch)
	},
}

func init() {
	rootCmd.AddCommand(prCmd)
}

// 获取当前的分支名称
func getCurrentBranch() (string, error) {
	// 执行 git 命令获取当前分支
	output, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		return "", err
	}

	// 去除换行符
	return strings.TrimSpace(string(output)), nil
}
