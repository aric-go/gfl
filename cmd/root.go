package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "github-flow",
	Short: "Github Flow CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("版本: %s\n", cmd.Version)
		fmt.Print("🌈 Welcome to GitHub Flow CLI!\n\n")
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
