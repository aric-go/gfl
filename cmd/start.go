// cmd/start.go
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"os/exec"
)

type Config struct {
	DefaultBranch string `yaml:"defaultBranch"`
}

var startCmd = &cobra.Command{
	Use:   "start [feature-name]",
	Short: "Start a new feature",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		featureName := args[0]

		// 读取配置文件
		configFile, err := ioutil.ReadFile(".gflow.config.yml")
		if err != nil {
			fmt.Println("Error reading config file:", err)
			return
		}
		var config Config
		yaml.Unmarshal(configFile, &config)

		// 执行 git fetch
		execCmd("git", "fetch", "origin", config.DefaultBranch)

		// 创建新分支
		branchName := fmt.Sprintf("feature/%s", featureName)
		execCmd("git", "checkout", "-b", branchName, fmt.Sprintf("origin/%s", config.DefaultBranch))

		fmt.Printf("Feature branch '%s' started successfully.\n", branchName)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func execCmd(command string, args ...string) {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error executing %s: %v\n", command, err)
	}
}
