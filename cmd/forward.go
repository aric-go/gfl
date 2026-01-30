package cmd

import (
	"fmt"
	"gfl/utils"
	str "gfl/utils/strings"
	"time"

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
			utils.Errorf(str.GetPath("forward.same_branch_error"))
			return
		}

		// Sync remote branches first to ensure we have latest info
		if err := utils.RunCommandWithSpin("git fetch origin", str.GetPath("forward.syncing")); err != nil {
			utils.Errorf(str.GetPath("forward.sync_error"), err)
			return
		}

		// Check if remote branches exist
		remoteBranches, err := utils.GetRemoteBranches()
		if err != nil {
			utils.Errorf(str.GetPath("forward.sync_error"), err)
			return
		}

		baseBranch := "origin/" + config.DevBaseBranch
		headBranch := "origin/" + config.ProductionBranch

		baseExists := false
		headExists := false
		for _, branch := range remoteBranches {
			if branch == baseBranch {
				baseExists = true
			}
			if branch == headBranch {
				headExists = true
			}
		}

		if !baseExists {
			utils.Errorf(str.GetPath("forward.base_branch_not_exist"), config.DevBaseBranch)
			return
		}

		if !headExists {
			utils.Errorf(str.GetPath("forward.head_branch_not_exist"), config.ProductionBranch)
			return
		}

		// Set default PR title if not provided
		prTitle := forwardTitle
		if prTitle == "" {
			now := time.Now()
			prTitle = fmt.Sprintf("chore: forward %s to %s (%s)", config.ProductionBranch, config.DevBaseBranch, now.Format("2006-01-02 Mon 15:04"))
		}

		// Set default PR body if not provided
		prBody := forwardBody
		if prBody == "" {
			prBody = fmt.Sprintf("✨ Automated forward from `%s` to `%s`.\n\n_Powered by [gfl](https://github.com/aric-go/gfl) 🚀_", config.ProductionBranch, config.DevBaseBranch)
		}

		// Create PR using gh CLI with remote branches
		// GitHub PR automatically uses remote branches for both base and head
		prArgs := []string{
			"pr", "create",
			"--base", config.DevBaseBranch,
			"--head", config.ProductionBranch,
			"--title", prTitle,
			"--body", prBody,
		}

		if err := utils.RunCommandWithArgs("gh", prArgs, str.GetPath("forward.creating_pr")); err != nil {
			utils.Errorf(str.GetPath("forward.create_pr_error"), err)
			return
		}

		utils.Successf(str.GetPath("forward.success"), config.ProductionBranch, config.DevBaseBranch)
	},
}

func init() {
	forwardCmd.Flags().StringVarP(&forwardTitle, "title", "t", "", "PR title")
	forwardCmd.Flags().StringVarP(&forwardBody, "body", "b", "", "PR description")
	rootCmd.AddCommand(forwardCmd)
}
