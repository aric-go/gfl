package cmd

import (
	"fmt"
	"github-flow/utils"
	"github.com/spf13/cobra"
	"log"
)

// prCmd represents the pr command
var prCmd = &cobra.Command{
	Use:   "pr",
	Short: "打开 GitHub 的 PR 页面",
	Run: func(cmd *cobra.Command, args []string) {
		// 读取配置文件获取默认分支和仓库
		config := utils.ReadConfig()
		if config == nil {
			return
		}
		isSync, _ := cmd.Flags().GetBool("sync")
		isOpen, _ := cmd.Flags().GetBool("open")

		repo, _ := utils.GetRepository()

		if isSync {
			utils.CreatePr(config.DevBaseBranch, config.ProductionBranch)
			return
		}

		if isOpen {
			prsUrl := fmt.Sprintf("https://github.com/%s/pulls", repo)
			err := utils.OpenBrowser(prsUrl)
			if err != nil {
				log.Fatal(err)
			}
			return
		}

		// 获取当前的分支名称
		currentBranch, err := utils.GetCurrentBranch()
		if err != nil {
			fmt.Println("无法获取当前分支:", err)
			return
		}

		var baseBranch = config.DevBaseBranch
		if args != nil && len(args) > 0 {
			baseBranch = args[0]
		}

		utils.CreatePr(baseBranch, currentBranch)
	},
}

func init() {
	rootCmd.AddCommand(prCmd)
}

func init() {
	// add sync flag bool
	prCmd.Flags().BoolP("sync", "s", false, "不定期同步 production 分支 develop 分支")
	// open pull requests page
	prCmd.Flags().BoolP("open", "o", false, "打开 GitHub 的 PR 列表页面")
}
