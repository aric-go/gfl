package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the github flow config",
	Run: func(cmd *cobra.Command, args []string) {
		configContent := "defaultBranch: origin/develop\n"
		err := ioutil.WriteFile(".gflow.config.yml", []byte(configContent), 0644)
		if err != nil {
			fmt.Println("Error creating config file:", err)
			return
		}
		fmt.Println(".gflow.config.yml created successfully.")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
