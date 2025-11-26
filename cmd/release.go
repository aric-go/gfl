package cmd

import (
	"fmt"
	"gfl/utils"
	"gfl/utils/strings"

	"github.com/spf13/cobra"
)

// 这里有个小问题 release + tag 这个过程，应该是先创建 release 分支，然后再创建 tag，最后再切换回原分支。？？
// releaseCmd represents the release command
var releaseCmd = &cobra.Command{
	Use:     "release",
	Aliases: []string{"rel"},
	Short:   "Generate new release version based on latest tag (eg:v1.0.0)",
	Run: func(cmd *cobra.Command, args []string) {
		version := utils.GetLatestVersion()
		versionType, _ := cmd.Flags().GetString("type")
		hotfix, _ := cmd.Flags().GetBool("hotfix")
		newVersion, err := utils.IncrementVersion(version, versionType)
		if err != nil {
			utils.Error(err.Error())
		}

		// print new version
		utils.Infof(strings.GetString("release", "previous_version"), version)
		utils.Successf(strings.GetString("release", "new_version"), newVersion)

		config := utils.ReadConfig()
		if config == nil {
			return
		}

		remoteBranch := config.DevBaseBranch
		if hotfix {
			remoteBranch = config.ProductionBranch
		}

		branchName := fmt.Sprintf("%s/release-%s", "releases", newVersion)
		baseRemoteBranch := fmt.Sprintf("origin/%s", remoteBranch)
		// 1. fetch remote branch
		command1 := "git fetch origin"
		if err := utils.RunCommandWithSpin(command1, strings.GetString("release", "step1")); err != nil {
			utils.Errorf("step 1 failed: %v", err)
			return
		}

		// 2. create release branch
		command2 := fmt.Sprintf("git checkout -b %s %s", branchName, baseRemoteBranch)
		if err := utils.RunCommandWithSpin(command2, strings.GetString("release", "step2")); err != nil {
			return
		}
		// 3. push release branch
		command3 := fmt.Sprintf("git push -u origin %s", branchName)
		if err := utils.RunCommandWithSpin(command3, strings.GetString("release", "step3")); err != nil {
			utils.Errorf("step 2 failed: %v", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(releaseCmd)
	// Here you will define your flags and configuration settings.
	// add Type (MAJOR, MINOR, PATCH) enum
	releaseCmd.Flags().StringP("type", "t", "patch", strings.GetString("release", "type_flag"))
	// add hotfix flag
	releaseCmd.Flags().BoolP("hotfix", "x", false, strings.GetString("release", "hotfix_flag"))
}
