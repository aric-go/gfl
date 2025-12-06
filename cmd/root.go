package cmd

import (
	"fmt"
	"gfl/utils"
	"gfl/utils/strings"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "gfl",
	Short:   "GitHub Flow CLI", // Will be updated after strings load
	Version: "1.0.8",
	Run: func(cmd *cobra.Command, args []string) {
		utils.DisplayLogo()
		fmt.Println() // Keep for spacing
		fmt.Printf("%s", strings.GetPath("root.welcome"))
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
	rootCmd.Short = strings.GetPath("root.short")

	// Update flag description
	rootCmd.PersistentFlags().Lookup("confirm").Usage = strings.GetPath("root.confirm_flag")

	// Update start command
	if startCmd != nil {
		startCmd.Short = strings.GetPath("start.short")
	}

	// Update init command
	if initCmd != nil {
		initCmd.Short = strings.GetPath("init.short")
		initCmd.Flags().Lookup("force").Usage = strings.GetPath("init.force_flag")
		initCmd.Flags().Lookup("nickname").Usage = strings.GetPath("init.nickname_flag")
	}

	// Update publish command
	if publishCmd != nil {
		publishCmd.Short = strings.GetPath("publish.short")
	}

	// Update hotfix command
	if hotfixCmd != nil {
		hotfixCmd.Short = strings.GetPath("hotfix.short")
	}

	// Update checkout command
	if checkoutCmd != nil {
		checkoutCmd.Short = strings.GetPath("checkout.short")
	}

	// Update sync command
	if syncCmd != nil {
		syncCmd.Short = strings.GetPath("sync.short")
	}

	// Update tag command
	if tagCmd != nil {
		tagCmd.Short = strings.GetPath("tag.short")
		tagCmd.Flags().Lookup("type").Usage = strings.GetPath("tag.type_flag")
	}

	// Update pr command
	if prCmd != nil {
		prCmd.Short = strings.GetPath("pr.short")
		prCmd.Flags().Lookup("sync").Usage = strings.GetPath("pr.sync_flag")
		prCmd.Flags().Lookup("open").Usage = strings.GetPath("pr.open_flag")
	}

	// Update sweep command
	if sweepCmd != nil {
		sweepCmd.Short = strings.GetPath("sweep.short")
		sweepCmd.Flags().Lookup("local").Usage = strings.GetPath("sweep.local_flag")
		sweepCmd.Flags().Lookup("remote").Usage = strings.GetPath("sweep.remote_flag")
	}

	// Update release command
	if releaseCmd != nil {
		releaseCmd.Short = strings.GetPath("release.short")
		releaseCmd.Flags().Lookup("type").Usage = strings.GetPath("release.type_flag")
		releaseCmd.Flags().Lookup("hotfix").Usage = strings.GetPath("release.hotfix_flag")
	}

	// Update config command
	if configCmd != nil {
		configCmd.Short = strings.GetPath("config.short")
		configCmd.Long = strings.GetPath("config.long")
	}

	if rebaseCmd != nil {
		rebaseCmd.Short = strings.GetPath("rebase.short")
	}

	// Update rename command
	if renameCmd != nil {
		renameCmd.Short = strings.GetPath("rename.short")
		renameCmd.Flags().Lookup("local").Usage = strings.GetPath("rename.local_flag")
		renameCmd.Flags().Lookup("remote").Usage = strings.GetPath("rename.remote_flag")
		renameCmd.Flags().Lookup("delete").Usage = strings.GetPath("rename.delete_flag")
	}

	// Update restore command
	if restoreCmd != nil {
		restoreCmd.Short = strings.GetPath("restore.short")
		restoreCmd.Long = strings.GetPath("restore.long")
	}

	// Update bugfix command
	if bugfixCmd != nil {
		bugfixCmd.Short = strings.GetPath("bugfix.short")
	}
}
