/**
 * @Author: aric 1290657123@qq.com
 * @Date: 2024-10-10 23:19:41
 * @LastEditors: aric 1290657123@qq.com
 * @LastEditTime: 2024-10-10 23:19:48
 * @FilePath: cmd/version.go
 */
package cmd

import "github.com/spf13/cobra"

var (
	Version   string
	BuildTime string
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of the program",
	Long:  `All software has versions. This is the version number of the program`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Version:", Version)
		cmd.Println("Build Time:", BuildTime)
	},
	DisableAutoGenTag: true,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
