package cmd

import (
	"fmt"
	"github-flow/utils"
	"github.com/spf13/cobra"
)

// releaseCmd represents the release command
var releaseCmd = &cobra.Command{
	Use:     "release",
	Aliases: []string{"r"},
	Short:   "以最近 tag(eg:v1.0.0) 为基准，生成新的 release 版本",
	Run: func(cmd *cobra.Command, args []string) {
		version := utils.GetLatestVersion()
		versionType, _ := cmd.Flags().GetString("type")
		newVersion, err := utils.IncrementVersion(version, versionType)
		if err != nil {
			fmt.Println(err)
		}

		// print new version
		fmt.Printf("🌈 最新版本: %s\n", version)

		config := readConfig()
		if config == nil {
			return
		}
		branchName := fmt.Sprintf("%s/release-%s", "releases", newVersion)
		baseRemoteBranch := fmt.Sprintf("origin/%s", config.DevBaseBranch)
		// 0. fetch remote branch
		command0 := fmt.Sprintf("git fetch origin")
		if err := utils.RunCommandWithSpin(command0, "0. 正在同步远程分支...\n"); err != nil {
			return
		}

		// 1. create release branch
		command1 := fmt.Sprintf("git checkout -b %s %s", branchName, baseRemoteBranch)
		if err := utils.RunCommandWithSpin(command1, "1.正在创建 Release...\n"); err != nil {
			return
		}
		// 2. push release branch
		command2 := fmt.Sprintf("git push -u origin %s", branchName)
		if err := utils.RunCommandWithSpin(command2, "2.正在推送 Release...\n"); err != nil {
			fmt.Println("step 2 failed: ", err)
			return
		}
		// 3. create release tag
		command3 := fmt.Sprintf("git tag -a %s -m 'Release %s'", newVersion, newVersion)
		if err := utils.RunCommandWithSpin(command3, "3.正在创建 Release Tag...\n"); err != nil {
			return
		}
		// 4. push release tag
		command4 := fmt.Sprintf("git push origin %s", newVersion)
		if err := utils.RunCommandWithSpin(command4, "4.正在推送 Release Tag...\n"); err != nil {
			return
		}
		fmt.Printf("Release %s 创建成功！\n", newVersion)
	},
}

func init() {
	rootCmd.AddCommand(releaseCmd)
	// Here you will define your flags and configuration settings.
	// add Type (MAJOR, MINOR, PATCH) enum
	releaseCmd.Flags().StringP("type", "t", "PATCH", "版本类型: major, minor, patch")
}
