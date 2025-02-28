package utils

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type CreateGflConfigOptions struct {
	Filename     string
	Force        bool
	AddGitIgnore bool
}

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

// CreateGflConfig 创建配置文件
func CreateGflConfig(config YamlConfig, opts CreateGflConfigOptions) error {
	// 检查文件是否存在
	if _, err := os.Stat(opts.Filename); err == nil && !opts.Force {
		return fmt.Errorf("配置文件 %s 已存在，如需覆盖请使用 force 选项", opts.Filename)
	}

	// 创建或覆盖配置文件
	file, err := os.Create(opts.Filename)
	if err != nil {
		return fmt.Errorf("无法创建配置文件: %w", err)
	}
	defer file.Close()

	// 序列化配置
	data, err := yaml.Marshal(&config)
	if err != nil {
		return fmt.Errorf("无法生成 YAML: %w", err)
	}

	// 写入配置
	if _, err := file.Write(data); err != nil {
		return fmt.Errorf("无法写入配置文件: %w", err)
	}

	// 如果需要，添加到 .gitignore
	if opts.AddGitIgnore {
		if f, err := os.OpenFile(".gitignore", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600); err == nil {
			defer f.Close()
			f.WriteString(fmt.Sprintf("\n%s\n", opts.Filename))
		}
	}

	return nil
}
