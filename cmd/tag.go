package cmd

import (
	"fmt"
	"gfl/utils"
	"gfl/utils/strings"

	"github.com/spf13/cobra"
)

var tagCmd = &cobra.Command{
	Use:     "tag",
	Aliases: []string{"t"},
	Short:   "Generate new tag version for release branch based on latest tag (eg:v1.0.0), or generate new tag version based on previous tag", // Will be updated after strings load
	Run: func(cmd *cobra.Command, args []string) {
		version := utils.GetLatestVersion()
		versionType, _ := cmd.Flags().GetString("type")
		newVersion, err := utils.IncrementVersion(version, versionType)
		if err != nil {
			utils.Error(err.Error())
		}

		// print new version
		utils.Infof(strings.GetString("tag", "previous_version"), version)
		utils.Successf(strings.GetString("tag", "new_version"), newVersion)

		config := utils.ReadConfig()
		if config == nil {
			return
		}
		// 2. checkout to releases/release-x.x.x branch
		command1 := fmt.Sprintf("git checkout releases/release-%s", newVersion)
		if err := utils.RunCommandWithSpin(command1, strings.GetString("tag", "step1")); err != nil {
			return
		}

		// 2. fetch remote branch
		command2 := "git fetch --tags"
		if err := utils.RunCommandWithSpin(command2, strings.GetString("tag", "step2")); err != nil {
			utils.Errorf("step 1 failed: %v", err)
			return
		}

		// 3. create release tag
		command3 := fmt.Sprintf("git tag -a %s -m 'Release-%s'", newVersion, newVersion)
		if err := utils.RunCommandWithSpin(command3, strings.GetString("tag", "step3")); err != nil {
			return
		}
		// 4. push release tag
		command4 := fmt.Sprintf("git push origin %s", newVersion)
		if err := utils.RunCommandWithSpin(command4, strings.GetString("tag", "step4")); err != nil {
			return
		}
		utils.Successf(strings.GetString("tag", "release_success"), newVersion)

		// 5. create release use gh cli
		// ❯ gh release create v1.1.2 --generate-notes
		command5 := fmt.Sprintf("gh release create %s --generate-notes", newVersion)
		if utils.IsCommandAvailable("gh") {
			if err := utils.RunCommandWithSpin(command5, strings.GetString("tag", "step5")); err != nil {
				return
			}
			utils.Successf(strings.GetString("tag", "release_success"), newVersion)
		} else {
			utils.Warning(strings.GetString("tag", "gh_not_installed"))
		}
	},
}

func init() {
	rootCmd.AddCommand(tagCmd)
	// Here you will define your flags and configuration settings.
	// add Type (MAJOR, MINOR, PATCH) enum
	tagCmd.Flags().StringP("type", "t", "patch", "Version type: major, minor, patch") // Will be updated after strings load
}
