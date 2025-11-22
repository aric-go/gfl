package utils

import "fmt"

// GenerateBranchName 生成分支名称
// 如果有 nickname，格式为: type/nickname/name
// 如果没有 nickname，格式为: type/name
func GenerateBranchName(branchType, nickname, name string) string {
	if nickname != "" {
		return fmt.Sprintf("%s/%s/%s", branchType, nickname, name)
	}
	return fmt.Sprintf("%s/%s", branchType, name)
}