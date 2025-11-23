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
		Error(err.Error())
		return
	}

	// git checkout the branch
	command := fmt.Sprintf("git checkout %s", answers.Module)
	if _, err := RunShell(command); err != nil {
		Error(err.Error())
	}
}
