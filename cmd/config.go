package cmd

import (
	"fmt"
	"gfl/utils"
	"gfl/utils/strings"
	"os"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:     "config",
	Aliases: []string{"c"},
	Short:   "View current configuration",
	Long:    strings.GetPath("config.long"),
	Run: func(cmd *cobra.Command, args []string) {
		configInfo := utils.ReadConfigWithSources()
		finalConfig := configInfo.FinalConfig

		// 1. 显示最终配置 - 使用表格格式
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.SetTitle(strings.GetPath("config.title"))
		t.SetStyle(table.StyleRounded)
		t.Style().Options.SeparateRows = true
		t.Style().Options.DrawBorder = true

		t.AppendHeader(table.Row{
			color.New(color.FgCyan, color.Bold).Sprint(strings.GetPath("config.config_key")),
			color.New(color.FgGreen, color.Bold).Sprint(strings.GetPath("config.final_value")),
			color.New(color.FgMagenta, color.Bold).Sprint(strings.GetPath("config.source")),
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
					// 如果显式设置了 nickname（包括空字符串），则作为来源
					if source.Config.NicknameSet {
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
				case "branchCaseFormat":
					if source.Config.BranchCaseFormat != "" {
						return source.Name
					}
				}
			}

			return strings.GetPath("config.default_value")
		}

		// 辅助函数：为来源添加颜色
		colorizeSource := func(source string) string {
			switch source {
			case strings.GetPath("config.custom_config"):
				return color.New(color.FgRed, color.Bold).Sprint(source)
			case strings.GetPath("config.local_config"):
				return color.New(color.FgYellow, color.Bold).Sprint(source)
			case strings.GetPath("config.global_config"):
				return color.New(color.FgBlue, color.Bold).Sprint(source)
			case strings.GetPath("config.default_value"):
				return color.New(color.FgCyan).Sprint(source)
			default:
				return source
			}
		}

		// 辅助函数：为值添加颜色
		colorizeValue := func(value string, source string) string {
			switch source {
			case strings.GetPath("config.custom_config"):
				return color.New(color.FgRed).Sprint(value)
			case strings.GetPath("config.local_config"):
				return color.New(color.FgYellow).Sprint(value)
			case strings.GetPath("config.global_config"):
				return color.New(color.FgBlue).Sprint(value)
			default:
				return value
			}
		}

		debugSource := getSource("debug")
		t.AppendRow(table.Row{
			strings.GetPath("config.debug_mode"),
			fmt.Sprintf("%v", finalConfig.Debug),
			colorizeSource(debugSource),
		})

		devBaseSource := getSource("devBaseBranch")
		t.AppendRow(table.Row{
			strings.GetPath("config.develop_base_branch"),
			colorizeValue(finalConfig.DevBaseBranch, devBaseSource),
			colorizeSource(devBaseSource),
		})

		prodSource := getSource("productionBranch")
		t.AppendRow(table.Row{
			strings.GetPath("config.production_branch"),
			colorizeValue(finalConfig.ProductionBranch, prodSource),
			colorizeSource(prodSource),
		})

		nicknameSource := getSource("nickname")
		t.AppendRow(table.Row{
			strings.GetPath("config.nickname"),
			colorizeValue(finalConfig.Nickname, nicknameSource),
			colorizeSource(nicknameSource),
		})

		featureSource := getSource("featurePrefix")
		t.AppendRow(table.Row{
			strings.GetPath("config.feature_prefix"),
			colorizeValue(finalConfig.FeaturePrefix, featureSource),
			colorizeSource(featureSource),
		})

		fixSource := getSource("fixPrefix")
		t.AppendRow(table.Row{
			strings.GetPath("config.fix_prefix"),
			colorizeValue(finalConfig.FixPrefix, fixSource),
			colorizeSource(fixSource),
		})

		hotfixSource := getSource("hotfixPrefix")
		t.AppendRow(table.Row{
			strings.GetPath("config.hotfix_prefix"),
			colorizeValue(finalConfig.HotfixPrefix, hotfixSource),
			colorizeSource(hotfixSource),
		})

		caseFormatSource := getSource("branchCaseFormat")
		t.AppendRow(table.Row{
			strings.GetPath("config.branch_case_format"),
			colorizeValue(finalConfig.BranchCaseFormat, caseFormatSource),
			colorizeSource(caseFormatSource),
		})

		t.AppendSeparator()
		exampleBranch := utils.GenerateBranchName(&finalConfig, "feature", "new-feature")
		t.AppendRow(table.Row{
			strings.GetPath("config.example_feature_branch"),
			color.New(color.FgGreen, color.Bold).Sprint(exampleBranch),
			"",
		})

		t.Render()

		// 2. 显示配置来源详情 - 简化列表格式
		fmt.Printf(strings.GetPath("config.config_sources_title"))

		for _, source := range configInfo.Sources {
			if source.Exists {
				var emoji string
				switch source.Name {
				case strings.GetPath("config.global_config"):
					emoji = "🌍"
				case strings.GetPath("config.local_config"):
					emoji = "🏠"
				case strings.GetPath("config.custom_config"):
					emoji = "🎯"
				default:
					emoji = "📄"
				}
				fmt.Printf("  %s %s: %s\n", emoji, source.Name, source.Path)
			}
		}

		// GFL_CONFIG_FILE 环境变量
		configFile := os.Getenv("GFL_CONFIG_FILE")
		if configFile != "" {
			fmt.Print(strings.GetPath("config.custom_config_file", configFile))
		}

		// 3. 显示配置优先级说明
		fmt.Printf(strings.GetPath("config.priority_title"))
		fmt.Printf(strings.GetPath("config.priority_custom"))
		fmt.Printf(strings.GetPath("config.priority_local"))
		fmt.Printf(strings.GetPath("config.priority_global"))
		fmt.Printf(strings.GetPath("config.priority_default"))
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
