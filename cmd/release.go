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
		fmt.Println("release called with: ", versionType, version, newVersion)
	},
}

func init() {
	rootCmd.AddCommand(releaseCmd)
	// Here you will define your flags and configuration settings.
	// add Type (MINOR, MAJOR, PATCH) enum
	releaseCmd.Flags().StringP("type", "t", "PATCH", "版本类型")
}
