package cmd

import (
	"embed"
	"github-flow/utils"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

//go:embed assets/.gfl.config.yml
//go:embed assets/.gfl.local.config.yml
var assets embed.FS

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "初始化 Github Flow 配置",
	Run: func(cmd *cobra.Command, args []string) {

		// get flag
		force, _ := cmd.Flags().GetBool("force")
		nickname, _ := cmd.Flags().GetString("nickname")

		gflConfig, _ := assets.ReadFile("assets/.gfl.config.yml")
		gflLocalConfig, _ := assets.ReadFile("assets/.gfl.local.config.yml")

		// convert to YamlConfig
		var gflConfigYaml utils.YamlConfig
		var gflLocalConfigYaml utils.YamlConfig

		_ = yaml.Unmarshal(gflConfig, &gflConfigYaml)
		_ = yaml.Unmarshal(gflLocalConfig, &gflLocalConfigYaml)

		gflLocalConfigYaml.Nickname = nickname

		// remove empty fields for local config
		utils.RemoveEmptyFields(&gflLocalConfigYaml)

		// create .gfl.config.yml file
		_ = utils.CreateGflConfig(gflConfigYaml, utils.CreateGflConfigOptions{
			Filename:     ".gfl.config.yml",
			Force:        force,
			AddGitIgnore: false,
		})

		// create .gfl.config.yml file
		_ = utils.CreateGflConfig(gflLocalConfigYaml, utils.CreateGflConfigOptions{
			Filename:     ".gfl.local.config.yml",
			Force:        force,
			AddGitIgnore: true,
		})
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// 添加 --force 标志
	initCmd.Flags().BoolP("force", "f", false, "强制覆盖已存在的配置文件")
	// 添加 --nickname 标志
	initCmd.Flags().StringP("nickname", "n", "", "设置 Github Flow 昵称")

	// mark nickname as required
	_ = initCmd.MarkFlagRequired("nickname")
}
