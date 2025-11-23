package utils

import (
	"fmt"
	"os/exec"
	"strings"
	"github.com/pkg/browser"
)

func CreatePr(base string, head string) {
	Warning("The 'pr' command is deprecated. Please use 'gfl review' for a better experience.")
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
func SyncProductionToDev(productionBranch, devBranch string) {
	Infof("正在同步生产分支 %s 到开发分支 %s...", productionBranch, devBranch)

	// 执行 git 命令同步分支
	commands := [][]string{
		{"git", "fetch", "origin"},
		{"git", "checkout", devBranch},
		{"git", "pull", "origin", devBranch},
		{"git", "merge", "origin/" + productionBranch},
		{"git", "push", "origin", devBranch},
	}

	for _, cmd := range commands {
		Infof("执行: %s", strings.Join(cmd, " "))
		output, err := exec.Command(cmd[0], cmd[1:]...).CombinedOutput()
		if err != nil {
			Errorf("命令执行失败: %v", err)
			Errorf("错误输出: %s", string(output))
			return
		}
		if len(output) > 0 {
			Infof("输出: %s", string(output))
		}
	}

	Successf("成功同步 %s 到 %s", productionBranch, devBranch)
}
