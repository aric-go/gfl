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
	Long:  `显示当前 GFL 的配置信息，包括所有配置来源和最终值`,
	Run: func(cmd *cobra.Command, args []string) {
		configInfo := utils.ReadConfigWithSources()
		finalConfig := configInfo.FinalConfig

		// 1. 显示最终配置 - 使用表格格式
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.SetTitle("GFL 最终配置")
		t.SetStyle(table.StyleRounded)
		t.Style().Options.SeparateRows = true
		t.Style().Options.DrawBorder = true

		t.AppendHeader(table.Row{"配置项", "最终值", "来源"})

		// 确定每个配置项的来源
		getSource := func(field string) string {
			// 检查各个配置文件（按优先级从高到低）
			for i := len(configInfo.Sources) - 1; i >= 0; i-- {
				source := configInfo.Sources[i]
				if !source.Exists {
					continue
				}

				switch field {
				case "debug":
					if source.Config.Debug {
						return source.Name
					}
				case "devBaseBranch":
					if source.Config.DevBaseBranch != "" {
						return source.Name
					}
				case "productionBranch":
					if source.Config.ProductionBranch != "" {
						return source.Name
					}
				case "nickname":
					if source.Config.Nickname != "" {
						return source.Name
					}
				case "featurePrefix":
					if source.Config.FeaturePrefix != "" {
						return source.Name
					}
				case "fixPrefix":
					if source.Config.FixPrefix != "" {
						return source.Name
					}
				case "hotfixPrefix":
					if source.Config.HotfixPrefix != "" {
						return source.Name
					}
				}
			}

			return "默认值"
		}

		t.AppendRow(table.Row{"调试模式", fmt.Sprintf("%v", finalConfig.Debug), getSource("debug")})
		t.AppendRow(table.Row{"开发基础分支", finalConfig.DevBaseBranch, getSource("devBaseBranch")})
		t.AppendRow(table.Row{"生产分支", finalConfig.ProductionBranch, getSource("productionBranch")})
		t.AppendRow(table.Row{"昵称", finalConfig.Nickname, getSource("nickname")})
		t.AppendRow(table.Row{"功能分支前缀", finalConfig.FeaturePrefix, getSource("featurePrefix")})
		t.AppendRow(table.Row{"修复分支前缀", finalConfig.FixPrefix, getSource("fixPrefix")})
		t.AppendRow(table.Row{"热修复分支前缀", finalConfig.HotfixPrefix, getSource("hotfixPrefix")})

		t.AppendSeparator()
		t.AppendRow(table.Row{"示例功能分支", utils.GenerateBranchName(&finalConfig, "feature", "new-feature"), ""})

		t.Render()

		// 2. 显示配置来源详情 - 简化列表格式
		fmt.Printf("\n配置来源详情:\n\n")

		for _, source := range configInfo.Sources {
			if source.Exists {
				fmt.Printf("  • %s: %s\n", source.Name, source.Path)
			}
		}

		// GFL_CONFIG_FILE 环境变量
		configFile := os.Getenv("GFL_CONFIG_FILE")
		if configFile != "" {
			fmt.Printf("  • 自定义配置: %s (GFL_CONFIG_FILE)\n", configFile)
		}

		// 3. 显示配置优先级说明
		fmt.Printf("\n配置优先级 (从高到低):\n")
		fmt.Printf("  1. 自定义配置文件 (GFL_CONFIG_FILE)\n")
		fmt.Printf("  2. 本地配置文件 (.gfl.config.local.yml)\n")
		fmt.Printf("  3. 全局配置文件 (.gfl.config.yml)\n")
		fmt.Printf("  4. 默认值\n")
	},
}


func init() {
	rootCmd.AddCommand(configCmd)
}