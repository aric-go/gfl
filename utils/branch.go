package utils

import "fmt"

// GetBranchTypePrefix 获取分支类型的前缀，如果配置中没有设置则使用默认值
func GetBranchTypePrefix(config *YamlConfig, branchType string) string {
	switch branchType {
	case "feature":
		if config.FeaturePrefix != "" {
			return config.FeaturePrefix
		}
		return "feature"
	case "fix":
		if config.FixPrefix != "" {
			return config.FixPrefix
		}
		return "fix"
	case "hotfix":
		if config.HotfixPrefix != "" {
			return config.HotfixPrefix
		}
		return "hotfix"
	default:
		return branchType
	}
}

// GenerateBranchName 生成分支名称
// 如果有 nickname，格式为: type/nickname/name
// 如果没有 nickname，格式为: type/name
// branchType 可以是 "feature", "fix", "hotfix" 或其他自定义类型
func GenerateBranchName(config *YamlConfig, branchType, name string) string {
	prefix := GetBranchTypePrefix(config, branchType)
	if config.Nickname != "" {
		return fmt.Sprintf("%s/%s/%s", prefix, config.Nickname, name)
	}
	return fmt.Sprintf("%s/%s", prefix, name)
}