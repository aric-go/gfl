package utils

import (
	"fmt"
	"github.com/pkg/browser"
)

func CreatePr(base string, head string) {
	fmt.Println("⚠️  Warning: The 'pr' command is deprecated. Please use 'gfl review' for a better experience.")
	repo, _ := GetRepository()

	// 生成 GitHub PR URL
	// @example: https://github.com/applyai-dev/applyai-frontend/compare/dev...feature/aric/gogogo?expand=1
	// https://github.com/applyai-dev/applyai-frontend/compare/${baseBranch}...${headBranch}?expand=1
	url := fmt.Sprintf("https://github.com/%s/compare/%s...%s?expand=1", repo, base, head)

	err := browser.OpenURL(url)
	if err != nil {
		fmt.Println("无法打开浏览器:", err)
	} else {
		fmt.Printf("已打开 PR 页面: %s\n", url)
	}
}
