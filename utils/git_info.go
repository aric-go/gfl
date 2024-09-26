/**
 * @Author: aric 1290657123@qq.com
 * @Date: 2024-09-26 15:22:35
 * @LastEditors: aric 1290657123@qq.com
 * @LastEditTime: 2024-09-26 15:28:05
 * @FilePath: utils/git_info.go
 */
package utils

import (
	"fmt"
	"strings"
)

// https://github.com/applyai-dev/applyai-frontend/compare/dev...feature/aric/gogogo?expand=1

func GetGitInfo() string {
	var gitInfo, _ = RunShell("git branch --show-current")
	return gitInfo
}

func GetGitURL() string {
	var gitURL, _ = RunShell("git config --get remote.origin.url")
	return gitURL
}

func ExtractRepoPath(gitURL string) (string, error) {
	// 分割 "git@github.com:applyai-dev/applyai-frontend.git" 的 ":" 后部分
	parts := strings.Split(gitURL, ":")
	if len(parts) != 2 {
		return "", fmt.Errorf("无效的 Git URL")
	}

	// 去掉 .git 后缀
	repoPath := strings.TrimSuffix(parts[1], ".git")
	return repoPath, nil
}
