package utils

import (
	"fmt"
)

func CreateMr(base string, head string) {
	repo, _ := GetRepository()
	config := ReadConfig()

	baseURL := fmt.Sprintf("https://%s/%s", config.GitlabHost, repo)

	// 生成 GitHub MR URL
	// @example: https://gitlab.com/myteam/awesome-project/-/merge_requests/new?merge_request%5Bsource_branch%5D=feature/login-page&merge_request%5Btarget_branch%5D=main
	url := fmt.Sprintf("%s/-/merge_requests/new?merge_request%5Bsource_branch%5D=%s&merge_request%5Btarget_branch%5D=%s", baseURL, base, head)

	err := OpenBrowser(url)
	if err != nil {
		fmt.Println("无法打开浏览器:", err)
	} else {
		fmt.Printf("已打开 PR 页面: %s\n", url)
	}
}
