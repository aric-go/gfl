package cmd

import (
	"fmt"
	"gfl/utils"
	"gfl/utils/strings"

	"github.com/spf13/cobra"
)

var forwardTitle string
var forwardBody string

var forwardCmd = &cobra.Command{
	Use:     "forward",
	Aliases: []string{"fwd"},
	Short:   "Forward main branch to dev branch via PR (alias: fwd)", // Will be updated after strings load
	Run: func(cmd *cobra.Command, args []string) {
		config := utils.ReadConfig()

		if config == nil {
			return
		}

		// Check if production branch and dev branch are the same
		if config.ProductionBranch == config.DevBaseBranch {
			utils.Errorf(strings.GetPath("forward.same_branch_error"))
			return
		}

		// Get current branch
		currentBranch, err := utils.GetCurrentBranch()
		if err != nil {
			utils.Errorf(strings.GetPath("forward.get_branch_error"), err)
			return
		}

		// Check if current branch is production branch (main)
		if currentBranch != config.ProductionBranch {
			utils.Errorf(strings.GetPath("forward.must_be_on_main"), config.ProductionBranch)
			return
		}

		// Set default PR title if not provided
		prTitle := forwardTitle
		if prTitle == "" {
			prTitle = fmt.Sprintf("Sync %s to %s", config.ProductionBranch, config.DevBaseBranch)
		}

		// Set default PR body if not provided
		prBody := forwardBody
		if prBody == "" {
			prBody = fmt.Sprintf("Forwarding changes from `%s` to `%s`.", config.ProductionBranch, config.DevBaseBranch)
		}

		// Create PR using gh CLI
		prCmd := fmt.Sprintf("gh pr create --base %s --head %s --title \"%s\" --body \"%s\"",
			config.DevBaseBranch,
			config.ProductionBranch,
			prTitle,
			prBody)

		if err := utils.RunCommandWithSpin(prCmd, strings.GetPath("forward.creating_pr")); err != nil {
			utils.Errorf(strings.GetPath("forward.create_pr_error"), err)
			return
		}

		utils.Successf(strings.GetPath("forward.success"), config.ProductionBranch, config.DevBaseBranch)
	},
}

func init() {
	forwardCmd.Flags().StringVarP(&forwardTitle, "title", "t", "", "PR title")
	forwardCmd.Flags().StringVarP(&forwardBody, "body", "b", "", "PR description")
	rootCmd.AddCommand(forwardCmd)
}
