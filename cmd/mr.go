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
		if config == nil {
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mrCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mrCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
