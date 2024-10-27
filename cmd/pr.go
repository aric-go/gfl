package cmd

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

// prCmd represents the pr command
var prCmd = &cobra.Command{
	Use:   "pr",
	Short: "打开 GitHub 的 PR 页面",
	Run: func(cmd *cobra.Command, args []string) {
		// 读取配置文件获取默认分支和仓库
		config := readConfig()
		if config == nil {
			return
		}

		var headBranch string

		// 取得参数
		argBranch := args[0]

		// 获取当前的分支名称
		currentBranch, err := getCurrentBranch()
		if err != nil {
			fmt.Println("无法获取当前分支:", err)
			return
		}

		if argBranch == "" {
			headBranch = currentBranch
		}

		// 生成 GitHub PR URL
		prURL := fmt.Sprintf("https://github.com/%s/compare/%s...%s?expand=1", config.Repository, config.DevBaseBranch, headBranch)

		// 打开浏览器
		err = openBrowser(prURL)
		if err != nil {
			fmt.Println("无法打开浏览器:", err)
		} else {
			fmt.Printf("已打开 PR 页面: %s\n", prURL)
		}
	},
}

func init() {
	rootCmd.AddCommand(prCmd)
}

// 获取当前的分支名称
func getCurrentBranch() (string, error) {
	// 执行 git 命令获取当前分支
	output, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		return "", err
	}

	// 去除换行符
	return strings.TrimSpace(string(output)), nil
}

// 打开浏览器函数
func openBrowser(url string) error {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("无法识别的操作系统")
	}

	return err
}
