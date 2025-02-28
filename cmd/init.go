package cmd

import (
	"embed"
	"fmt"
	"github-flow/utils"
	"os"

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
		local, _ := cmd.Flags().GetBool("local")
		force, _ := cmd.Flags().GetBool("force")
		nickname, _ := cmd.Flags().GetString("nickname")

		gflConfig, _ := assets.ReadFile("assets/.gfl.config.yml")
		gflLocalConfig, _ := assets.ReadFile("assets/.gfl.local.config.yml")

		// convert to YamlConfig
		var gflConfigYaml utils.YamlConfig
		var gflLocalConfigYaml utils.YamlConfig

		_ = yaml.Unmarshal(gflConfig, &gflConfigYaml)
		_ = yaml.Unmarshal(gflLocalConfig, &gflLocalConfigYaml)

		fmt.Println("gflConfigYaml: ", gflConfigYaml)
		fmt.Println("gflLocalConfigYaml: ", gflLocalConfigYaml)
		fmt.Println("初始化 Github Flow 配置", local)

		config := utils.YamlConfig{
			Debug:            false,
			DevBaseBranch:    "develop",
			ProductionBranch: "main",
			Nickname:         nickname,
			GitlabHost:       "https://git.saybot.net",
		}

		if _, err := os.Stat(".gflow.config.yml"); !os.IsNotExist(err) && !force {
			fmt.Println(".gflow.config.yml 文件已存在，如需覆盖请使用 --force 选项")
			return
		}

		// 将配置写入 .gflow.config.yml 文件
		file, err := os.Create(".gflow.config.yml")
		if err != nil {
			fmt.Println("无法创建配置文件:", err)
			return
		}
		defer file.Close()

		data, err := yaml.Marshal(&config)
		if err != nil {
			fmt.Println("无法生成 YAML:", err)
			return
		}

		_, err = file.Write(data)

		utils.AddGitIgnore()

		if err != nil {
			fmt.Println("无法写入配置文件:", err)
		} else {
			fmt.Println(".gflow.config.yml 已生成")
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// 添加 --force 标志
	initCmd.Flags().BoolP("force", "f", false, "强制覆盖已存在的配置文件")
	// 添加 --local 标志
	initCmd.Flags().BoolP("local", "l", false, "初始化本地配置文件")
	// 添加 --nickname 标志
	initCmd.Flags().StringP("nickname", "n", "", "设置 Github Flow 昵称")
}
