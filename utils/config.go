package utils

import (
  "fmt"
  "os"

  "gopkg.in/yaml.v3"
)

type PublishItem struct {
  Name   string `yaml:"name"`
  Source string `yaml:"source"`
  Target string `yaml:"target"`
}

type YamlConfig struct {
  Debug            bool          `yaml:"debug"`
  DevBaseBranch    string        `yaml:"devBaseBranch"`
  ProductionBranch string        `yaml:"productionBranch"`
  Nickname         string        `yaml:"nickname"`
  GitlabHost       string        `yaml:"gitlabHost"`
  PublishList      []PublishItem `yaml:"publishList"`
}

func ReadConfig() *YamlConfig {
  data, err := os.ReadFile(".gflow.config.yml")
  if err != nil {
    fmt.Println("读取配置文件失败:", err)
    return nil
  }

  var config YamlConfig
  if err := yaml.Unmarshal(data, &config); err != nil {
    fmt.Println("解析配置文件失败:", err)
    return nil
  }

  return &config
}
