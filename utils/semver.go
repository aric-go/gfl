package utils

import (
	"bytes"
	"fmt"
	"golang.org/x/mod/semver"
	"os/exec"
	"sort"
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

	if result, err := GetLatestLocalVersion(); err == nil {
		return result
	} else {
		fmt.Println(err)
	}

	return ""
}

func GetLatestLocalVersion() (string, error) {
	// 运行 `git tag` 获取本地所有标签
	cmd := exec.Command("git", "tag")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("执行命令失败: %w", err)
	}

	// 解析标签并过滤出语义化版本
	var versions []string
	lines := strings.Split(out.String(), "\n")
	for _, line := range lines {
		tag := strings.TrimSpace(line)
		if semver.IsValid(tag) {
			versions = append(versions, tag)
		}
	}

	// 检查是否有语义化版本标签
	if len(versions) == 0 {
		return "v1.0.0", nil
	}

	// 对标签进行排序并获取最大版本
	sort.Slice(versions, func(i, j int) bool {
		return semver.Compare(versions[i], versions[j]) < 0
	})

	// 返回最大版本
	return versions[len(versions)-1], nil
}
