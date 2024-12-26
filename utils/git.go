package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

func extractOwnerAndRepo(url string) (string, error) {
	if strings.HasPrefix(url, "git@") {
		// 处理 SSH 格式
		parts := strings.Split(url, ":")
		if len(parts) != 2 {
			return "", fmt.Errorf("invalid git URL format: %s", url)
		}
		path := strings.TrimSuffix(parts[1], ".git")
		return path, nil
	} else if strings.HasPrefix(url, "https://") {
		// 处理 HTTPS 格式
		parts := strings.Split(url, "/")
		if len(parts) < 2 {
			return "", fmt.Errorf("invalid git URL format: %s", url)
		}
		orgRepo := strings.Join(parts[len(parts)-2:], "/")
		orgRepo = strings.TrimSuffix(orgRepo, ".git")
		return orgRepo, nil
	}
	return "", fmt.Errorf("unsupported git URL format: %s", url)
}

func GetRepository() (string, error) {
	// Get current repository URL
	url, err := exec.Command("git", "config", "--get", "remote.origin.url").Output()
	if err != nil {
		return "", err
	}
	// Extract owner and repository name from URL
	return extractOwnerAndRepo(strings.TrimSpace(string(url)))
}

func GetCurrentBranch() (string, error) {
	// 执行 git 命令获取当前分支
	output, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		return "", err
	}

	// 去除换行符
	return strings.TrimSpace(string(output)), nil
}
