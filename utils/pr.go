package utils

import (
	"fmt"
	"os/exec"
	"runtime"
)

func CreatePr(url string) {
	prURL := fmt.Sprintf(url)
	err := openBrowser(prURL)
	err = openBrowser(prURL)
	if err != nil {
		fmt.Println("无法打开浏览器:", err)
	} else {
		fmt.Printf("已打开 PR 页面: %s\n", prURL)
	}
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
