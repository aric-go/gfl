package cmd

import (
	"fmt"

	"gfl/utils"
	"gfl/utils/strings"

	"github.com/spf13/cobra"
)

// rebaseCmd represents the rebase command
var rebaseCmd = &cobra.Command{
	Use:     "rebase",
	Aliases: []string{"rb"},
	Short:   "", // Will be updated after strings load
	Long:    "", // Will be updated after strings load
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// Read configuration
		config := utils.ReadConfig()

		// Get current branch
		currentBranch, err := utils.GetCurrentBranch()
		if err != nil {
			utils.Errorf(strings.GetString("rebase", "current_branch_error", err))
			return
		}

		// Use configured dev base branch or default
		devBranch := "dev"
		if config != nil && config.DevBaseBranch != "" {
			devBranch = config.DevBaseBranch
		} else {
			utils.Infof(strings.GetString("rebase", "no_config"))
		}

		// Check if we're already on the target branch
		if currentBranch == devBranch {
			utils.Warningf("Already on branch '%s', no need to rebase", devBranch)
			return
		}

		// Fetch remote branch information with fallback
		fetchCmd := fmt.Sprintf("git fetch origin %s", devBranch)
		if err := utils.RunCommandWithSpin(fetchCmd, strings.GetString("rebase", "fetching")); err != nil {
			// Try to fall back to main branch if configured branch doesn't exist
			if devBranch != "main" {
				utils.Warningf("Failed to fetch '%s' branch, trying 'main' instead: %v", devBranch, err)
				devBranch = "main"
				fetchCmd = fmt.Sprintf("git fetch origin %s", devBranch)
				if err := utils.RunCommandWithSpin(fetchCmd, strings.GetString("rebase", "fetching")); err != nil {
					utils.Errorf("Failed to fetch remote branch 'main': %v", err)
					return
				}
			} else {
				utils.Errorf("Failed to fetch remote branch: %v", err)
				return
			}
		}

		// Check again if we're already on the target branch (after fallback)
		if currentBranch == devBranch {
			utils.Warningf("Already on branch '%s', no need to rebase", devBranch)
			return
		}

		// Perform rebase
		rebaseCmd := fmt.Sprintf("git rebase origin/%s", devBranch)
		if err := utils.RunCommandWithSpin(rebaseCmd, fmt.Sprintf(strings.GetString("rebase", "rebasing"), devBranch)); err != nil {
			utils.Errorf(strings.GetString("rebase", "rebase_failed", err))
			return
		}

		utils.Infof(strings.GetString("rebase", "success", devBranch))
	},
}

func init() {
	rootCmd.AddCommand(rebaseCmd)

	// Update command description after strings are loaded
	rebaseCmd.Short = strings.GetString("rebase", "short")
}
