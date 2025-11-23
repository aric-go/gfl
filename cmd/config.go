package cmd

import (
	"fmt"
	"gfl/utils"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "查看当前配置",
	Long:  `显示当前 GFL 的配置信息，包括计算后的最终值`,
	Run: func(cmd *cobra.Command, args []string) {
		config := utils.ReadConfig()
		if config == nil {
			utils.Error("无法读取配置文件")
			return
		}

		// 创建表格
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.SetTitle("GFL 配置信息")
		t.Style().Options.SeparateRows = true
		t.Style().Options.DrawBorder = true

		// 添加表头
		t.AppendHeader(table.Row{"配置项", "配置值", "最终值"})

		// 基础配置
		if config.DevBaseBranch != "" {
			t.AppendRow(table.Row{"开发基础分支", config.DevBaseBranch, config.DevBaseBranch})
		} else {
			t.AppendRow(table.Row{"开发基础分支", "<未设置>", "dev"})
		}

		if config.ProductionBranch != "" {
			t.AppendRow(table.Row{"生产分支", config.ProductionBranch, config.ProductionBranch})
		} else {
			t.AppendRow(table.Row{"生产分支", "<未设置>", "main"})
		}

		t.AppendRow(table.Row{"昵称", config.Nickname, config.Nickname})
		t.AppendRow(table.Row{"调试模式", fmt.Sprintf("%v", config.Debug), fmt.Sprintf("%v", config.Debug)})

		// 分支前缀配置 - 显示配置值和最终计算值
		featurePrefix := config.FeaturePrefix
		if featurePrefix == "" {
			t.AppendRow(table.Row{"功能分支前缀", "<未设置>", "feature"})
		} else {
			t.AppendRow(table.Row{"功能分支前缀", featurePrefix, featurePrefix})
		}

		fixPrefix := config.FixPrefix
		if fixPrefix == "" {
			t.AppendRow(table.Row{"修复分支前缀", "<未设置>", "fix"})
		} else {
			t.AppendRow(table.Row{"修复分支前缀", fixPrefix, fixPrefix})
		}

		hotfixPrefix := config.HotfixPrefix
		if hotfixPrefix == "" {
			t.AppendRow(table.Row{"热修复分支前缀", "<未设置>", "hotfix"})
		} else {
			t.AppendRow(table.Row{"热修复分支前缀", hotfixPrefix, hotfixPrefix})
		}

		// 添加示例分支名称
		t.AppendSeparator()
		t.AppendRow(table.Row{"示例功能分支", "", utils.GenerateBranchName(config, "feature", "new-feature")})

		// 渲染表格
		t.Render()

		// 显示配置文件信息
		configFile := os.Getenv("GFL_CONFIG_FILE")
		if configFile == "" {
			configFile = ".gfl.config.yml"
		}

		fmt.Printf("\n配置文件: %s\n", configFile)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}