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
}

var nickname string

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "初始化 Git Flow 配置",
	Run: func(cmd *cobra.Command, args []string) {
		config := YamlConfig{
			BaseBranch: "origin/develop",
			Nickname:   nickname,
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

	// 添加 --nickname 标志
	initCmd.Flags().StringVarP(&nickname, "nickname", "n", "default_nickname", "设置昵称")
}
