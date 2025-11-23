package utils

import (
	"fmt"
	"os/exec"
	"strings"
	"github.com/pkg/browser"
)

func CreatePr(base string, head string) {
	repo, _ := GetRepository()

	// 生成 GitHub PR URL
	// @example: https://github.com/applyai-dev/applyai-frontend/compare/dev...feature/aric/gogogo?expand=1
	// https://github.com/applyai-dev/applyai-frontend/compare/${baseBranch}...${headBranch}?expand=1
	url := fmt.Sprintf("https://github.com/%s/compare/%s...%s?expand=1", repo, base, head)

	err := browser.OpenURL(url)
	if err != nil {
		Errorf("无法打开浏览器: %v", err)
	} else {
		Infof("已打开 PR 页面: %s", url)
	}
}

// SyncProductionToDev 同步生产分支到开发分支
func SyncProductionToDev(productionBranch, devBranch string) bool {
	Infof("正在同步生产分支 %s 到开发分支 %s...", productionBranch, devBranch)

	// 1. 检查工作目录是否干净
	if !isWorkingDirectoryClean() {
		Errorf("工作目录不干净，请先提交或暂存当前的更改")
		return false
	}

	// 2. 保存当前分支
	currentBranch, err := GetCurrentBranch()
	if err != nil {
		Errorf("无法获取当前分支: %v", err)
		return false
	}

	// 3. 执行同步操作
	commands := [][]string{
		{"git", "fetch", "origin"},
		{"git", "checkout", devBranch},
		{"git", "pull", "origin", devBranch},
		{"git", "merge", "origin/" + productionBranch},
		{"git", "push", "origin", devBranch},
	}

	for i, cmd := range commands {
		Infof("执行: %s", strings.Join(cmd, " "))
		output, err := exec.Command(cmd[0], cmd[1:]...).CombinedOutput()
		if err != nil {
			Errorf("命令执行失败: %v", err)
			Errorf("错误输出: %s", string(output))

			// 尝试回滚到原始分支
			Warningf("正在回滚到原始分支 %s...", currentBranch)
			if i > 0 { // 如果已经切换了分支
				exec.Command("git", "checkout", currentBranch).CombinedOutput()
			}
			return false
		}
		if len(output) > 0 {
			// 只显示重要的输出信息
			outputStr := string(output)
			if strings.Contains(outputStr, "Already up to date") ||
			   strings.Contains(outputStr, "Fast-forward") ||
			   strings.Contains(outputStr, "Merge") ||
			   strings.Contains(outputStr, "file changed") {
				Infof("输出: %s", outputStr)
			}
		}
	}

	Successf("成功同步 %s 到 %s", productionBranch, devBranch)
	return true
}

// isWorkingDirectoryClean 检查工作目录是否干净
func isWorkingDirectoryClean() bool {
	output, err := exec.Command("git", "status", "--porcelain").CombinedOutput()
	if err != nil {
		return false
	}
	return len(strings.TrimSpace(string(output))) == 0
}
