package cmd

import (
	"fmt"
	"gfl/utils"
	"gfl/utils/strings"

	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

// prCmd represents the pr command
var prCmd = &cobra.Command{
	Use:     "pr",
	Aliases: []string{"rv"},
	Short:   "Create pull request (PR)",
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
			if !utils.SyncProductionToDev(config.ProductionBranch, config.DevBaseBranch) {
				utils.Errorf(strings.GetString("pr", "sync_failed"))
			}
			return
		}

		// 处理打开列表页面标志
		isOpen, _ := cmd.Flags().GetBool("open")
		if isOpen {
			listUrl := fmt.Sprintf("https://github.com/%s/pulls", repo)

			err := browser.OpenURL(listUrl)
			if err != nil {
				utils.Errorf(strings.GetString("pr", "browser_error"), err)
				return
			}
			return
		}

		// 获取当前分支名称
		currentBranch, err := utils.GetCurrentBranch()
		if err != nil {
			utils.Errorf(strings.GetString("pr", "current_branch_error"), err)
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
	prCmd.Flags().BoolP("sync", "s", false, strings.GetString("pr", "sync_flag"))
	prCmd.Flags().BoolP("open", "o", false, strings.GetString("pr", "open_flag"))
}
