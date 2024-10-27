package cmd

import (
	"fmt"
	"github-flow/utils"
	"github.com/spf13/cobra"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "定期同步 main 到 develop 分支",
	Run: func(cmd *cobra.Command, args []string) {
		config := readConfig()
		if config == nil {
			return
		}
		prURL := fmt.Sprintf("https://github.com/%s/compare/%s...%s?expand=1", config.Repository, config.ProductionBranch, config.DevBaseBranch)
		utils.CreatePr(prURL)
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
