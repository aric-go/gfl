package cmd

import (
	"gfl/utils"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "gfl",
	Short:   "GitHub Flow CLI",
	Version: "1.0.7",
	Run: func(cmd *cobra.Command, args []string) {
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
	// Cobra will automatically add --version/-v flag when Version field is set
	rootCmd.PersistentFlags().BoolP("confirm", "y", false, "确认操作")
}
