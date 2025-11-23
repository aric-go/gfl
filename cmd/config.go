package cmd

import (
	"fmt"
	"gfl/utils"
	"os"

	"github.com/fatih/color"
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

		t.AppendHeader(table.Row{
			color.New(color.FgCyan, color.Bold).Sprint("配置项"),
			color.New(color.FgGreen, color.Bold).Sprint("最终值"),
			color.New(color.FgMagenta, color.Bold).Sprint("来源"),
		})

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

		// 辅助函数：为来源添加颜色
		colorizeSource := func(source string) string {
			switch source {
			case "自定义配置":
				return color.New(color.FgRed, color.Bold).Sprint(source)
			case "本地配置":
				return color.New(color.FgYellow, color.Bold).Sprint(source)
			case "全局配置":
				return color.New(color.FgBlue, color.Bold).Sprint(source)
			case "默认值":
				return color.New(color.FgCyan).Sprint(source)
			default:
				return source
			}
		}

		// 辅助函数：为值添加颜色
		colorizeValue := func(value string, source string) string {
			switch source {
			case "自定义配置":
				return color.New(color.FgRed).Sprint(value)
			case "本地配置":
				return color.New(color.FgYellow).Sprint(value)
			case "全局配置":
				return color.New(color.FgBlue).Sprint(value)
			default:
				return value
			}
		}

		debugSource := getSource("debug")
		t.AppendRow(table.Row{
			"调试模式",
			fmt.Sprintf("%v", finalConfig.Debug),
			colorizeSource(debugSource),
		})

		devBaseSource := getSource("devBaseBranch")
		t.AppendRow(table.Row{
			"开发基础分支",
			colorizeValue(finalConfig.DevBaseBranch, devBaseSource),
			colorizeSource(devBaseSource),
		})

		prodSource := getSource("productionBranch")
		t.AppendRow(table.Row{
			"生产分支",
			colorizeValue(finalConfig.ProductionBranch, prodSource),
			colorizeSource(prodSource),
		})

		nicknameSource := getSource("nickname")
		t.AppendRow(table.Row{
			"昵称",
			colorizeValue(finalConfig.Nickname, nicknameSource),
			colorizeSource(nicknameSource),
		})

		featureSource := getSource("featurePrefix")
		t.AppendRow(table.Row{
			"功能分支前缀",
			colorizeValue(finalConfig.FeaturePrefix, featureSource),
			colorizeSource(featureSource),
		})

		fixSource := getSource("fixPrefix")
		t.AppendRow(table.Row{
			"修复分支前缀",
			colorizeValue(finalConfig.FixPrefix, fixSource),
			colorizeSource(fixSource),
		})

		hotfixSource := getSource("hotfixPrefix")
		t.AppendRow(table.Row{
			"热修复分支前缀",
			colorizeValue(finalConfig.HotfixPrefix, hotfixSource),
			colorizeSource(hotfixSource),
		})

		t.AppendSeparator()
		exampleBranch := utils.GenerateBranchName(&finalConfig, "feature", "new-feature")
		t.AppendRow(table.Row{
			"示例功能分支",
			color.New(color.FgGreen, color.Bold).Sprint(exampleBranch),
			"",
		})

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