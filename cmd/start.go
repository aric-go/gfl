package cmd

import (
	"fmt"
	"gfl/utils"
	"gfl/utils/strings"
	str "strings"

	"github.com/spf13/cobra"
)

var startBaseBranch string // 存储 --base 参数值

var startCmd = &cobra.Command{
	Use:     "start [feature-name]",
	Short:   "Start a new feature (alias: s)", // Will be updated after strings load
	Aliases: []string{"s"},
	Args:    cobra.ExactArgs(1), // 要求提供一个参数
	Run: func(cmd *cobra.Command, args []string) {
		config := utils.ReadConfig()

		if config == nil {
			return
		}

		startName := parseStartName(args[0])
		branchName := utils.GenerateBranchName(config, startName.ActionName, startName.FeatureName)

		// 解析基础分支
		baseBranch := determineBaseBranch()
		if baseBranch == "" {
			baseBranch = config.DevBaseBranch
		}

		// 执行命令: git fetch origin（始终执行，确保远程信息最新）
		fetchCmd := "git fetch origin"
		if err := utils.RunCommandWithSpin(fetchCmd, strings.GetPath("start.syncing")); err != nil {
			return
		}

		// 验证指定的 base 分支在远程是否存在
		baseExists, err := utils.RemoteBranchExists(baseBranch)
		if err != nil {
			utils.Errorf(strings.GetPath("start.check_remote_failed"), err)
			return
		}

		if !baseExists {
			// 根据 base 来源选择不同的错误消息
			if startBaseBranch == "@" {
				// 当前分支在远程不存在
				currentBranch, _ := utils.GetCurrentBranch()
				utils.Errorf(strings.GetPath("start.current_branch_not_exist"), currentBranch)
			} else {
				// 指定的 base 分支在远程不存在
				utils.Errorf(strings.GetPath("start.base_not_exist"), baseBranch)
			}
			return
		}

		// 执行命令: git checkout -b feature/aric/new-feature origin/develop
		baseRemoteBranch := fmt.Sprintf("origin/%s", baseBranch)
		checkoutCmd := fmt.Sprintf("git checkout -b %s %s", branchName, baseRemoteBranch)
		if err := utils.RunCommandWithSpin(checkoutCmd, strings.GetPath("start.creating")); err != nil {
			return
		}
		utils.Successf(strings.GetPath("start.success", startName.ActionName, branchName))
	},
}

// determineBaseBranch 确定使用的基础分支
// 优先级：--base 参数 > config.DevBaseBranch
// 如果 --base=@，则返回当前分支名
func determineBaseBranch() string {
	// 如果用户指定了 --base 参数
	if startBaseBranch != "" {
		// 特殊值 @ 表示使用当前分支
		if startBaseBranch == "@" {
			currentBranch, err := utils.GetCurrentBranch()
			if err != nil {
				// 如果获取失败，返回空字符串（由上层使用 config.DevBaseBranch）
				utils.Errorf("Failed to get current branch: %v", err)
				return ""
			}
			return currentBranch
		}
		// 直接使用用户指定的分支名
		return startBaseBranch
	}

	// 没有指定 --base 参数，返回空字符串（由上层使用 config.DevBaseBranch）
	return ""
}

func init() {
	rootCmd.AddCommand(startCmd)
	// 添加 --base flag
	startCmd.Flags().StringVarP(&startBaseBranch, "base", "b", "", strings.GetPath("start.base_flag"))
}

type StartName struct {
	ActionName  string
	FeatureName string
}

func parseStartName(name string) *StartName {
	hasColon := str.Contains(name, ":")
	if hasColon {
		parts := str.Split(name, ":")
		return &StartName{
			ActionName:  parts[0],
			FeatureName: parts[1],
		}
	} else {
		return &StartName{
			ActionName:  "feature",
			FeatureName: name,
		}
	}
}
