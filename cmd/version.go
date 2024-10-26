/**
 * @Author: aric 1290657123@qq.com
 * @Date: 2024-10-10 23:19:41
 * @LastEditors: aric 1290657123@qq.com
 * @LastEditTime: 2024-10-11 21:49:35
 * @FilePath: cmd/version.go
 */
package cmd

import (
	"github-flow/utils"
	"github.com/spf13/cobra"
)

var (
	Version   string
	BuildTime string
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "获取程序版本",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("🌈 Github Flow Version:", utils.GetLatestVersion())
	},
	DisableAutoGenTag: true,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
