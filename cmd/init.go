package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"os"
)

type YamlConfig struct {
	BaseBranch string `yaml:"baseBranch"`
	Nickname   string `yaml:"nickname"`
	Repository string `yaml:"repository"`
}

var nickname string
var force bool

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "初始化 Github Flow 配置",
	Run: func(cmd *cobra.Command, args []string) {
		config := YamlConfig{
			BaseBranch: "develop",
			Nickname:   nickname,
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
	initCmd.Flags().BoolVarP(&force, "force", "f", false, "强制覆盖已存在的配置文件")
}
