package utils

import (
	"fmt"
	"github.com/pkg/browser"
)

func CreateMr(base string, head string) {
	fmt.Println("⚠️  Warning: The 'mr' command is deprecated. Please use 'gfl review' for a better experience.")
	repo, _ := GetRepository()
	config := ReadConfig()

	baseURL := fmt.Sprintf("%s/%s", config.GitlabHost, repo)

	// 生成 GitHub MR URL
	// @example: https://gitlab.com/myteam/awesome-project/-/merge_requests/new?merge_request[source_branch]=feature/login-page&merge_request[target_branch]=main
	url := fmt.Sprintf("%s/-/merge_requests/new?merge_request[source_branch]=%s&merge_request[target_branch]=%s", baseURL, head, base)

	err := browser.OpenURL(url)
	if err != nil {
		fmt.Println("无法打开浏览器:", err)
	} else {
		fmt.Printf("已打开 PR 页面: %s\n", url)
	}
}
