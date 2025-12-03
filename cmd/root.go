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

	// Update init command
	if initCmd != nil {
		initCmd.Short = strings.GetString("init", "short")
		initCmd.Flags().Lookup("force").Usage = strings.GetString("init", "force_flag")
		initCmd.Flags().Lookup("nickname").Usage = strings.GetString("init", "nickname_flag")
	}

	// Update publish command
	if publishCmd != nil {
		publishCmd.Short = strings.GetString("publish", "short")
	}

	// Update hotfix command
	if hotfixCmd != nil {
		hotfixCmd.Short = strings.GetString("hotfix", "short")
	}

	// Update checkout command
	if checkoutCmd != nil {
		checkoutCmd.Short = strings.GetString("checkout", "short")
	}

	// Update sync command
	if syncCmd != nil {
		syncCmd.Short = strings.GetString("sync", "short")
	}

	// Update tag command
	if tagCmd != nil {
		tagCmd.Short = strings.GetString("tag", "short")
		tagCmd.Flags().Lookup("type").Usage = strings.GetString("tag", "type_flag")
	}

	// Update pr command
	if prCmd != nil {
		prCmd.Short = strings.GetString("pr", "short")
		prCmd.Flags().Lookup("sync").Usage = strings.GetString("pr", "sync_flag")
		prCmd.Flags().Lookup("open").Usage = strings.GetString("pr", "open_flag")
	}

	// Update sweep command
	if sweepCmd != nil {
		sweepCmd.Short = strings.GetString("sweep", "short")
		sweepCmd.Flags().Lookup("local").Usage = strings.GetString("sweep", "local_flag")
		sweepCmd.Flags().Lookup("remote").Usage = strings.GetString("sweep", "remote_flag")
	}

	// Update release command
	if releaseCmd != nil {
		releaseCmd.Short = strings.GetString("release", "short")
		releaseCmd.Flags().Lookup("type").Usage = strings.GetString("release", "type_flag")
		releaseCmd.Flags().Lookup("hotfix").Usage = strings.GetString("release", "hotfix_flag")
	}

	// Update config command
	if configCmd != nil {
		configCmd.Short = strings.GetString("config", "short")
		configCmd.Long = strings.GetString("config", "long")
	}

	if rebaseCmd != nil {
		rebaseCmd.Short = strings.GetString("rebase", "short")
	}

	// Update rename command
	if renameCmd != nil {
		renameCmd.Short = strings.GetString("rename", "short")
		renameCmd.Flags().Lookup("local").Usage = strings.GetString("rename", "local_flag")
		renameCmd.Flags().Lookup("remote").Usage = strings.GetString("rename", "remote_flag")
		renameCmd.Flags().Lookup("delete").Usage = strings.GetString("rename", "delete_flag")
	}

	// Update restore command
	if restoreCmd != nil {
		restoreCmd.Short = strings.GetString("restore", "short")
		restoreCmd.Long = strings.GetString("restore", "long")
	}
}
