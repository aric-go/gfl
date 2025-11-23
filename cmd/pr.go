package cmd

import (
	"fmt"
	"gfl/utils"

	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

// prCmd represents the pr command
var prCmd = &cobra.Command{
	Use:     "pr",
	Aliases: []string{"rv"},
	Short:   "创建代码审查请求（PR）",
	Run: func(cmd *cobra.Command, args []string) {
		// 读取配置文件获取默认分支和仓库
		config := utils.ReadConfig()
		if config == nil {
			return
		}

		repo, _ := utils.GetRepository()

		// 处理同步标志
		isSync, _ := cmd.Flags().GetBool("sync")
		if isSync {
			utils.SyncProductionToDev(config.ProductionBranch, config.DevBaseBranch)
			return
		}

		// 处理打开列表页面标志
		isOpen, _ := cmd.Flags().GetBool("open")
		if isOpen {
			listUrl := fmt.Sprintf("https://github.com/%s/pulls", repo)

			err := browser.OpenURL(listUrl)
			if err != nil {
				utils.Errorf("无法打开浏览器: %v", err)
				return
			}
			return
		}

		// 获取当前分支名称
		currentBranch, err := utils.GetCurrentBranch()
		if err != nil {
			utils.Errorf("无法获取当前分支: %v", err)
			return
		}

		// 确定目标分支
		var baseBranch = config.DevBaseBranch
		if len(args) > 0 {
			baseBranch = args[0]
		}

		// 创建 GitHub PR
		utils.CreatePr(baseBranch, currentBranch)
	},
}

func init() {
	rootCmd.AddCommand(prCmd)

	// 添加命令标志
	prCmd.Flags().BoolP("sync", "s", false, "同步 production 分支到 develop 分支")
	prCmd.Flags().BoolP("open", "o", false, "打开代码审查列表页面")
}
