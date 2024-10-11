package cmd

import (
	"fmt"
	"github-flow/utils"
	"github.com/spf13/cobra"
)

// coCmd represents the co command
var coCmd = &cobra.Command{
	Use:   "co",
	Short: "交互式的git分支切换",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("co called", utils.GetLocalBranches())
	},
}

func init() {
	rootCmd.AddCommand(coCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// coCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// coCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
