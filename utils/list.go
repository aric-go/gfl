/**
 * @Author: aric 1290657123@qq.com
 * @Date: 2024-10-11 21:30:36
 * @LastEditors: aric 1290657123@qq.com
 * @LastEditTime: 2024-10-11 21:31:33
 */
package utils

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
)

func BuildCommandList(branches []string) {
	// the answers will be written to this struct
	answers := struct {
		Module string `survey:"branch"`
	}{}

	// the questions to ask
	var qs = []*survey.Question{
		{
			Name: "module",
			Prompt: &survey.Select{
				Message: "Choose a branch:",
				Options: branches,
			},
		},
	}

	// perform the questions
	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 打印选择的模块名称
	//bun create answers.Module --no-git --no-install
	command := fmt.Sprintf("bun create %s --no-git --no-install", answers.Module)
	if _, err := RunShell(command); err != nil {
		fmt.Println(err.Error())
	}
}
