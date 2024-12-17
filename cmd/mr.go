/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github-flow/utils"

	"github.com/spf13/cobra"
)

// mrCmd represents the mr command
var mrCmd = &cobra.Command{
	Use:   "mr",
	Short: "Gitlab merge request",
	Run: func(cmd *cobra.Command, args []string) {
		config := utils.ReadConfig()
		isSync, _ := cmd.Flags().GetBool("sync")

		if config == nil {
			return
		}

		if isSync {
			utils.CreatePr(config.DevBaseBranch, config.ProductionBranch)
			return
		}

		// 获取当前的分支名称
		currentBranch, err := getCurrentBranch()
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
	prCmd.Flags().BoolP("sync", "s", false, "不定期同步 production 分支 develop 分支")
}
