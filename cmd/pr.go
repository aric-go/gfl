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
		config := readConfig()
		if config == nil {
			return
		}

		// 获取当前的分支名称
		currentBranch, err := getCurrentBranch()
		if err != nil {
			fmt.Println("无法获取当前分支:", err)
			return
		}

		// 生成 GitHub PR URL
		// @example: https://github.com/applyai-dev/applyai-frontend/compare/dev...feature/aric/gogogo?expand=1
		// https://github.com/applyai-dev/applyai-frontend/compare/${baseBranch}...${headBranch}?expand=1
		prURL := fmt.Sprintf("https://github.com/%s/compare/%s...%s?expand=1", config.Repository, config.DevBaseBranch, currentBranch)

		// 打开浏览器
		utils.CreatePr(prURL)
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
