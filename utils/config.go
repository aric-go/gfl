package utils

import (
  "fmt"
  "github.com/afeiship/go-ipt"
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

type PublishOption struct {
  Label string
  Value PublishItem
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

func IptPublishList(config *YamlConfig) {
  publishList := config.PublishList
  isGitlab := config.GitlabHost != ""

  var opts []ipt.Option[PublishOption]
  for _, item := range publishList {
    opts = append(opts, ipt.Option[PublishOption]{
      Label: item.Name,
      Value: PublishOption{
        Label: item.Name,
        Value: item,
      },
    })
  }

  selected, err := ipt.Ipt("What is your favorite color?", opts)
  if err != nil {
    fmt.Println("选择发布项目终止:", err)
    return
  }

  var currentBranch string

  if selected.Value.Source == "current" {
    currentBranch, _ = GetCurrentBranch()
  } else {
    currentBranch = selected.Value.Source
  }

  if isGitlab {
    CreateMr(selected.Value.Target, currentBranch)
  } else {
    CreatePr(selected.Value.Target, currentBranch)
  }
}
