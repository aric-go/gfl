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
	Short: "初始化 Github Flow 配置",
	Run: func(cmd *cobra.Command, args []string) {

		// get flag
		force, _ := cmd.Flags().GetBool("force")
		nickname, _ := cmd.Flags().GetString("nickname")

		gflConfig, _ := assets.ReadFile("assets/.gfl.config.yml")
		gflLocalConfig, _ := assets.ReadFile("assets/.gfl.config.local.yml")

		// convert to YamlConfig
		var gflConfigYaml utils.YamlConfig
		var gflLocalConfigYaml utils.YamlConfig

		_ = yaml.Unmarshal(gflConfig, &gflConfigYaml)
		_ = yaml.Unmarshal(gflLocalConfig, &gflLocalConfigYaml)

		if nickname != "" {
			gflLocalConfigYaml.Nickname = nickname
		}

		// remove empty fields for local config
		utils.RemoveEmptyFields(&gflLocalConfigYaml)

		// create .gfl.config.yml file
		err := utils.CreateGflConfig(gflConfigYaml, utils.CreateGflConfigOptions{
			Filename:     ".gfl.config.yml",
			Force:        false,
			AddGitIgnore: false,
		})

		if err != nil {
			utils.Error(err.Error())
		}

		// create .gfl.config.local.yml file
		err = utils.CreateGflConfig(gflLocalConfigYaml, utils.CreateGflConfigOptions{
			Filename:     ".gfl.config.local.yml",
			Force:        force,
			AddGitIgnore: true,
		})

		if err != nil {
			utils.Error(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// 添加 --force 标志
	initCmd.Flags().BoolP("force", "f", false, "强制覆盖已存在的配置文件")
	// 添加 --nickname 标志
	initCmd.Flags().StringP("nickname", "n", "", "设置 Github Flow 昵称 (可选)")
}
