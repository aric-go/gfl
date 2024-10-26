package utils

import (
	"fmt"
	"golang.org/x/mod/semver"
	"strconv"
	"strings"
)

// IncrementVersion 根据传入的 part (MAJOR, MINOR, PATCH) 更新相应版本号
func IncrementVersion(currentVersion string, versionType string) (string, error) {
	// 检查版本号是否有效
	if !semver.IsValid(currentVersion) {
		return "", fmt.Errorf("无效的版本号: %s", currentVersion)
	}

	// 去掉前缀 "v"
	version := currentVersion[1:]
	parts := strings.Split(version, ".")
	if len(parts) != 3 {
		return "", fmt.Errorf("版本号格式错误: %s", currentVersion)
	}

	// 解析 MAJOR, MINOR, PATCH 为整数
	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return "", fmt.Errorf("转换 MAJOR 部分出错: %v", err)
	}
	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", fmt.Errorf("转换 MINOR 部分出错: %v", err)
	}
	patch, err := strconv.Atoi(parts[2])
	if err != nil {
		return "", fmt.Errorf("转换 PATCH 部分出错: %v", err)
	}

	// 根据传入的 versionType 值递增对应部分
	switch strings.ToUpper(versionType) {
	case "MAJOR":
		major++
		minor = 0
		patch = 0
	case "MINOR":
		minor++
		patch = 0
	case "PATCH":
		patch++
	default:
		return "", fmt.Errorf("无效的版本部分: %s", versionType)
	}

	// 构造新的版本号
	newVersion := fmt.Sprintf("v%d.%d.%d", major, minor, patch)
	return newVersion, nil
}

func GetLatestVersion() string {
	//git fetch --tags
	command := "git fetch --tags"
	_, err := RunShell(command)
	if err != nil {
		fmt.Println(err)
	}

	version, err := RunShell("git describe --tags --abbrev=0")
	if err != nil {
		fmt.Println(err)
	}
	return strings.TrimSpace(version)
}
