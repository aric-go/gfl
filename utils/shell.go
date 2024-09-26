/**
 * @Author: aric 1290657123@qq.com
 * @Date: 2024-09-26 14:08:24
 * @LastEditors: aric 1290657123@qq.com
 * @LastEditTime: 2024-09-26 14:14:19
 * @FilePath: utils/shell.go
 */
package utils

import "os/exec"

func RunShell(cmd string) (string, error) {
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
