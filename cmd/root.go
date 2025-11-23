package cmd

import (
	"gfl/utils"
	"gfl/utils/strings"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "gfl",
	Short:   "GitHub Flow CLI", // Will be updated after strings load
	Version: "1.0.7",
	Run: func(cmd *cobra.Command, args []string) {
		utils.Info(strings.GetString("root", "welcome"))
		_ = cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// Initialize the strings package
	if err := strings.LoadStrings(); err != nil {
		utils.Errorf("Failed to initialize strings: %v", err)
		os.Exit(1)
	}

	// Update command descriptions after strings are loaded
	updateCommandDescriptions()

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Cobra will automatically add --version/-v flag when Version field is set
	rootCmd.PersistentFlags().BoolP("confirm", "y", false, "Confirm operation") // Will be updated after strings load
}

// updateCommandDescriptions updates all command descriptions after strings are loaded
func updateCommandDescriptions() {
	// Update root command
	rootCmd.Short = strings.GetString("root", "short")

	// Update flag description
	rootCmd.PersistentFlags().Lookup("confirm").Usage = strings.GetString("root", "confirm_flag")

	// Update start command
	if startCmd != nil {
		startCmd.Short = strings.GetString("start", "short")
	}
}
