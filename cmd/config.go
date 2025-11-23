package cmd

import (
	"fmt"
	"gfl/utils"
	"os"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "查看当前配置",
	Long:  `显示当前 GFL 的配置信息，包括所有配置来源和最终值`,
	Run: func(cmd *cobra.Command, args []string) {
		configInfo := utils.ReadConfigWithSources()
		finalConfig := configInfo.FinalConfig

		// 1. 显示最终配置
		fmt.Printf("📋 GFL 最终配置\n\n")

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

		fmt.Printf("  调试模式       : %v (%s)\n", finalConfig.Debug, getSource("debug"))
		fmt.Printf("  开发基础分支   : %s (%s)\n", finalConfig.DevBaseBranch, getSource("devBaseBranch"))
		fmt.Printf("  生产分支       : %s (%s)\n", finalConfig.ProductionBranch, getSource("productionBranch"))
		fmt.Printf("  昵称           : %s (%s)\n", finalConfig.Nickname, getSource("nickname"))
		fmt.Printf("  功能分支前缀   : %s (%s)\n", finalConfig.FeaturePrefix, getSource("featurePrefix"))
		fmt.Printf("  修复分支前缀   : %s (%s)\n", finalConfig.FixPrefix, getSource("fixPrefix"))
		fmt.Printf("  热修复分支前缀 : %s (%s)\n", finalConfig.HotfixPrefix, getSource("hotfixPrefix"))
		fmt.Printf("  示例功能分支   : %s\n", utils.GenerateBranchName(&finalConfig, "feature", "new-feature"))

		// 2. 显示配置来源详情
		fmt.Printf("\n📁 配置来源详情\n\n")

		for _, source := range configInfo.Sources {
			status := "✅ 存在"
			if !source.Exists {
				status = "❌ 不存在"
			}

			fmt.Printf("  %s (%s)\n", source.Name, status)
			fmt.Printf("    路径: %s\n", source.Path)

			if source.Exists {
				values := []string{}
				if source.Config.Debug {
					values = append(values, "debug=true")
				}
				if source.Config.DevBaseBranch != "" {
					values = append(values, fmt.Sprintf("devBaseBranch=%s", source.Config.DevBaseBranch))
				}
				if source.Config.ProductionBranch != "" {
					values = append(values, fmt.Sprintf("productionBranch=%s", source.Config.ProductionBranch))
				}
				if source.Config.Nickname != "" {
					values = append(values, fmt.Sprintf("nickname=%s", source.Config.Nickname))
				}
				if source.Config.FeaturePrefix != "" {
					values = append(values, fmt.Sprintf("featurePrefix=%s", source.Config.FeaturePrefix))
				}
				if source.Config.FixPrefix != "" {
					values = append(values, fmt.Sprintf("fixPrefix=%s", source.Config.FixPrefix))
				}
				if source.Config.HotfixPrefix != "" {
					values = append(values, fmt.Sprintf("hotfixPrefix=%s", source.Config.HotfixPrefix))
				}

				if len(values) > 0 {
					fmt.Printf("    配置: %s\n", joinStrings(values, ", "))
				} else {
					fmt.Printf("    配置: (无)\n")
				}
			} else {
				fmt.Printf("    配置: -\n")
			}
			fmt.Println()
		}

		// GFL_CONFIG_FILE 环境变量
		configFile := os.Getenv("GFL_CONFIG_FILE")
		if configFile != "" {
			fmt.Printf("  🔧 配置文件环境变量 (GFL_CONFIG_FILE)\n")
			fmt.Printf("    状态: ✅ 活动\n")
			fmt.Printf("    值: %s\n\n", configFile)
		}

		// 3. 显示配置优先级说明
		fmt.Printf("🏆 配置优先级（从高到低）\n\n")
		fmt.Printf("  1️⃣  自定义配置文件 - GFL_CONFIG_FILE 环境变量指定\n")
		fmt.Printf("  2️⃣  本地配置文件   - .gfl.config.local.yml\n")
		fmt.Printf("  3️⃣  全局配置文件   - .gfl.config.yml\n")
		fmt.Printf("  4️⃣  默认值         - 内置默认配置\n")
	},
}

// joinStrings 连接字符串数组（简单的 strings.Join 替代）
func joinStrings(strs []string, sep string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	result := strs[0]
	for i := 1; i < len(strs); i++ {
		result += sep + strs[i]
	}
	return result
}

func init() {
	rootCmd.AddCommand(configCmd)
}