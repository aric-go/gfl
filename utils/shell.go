package utils

import (
	"fmt"
	"github.com/briandowns/spinner"
	"log"
	"os/exec"
	"strings"
	"time"
)

var spin = spinner.New(spinner.CharSets[35], 200*time.Millisecond)

func RunShell(cmd string) (string, error) {
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func RunCommandWithSpin(command string, message string) error {
	_ = spin.Color("green")
	spin.Start()
	spin.Suffix = message
	config := ReadConfig()
	if config == nil {
		log.Fatalf("读取配置文件失败")
	}

	if config.Debug {
		fmt.Println("🌈 正在执行命令: ", command)
	}

	// 解析命令和参数
	cmdArgs := strings.Fields(command)
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...) // 第一个元素是命令，后面的元素是参数

	if err := cmd.Run(); err != nil {
		spin.Stop()
		return fmt.Errorf("执行命令失败: %w, 命令: %s", err, command)
	}

	spin.Stop()
	return nil
}

func GetLocalBranches() []string {
	output, err := RunShell("git branch")
	if err != nil {
		fmt.Println("执行命令失败:", err)
		return nil
	}

	// 将输出转换为字符串并按行分割
	branches := strings.Split(strings.TrimSpace(string(output)), "\n")

	return branches
}

func IsCommandAvailable(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}
