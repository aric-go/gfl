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
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...) // 第一个元素是命令，后面的元素是参数

	if err := cmd.Run(); err != nil {
		spin.Stop()
		return fmt.Errorf("执行命令失败: %w", err)
	}

	spin.Stop()
	return nil
}
