package cmd

import (
	"fmt"
	"github-flow/utils"
	"log"

	"github.com/spf13/cobra"
)

// mrCmd represents the mr command
var mrCmd = &cobra.Command{
	Use:   "mr",
	Short: "Gitlab merge request",
	Run: func(cmd *cobra.Command, args []string) {
		config := utils.ReadConfig()
		isSync, _ := cmd.Flags().GetBool("sync")
		isOpen, _ := cmd.Flags().GetBool("open")
		repo, _ := utils.GetRepository()

		if config == nil {
			return
		}

		if isSync {
			utils.CreateMr(config.DevBaseBranch, config.ProductionBranch)
			return
		}

		if isOpen {
			prsUrl := fmt.Sprintf("%s/%s/-/merge_requests", config.GitlabHost, repo)
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

		utils.CreateMr(baseBranch, currentBranch)
	},
}

func init() {
	rootCmd.AddCommand(mrCmd)

	// add sync flag bool
	mrCmd.Flags().BoolP("open", "o", false, "打开当前仓库的 merge requests 页面")
	mrCmd.Flags().BoolP("sync", "s", false, "不定期同步 production 分支 develop 分支")
}
