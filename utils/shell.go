package utils

import (
	"fmt"
	"github.com/briandowns/spinner"
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

	// 解析命令和参数
	cmdArgs := strings.Fields(command)
	fmt.Println("cmdArgs:", cmdArgs, len(cmdArgs))
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
