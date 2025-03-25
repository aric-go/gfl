/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"gfl/utils"
	"github.com/spf13/cobra"
)

// plCmd represents the pl command
var plCmd = &cobra.Command{
	Use:   "pl",
	Short: "发布清单",
	Run: func(cmd *cobra.Command, args []string) {
		config := utils.ReadConfig()
		utils.IptPublishList(config)
	},
}

func init() {
	rootCmd.AddCommand(plCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// plCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// plCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
