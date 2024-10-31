package cmd

import (
	"fmt"
	"github-flow/utils"

	"github.com/spf13/cobra"
)

// 这里有个小问题 release + tag 这个过程，应该是先创建 release 分支，然后再创建 tag，最后再切换回原分支。？？
// releaseCmd represents the release command
var releaseCmd = &cobra.Command{
	Use:     "release",
	Aliases: []string{"rls"},
	Short:   "以最近 tag(eg:v1.0.0) 为基准，生成新的 release 版本",
	Run: func(cmd *cobra.Command, args []string) {
		version := utils.GetLatestVersion()
		versionType, _ := cmd.Flags().GetString("type")
		newVersion, err := utils.IncrementVersion(version, versionType)
		if err != nil {
			fmt.Println(err)
		}

		// print new version
		fmt.Printf("🌈 上一版本: %s\n", version)
		fmt.Printf("🎉 新的版本: %s\n", newVersion)

		config := utils.ReadConfig()
		if config == nil {
			return
		}
		branchName := fmt.Sprintf("%s/release-%s", "releases", newVersion)
		baseRemoteBranch := fmt.Sprintf("origin/%s", config.DevBaseBranch)
		// 1. fetch remote branch
		command1 := fmt.Sprintf("git fetch origin")
		if err := utils.RunCommandWithSpin(command1, "1. 正在同步远程分支...\n"); err != nil {
			fmt.Println("step 1 failed: ", err)
			return
		}

		// 2. create release branch
		command2 := fmt.Sprintf("git checkout -b %s %s", branchName, baseRemoteBranch)
		if err := utils.RunCommandWithSpin(command2, "2.正在创建 Release...\n"); err != nil {
			return
		}
		// 3. push release branch
		command3 := fmt.Sprintf("git push -u origin %s", branchName)
		if err := utils.RunCommandWithSpin(command3, "3.正在推送 Release...\n"); err != nil {
			fmt.Println("step 2 failed: ", err)
			return
		}

		// 6. switch back to original branch
		command6 := fmt.Sprintf("git checkout -")
		if err := utils.RunCommandWithSpin(command6, "6.正在切换回原分支...\n"); err != nil {
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(releaseCmd)
	// Here you will define your flags and configuration settings.
	// add Type (MAJOR, MINOR, PATCH) enum
	releaseCmd.Flags().StringP("type", "t", "patch", "版本类型: major, minor, patch")
	// add version flag manual set verison
	releaseCmd.Flags().StringP("version", "v", "", "手动指定版本号")
}
