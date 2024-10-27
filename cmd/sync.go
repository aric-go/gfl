package cmd

import (
	"github-flow/utils"
	"github.com/spf13/cobra"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "定期同步 main 到 develop 分支",
	Run: func(cmd *cobra.Command, args []string) {
		config := utils.ReadConfig()
		if config == nil {
			return
		}
		utils.CreatePr(config.DevBaseBranch, config.ProductionBranch)
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
