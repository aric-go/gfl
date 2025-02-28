package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "github-flow",
	Short: "Github Flow CLI",
	Run: func(cmd *cobra.Command, args []string) {

		isVersion, _ := cmd.Flags().GetBool("version")
		if isVersion {
			cmd.Println("🌈 Github Flow version: 1.0.6")
			return
		}
		cmd.Print("🌈 Welcome to GitHub Flow CLI!\n\n")
		_ = cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// add --version flag
	rootCmd.Flags().BoolP("version", "v", false, "show version")
	rootCmd.PersistentFlags().BoolP("skip-fetch", "s", false, "跳过远程分支同步")
	rootCmd.PersistentFlags().BoolP("confirm", "y", false, "确认操作")
}
