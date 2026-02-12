package cmd

import (
	"strings"

	"gfl/utils"
	"github.com/spf13/cobra"
)

// checkoutCmd represents the co command
var checkoutCmd = &cobra.Command{
	Use:     "checkout [filter]",
	Aliases: []string{"co"},
	Short:   "Interactive git branch switching (alias: co)", // Will be updated after strings load
	Run: func(cmd *cobra.Command, args []string) {
		branches := utils.GetLocalBranches()

		// Filter branches if filter argument is provided
		if len(args) > 0 {
			filter := args[0]
			filteredBranches := []string{}
			for _, branch := range branches {
				if contains(branch, filter) {
					filteredBranches = append(filteredBranches, branch)
				}
			}
			branches = filteredBranches
		}

		utils.BuildCommandList(branches)
	},
}

// contains checks if a string contains a substring (case-insensitive)
func contains(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

func init() {
	rootCmd.AddCommand(checkoutCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// coCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// coCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
