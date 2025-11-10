package cmd

import (
	"fmt"
	"gfl/utils"
	"log"

	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

// reviewCmd represents the review command
var reviewCmd = &cobra.Command{
	Use:     "review",
	Aliases: []string{"rv"},
	Short:   "创建代码审查请求（PR/MR）",
	Run: func(cmd *cobra.Command, args []string) {
		// 读取配置文件获取默认分支和仓库
		config := utils.ReadConfig()
		if config == nil {
			return
		}

		// 判断是否为 GitLab 仓库
		isGitlab := config.GitlabHost != ""
		repo, _ := utils.GetRepository()

		// 处理同步标志
		isSync, _ := cmd.Flags().GetBool("sync")
		if isSync {
			if isGitlab {
				utils.CreateMr(config.DevBaseBranch, config.ProductionBranch)
			} else {
				utils.CreatePr(config.DevBaseBranch, config.ProductionBranch)
			}
			return
		}

		// 处理打开列表页面标志
		isOpen, _ := cmd.Flags().GetBool("open")
		if isOpen {
			var listUrl string
			if isGitlab {
				listUrl = fmt.Sprintf("%s/%s/-/merge_requests", config.GitlabHost, repo)
			} else {
				listUrl = fmt.Sprintf("https://github.com/%s/pulls", repo)
			}

			err := browser.OpenURL(listUrl)
			if err != nil {
				log.Fatal(err)
			}
			return
		}

		// 获取当前分支名称
		currentBranch, err := utils.GetCurrentBranch()
		if err != nil {
			fmt.Println("无法获取当前分支:", err)
			return
		}

		// 确定目标分支
		var baseBranch = config.DevBaseBranch
		if args != nil && len(args) > 0 {
			baseBranch = args[0]
		}

		// 根据仓库类型创建相应的审查请求
		if isGitlab {
			utils.CreateMr(baseBranch, currentBranch)
		} else {
			utils.CreatePr(baseBranch, currentBranch)
		}
	},
}

func init() {
	rootCmd.AddCommand(reviewCmd)

	// 添加命令标志
	// reviewCmd.Flags().BoolP("sync", "s", false, "同步 production 分支到 develop 分支")
	reviewCmd.Flags().BoolP("open", "o", false, "打开代码审查列表页面")
}
