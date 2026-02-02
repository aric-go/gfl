package cmd

import (
	"fmt"
	"gfl/utils"
	gflstrings "gfl/utils/strings"
	"os/exec"
	str "strings"

	"github.com/spf13/cobra"
)

var copyConfirm bool // 存储 --confirm 参数值

var copyCmd = &cobra.Command{
	Use:     "copy [new-branch-name]",
	Short:   "Copy current branch to new branch(alias: cp)", // Will be updated after strings load
	Aliases: []string{"cp"},
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		config := utils.ReadConfig()
		if config == nil {
			return
		}

		// Determine the new branch name
		var newBranchName string
		if len(args) == 0 {
			// No branch name provided, generate automatically
			baseBranchName := generateCopyBranchName(config)
			// Generate the full branch name with prefix and nickname
			newBranchName = utils.GenerateBranchName(config, "feature", baseBranchName)

			// Require -y flag if no branch name provided
			if !copyConfirm {
				currentBranch, _ := utils.GetCurrentBranch()
				utils.Warningf(gflstrings.GetPath("copy.error.confirm_required"), currentBranch, newBranchName)
				utils.Info(gflstrings.GetPath("copy.info.use_confirm_flag"))
				return
			}
		} else {
			newBranchName = args[0]
		}

		// Copy branch; skipGenerate is true when we already generated the full name
		skipGenerate := (len(args) == 0)
		copyBranch(config, newBranchName, skipGenerate)
	},
}

func init() {
	rootCmd.AddCommand(copyCmd)
	// Add --confirm flag for automatic name generation
	copyCmd.Flags().BoolVarP(&copyConfirm, "confirm", "y", false, gflstrings.GetPath("copy.confirm_flag"))
}

// generateCopyBranchName generates a branch name by appending "-copyed" to the current branch name.
// It extracts the base name from the current branch (removing prefix and nickname if present).
func generateCopyBranchName(config *utils.YamlConfig) string {
	currentBranch, err := utils.GetCurrentBranch()
	if err != nil {
		// If we can't get current branch, return a generic name
		return "branch-copyed"
	}

	// Extract the last part of the branch name (the actual feature name)
	parts := str.Split(currentBranch, "/")
	baseName := parts[len(parts)-1]

	// Append -copyed suffix
	return baseName + "-copyed"
}

// copyBranch creates a new branch from the current branch.
// It performs validation before creating the branch to ensure a safe copy operation.
func copyBranch(config *utils.YamlConfig, newBranchName string, skipGenerate bool) {
	// Step 1: Generate branch name with proper prefix and case formatting
	var generatedBranchName string
	if skipGenerate {
		// Name is already fully generated, use as-is
		generatedBranchName = newBranchName
	} else {
		// Generate the full branch name with prefix and nickname
		generatedBranchName = utils.GenerateBranchName(config, "feature", newBranchName)
	}

	// Step 2: Check if working directory is clean
	if !isWorkingDirectoryClean() {
		utils.Errorf(gflstrings.GetPath("copy.error.dirty"))
		return
	}

	// Step 3: Get current branch
	currentBranch, err := utils.GetCurrentBranch()
	if err != nil {
		utils.Errorf(gflstrings.GetPath("copy.error.failed"), err)
		return
	}

	// Step 4: Fetch remote to ensure up-to-date information
	fetchCmd := "git fetch origin"
	if err := utils.RunCommandWithSpin(fetchCmd, gflstrings.GetPath("start.syncing")); err != nil {
		return
	}

	// Step 5: Validate current branch exists in remote
	currentBranchExistsInRemote, err := utils.RemoteBranchExists(currentBranch)
	if err != nil {
		utils.Errorf(gflstrings.GetPath("copy.error.failed"), err)
		return
	}

	if !currentBranchExistsInRemote {
		utils.Errorf(gflstrings.GetPath("copy.error.notInRemote"), currentBranch)
		return
	}

	// Step 6: Check if generated branch name already exists locally
	localBranches := utils.GetLocalBranches()
	for _, branch := range localBranches {
		// Remove leading whitespace and asterisk from git branch output
		branchName := str.TrimSpace(str.TrimPrefix(branch, "*"))
		if branchName == generatedBranchName {
			utils.Errorf(gflstrings.GetPath("copy.error.alreadyExists"), generatedBranchName)
			return
		}
	}

	// Step 7: Create new branch from current branch's remote version
	remoteBranchRef := fmt.Sprintf("origin/%s", currentBranch)
	checkoutCmd := fmt.Sprintf("git checkout -b %s %s", generatedBranchName, remoteBranchRef)
	if err := utils.RunCommandWithSpin(checkoutCmd, gflstrings.GetPath("copy.copying")); err != nil {
		return
	}

	// Step 8: Display success message
	utils.Successf(gflstrings.GetPath("copy.success"), currentBranch, generatedBranchName)
}

// isWorkingDirectoryClean checks if the Git working directory has no uncommitted changes.
// This function uses 'git status --porcelain' to determine if there are any pending changes.
//
// Returns:
//   - bool: true if working directory is clean, false otherwise
//
// Method:
//   - Uses --porcelain format for machine-readable output
//   - Empty output indicates no pending changes
//   - Any output (staged or unstaged) indicates a dirty working directory
//
// This check ensures safe branch switching and synchronization operations.
func isWorkingDirectoryClean() bool {
	output, err := exec.Command("git", "status", "--porcelain").CombinedOutput()
	if err != nil {
		// If git command fails, assume working directory is not clean
		return false
	}

	// Trim whitespace and check if output is empty
	return str.TrimSpace(string(output)) == ""
}
