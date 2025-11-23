package cmd

import (
	"gfl/utils"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gfl",
	Short: "GitHub Flow CLI",
	Run: func(cmd *cobra.Command, args []string) {

		isVersion, _ := cmd.Flags().GetBool("version")
		if isVersion {
			utils.Info("🌈 GitHub Flow version: 1.0.6")
			return
		}
		utils.Info("🌈 Welcome to GitHub Flow CLI!")
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
	rootCmd.PersistentFlags().BoolP("confirm", "y", false, "确认操作")
}
