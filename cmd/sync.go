package cmd

import (
	"gfl/utils"
	"gfl/utils/strings"

	"github.com/spf13/cobra"
)

// syncCmd represents the fetch command
var syncCmd = &cobra.Command{
	Use:     "sync",
	Aliases: []string{"up"},
	Short:   "Sync remote repository to local repository/update all remote repository references", // Will be updated after strings load
	Run: func(cmd *cobra.Command, args []string) {
		if err := utils.RunCommandWithSpin("git fetch origin", strings.GetPath("sync.fetching")); err == nil {
			utils.Success(strings.GetPath("sync.fetch_success"))
		}

		if err := utils.RunCommandWithSpin("git remote update origin --prune", strings.GetPath("sync.updating")); err == nil {
			utils.Success(strings.GetPath("sync.sync_success"))
		}
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
