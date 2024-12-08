package cmd

import (
	"fmt"
	"github-flow/utils"

	"github.com/spf13/cobra"
)

var tagCmd = &cobra.Command{
	Use:     "tag",
	Aliases: []string{"t"},
	Short:   "以最近 tag(eg:v1.0.0) 为基准，为 release 分支生成新的 tag 版本，以上一个 tag 为基准，生成新的 tag 版本",
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
		// 2. checkout to releases/release-x.x.x branch
		command1 := fmt.Sprintf("git checkout releases/release-%s", newVersion)
		if err := utils.RunCommandWithSpin(command1, "1. 正在切换到 Release 分支...\n"); err != nil {
			return
		}

		// 2. fetch remote branch
		command2 := "git fetch --tags"
		if err := utils.RunCommandWithSpin(command2, "2. 正在同步远程tag...\n"); err != nil {
			fmt.Println("step 1 failed: ", err)
			return
		}

		// 3. create release tag
		command3 := fmt.Sprintf("git tag -a %s -m 'Release-%s'", newVersion, newVersion)
		if err := utils.RunCommandWithSpin(command3, "3.正在创建 Release Tag...\n"); err != nil {
			return
		}
		// 4. push release tag
		command4 := fmt.Sprintf("git push origin %s", newVersion)
		if err := utils.RunCommandWithSpin(command4, "4.正在推送 Release Tag...\n"); err != nil {
			return
		}
		fmt.Printf("Release %s 创建成功！\n", newVersion)

		// 5. create release use gh cli
		// ❯ gh release create v1.1.2 --generate-notes
		command5 := fmt.Sprintf("gh release create %s --generate-notes", newVersion)
		if utils.IsCommandAvailable("gh") {
			if err := utils.RunCommandWithSpin(command5, "5.正在创建 Release...\n"); err != nil {
				return
			}
			fmt.Printf("Release %s 创建成功！\n", newVersion)
		} else {
			fmt.Println("gh cli 未安装，请手动创建 Release...")
		}
	},
}

func init() {
	rootCmd.AddCommand(tagCmd)
	// Here you will define your flags and configuration settings.
	// add Type (MAJOR, MINOR, PATCH) enum
	tagCmd.Flags().StringP("type", "t", "patch", "版本类型: major, minor, patch")
}
