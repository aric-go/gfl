package utils

import "os"

func AddGitIgnore() {
	// test has .gitignore file
	if _, err := os.Stat(".gitignore"); os.IsNotExist(err) {
		return
	}

	// add `.gflow.config.yml` to `.gitignore`
	f, err := os.OpenFile(".gitignore", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return
	}
	defer f.Close()

	_, err = f.WriteString("\n.gflow.config.yml\n")
	if err != nil {
		return
	}
}
