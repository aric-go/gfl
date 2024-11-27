package utils

import (
	"fmt"
	"os/exec"
	"runtime"
)

func CreatePr(base string, head string) {
	repo, _ := GetRepository()

	// 生成 GitHub PR URL
	// @example: https://github.com/applyai-dev/applyai-frontend/compare/dev...feature/aric/gogogo?expand=1
	// https://github.com/applyai-dev/applyai-frontend/compare/${baseBranch}...${headBranch}?expand=1
	url := fmt.Sprintf("https://github.com/%s/compare/%s...%s?expand=1", repo, base, head)

	err := OpenBrowser(url)
	if err != nil {
		fmt.Println("无法打开浏览器:", err)
	} else {
		fmt.Printf("已打开 PR 页面: %s\n", url)
	}
}

// 打开浏览器函数
func OpenBrowser(url string) error {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("无法识别的操作系统")
	}

	return err
}
