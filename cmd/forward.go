package cmd

import (
	"fmt"
	"gfl/utils"
	str "gfl/utils/strings"
	"os/exec"
	"strings"

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

		// Check if PR already exists
		checkCmd := fmt.Sprintf("gh pr list --head %s --base %s --state open --json url,title,number --jq '. | length'",
			config.ProductionBranch,
			config.DevBaseBranch)
		output, _ := exec.Command("bash", "-c", checkCmd).Output()
		outputStr := strings.TrimSpace(string(output))
		if outputStr != "" && outputStr != "0" {
			// PR already exists, get its details
			listCmd := fmt.Sprintf("gh pr list --head %s --base %s --state open --json number,title,url --jq '.[0]\"",
				config.ProductionBranch,
			config.DevBaseBranch)
			prDetailsBytes, _ := exec.Command("bash", "-c", listCmd).Output()
			prDetails := string(prDetailsBytes)
			utils.Errorf(str.GetPath("forward.pr_already_exists"), prDetails)
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

		// Create PR using gh CLI with remote branches
		// GitHub PR automatically uses remote branches for both base and head
		prCmd := fmt.Sprintf("gh pr create --base %s --head %s --title \"%s\" --body \"%s\"",
			config.DevBaseBranch,
			config.ProductionBranch,
			prTitle,
			prBody)

		if err := utils.RunCommandWithSpin(prCmd, str.GetPath("forward.creating_pr")); err != nil {
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
