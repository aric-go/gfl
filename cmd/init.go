package cmd

import (
	"embed"
	"gfl/utils"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

//go:embed assets/.gfl.config.yml
//go:embed assets/.gfl.config.local.yml
var assets embed.FS

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Github Flow configuration", // Will be updated after strings load
	Run: func(cmd *cobra.Command, args []string) {

		// get flag
		force, _ := cmd.Flags().GetBool("force")
		nickname, _ := cmd.Flags().GetString("nickname")

		gflConfig, _ := assets.ReadFile("assets/.gfl.config.yml")
		gflLocalConfig, _ := assets.ReadFile("assets/.gfl.config.local.yml")

		// create .gfl.config.yml file (direct copy to preserve comments)
		err := utils.CreateGflConfigFromBytes(gflConfig, utils.CreateGflConfigOptions{
			Filename:     ".gfl.config.yml",
			Force:        force,
			AddGitIgnore: false,
		})

		if err != nil {
			utils.Error(err.Error())
		}

		// create .gfl.config.local.yml file
		if nickname != "" {
			// Parse and update local config if nickname is provided
			var gflLocalConfigYaml utils.YamlConfig
			_ = yaml.Unmarshal(gflLocalConfig, &gflLocalConfigYaml)
			gflLocalConfigYaml.Nickname = nickname
			utils.RemoveEmptyFields(&gflLocalConfigYaml)

			err = utils.CreateGflConfig(gflLocalConfigYaml, utils.CreateGflConfigOptions{
				Filename:     ".gfl.config.local.yml",
				Force:        force,
				AddGitIgnore: true,
			})
		} else {
			// Direct copy to preserve comments if no nickname
			err = utils.CreateGflConfigFromBytes(gflLocalConfig, utils.CreateGflConfigOptions{
				Filename:     ".gfl.config.local.yml",
				Force:        force,
				AddGitIgnore: true,
			})
		}

		if err != nil {
			utils.Error(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// 添加 --force 标志
	initCmd.Flags().BoolP("force", "f", false, "Force overwrite existing configuration file") // Will be updated after strings load
	// 添加 --nickname 标志
	initCmd.Flags().StringP("nickname", "n", "", "Set Github Flow nickname (optional)") // Will be updated after strings load
}
