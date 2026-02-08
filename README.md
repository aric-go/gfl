# GFL - GitHub Flow CLI

![GFL Logo](./assets/logo.svg)

> A powerful command-line tool that simplifies GitHub Flow workflows.

**Note:** For web projects, add this to your HTML head:
```html
<link rel="icon" type="image/x-icon" href="/assets/favicon.ico">
```

![Alt text](./__uml__/img.png)

## 📚 Documentation

- [🚀 快速开始](docs/quick-start.md) - 5分钟快速上手
- [📖 完整命令文档](docs/commands.md) - 所有命令的详细用法
- [⚙️ 配置指南](docs/configuration.md) - 自定义 GFL 配置
- [💡 最佳实践](docs/best-practices.md) - 团队协作和工作流程建议

## 🔧 Workflows

- [功能开发流程](docs/quick-start.md#基本工作流程)
- [热修复流程](docs/best-practices.md#紧急修复流程)
- [版本发布流程](docs/best-practices.md#版本发布流程)

## release
```shell
# update tag
nrcip
# release bin file
goreleaser release --clean
# upload dist to oss bucket
oss://web-alo7-com/assets/bins/gfl-releases/

# list oss
oss://web-alo7-com/assets/bins/gfl-releases/
# upload oss
cd dist && aliyun oss sync . oss://web-alo7-com/assets/bins/gfl-releases/ --delete --force
```

## 🚀 Quick Start

### Installation

```bash
# Install from source
go install github.com/your-repo/gfl@latest

# Or download binary from releases
# https://github.com/your-repo/gfl/releases
```

### Initialize

```bash
cd your-project
gfl init --nickname yourname
```

### Basic Usage

```bash
# Start a new feature
gfl s user-authentication

# Publish your branch
gfl p

# Create a Pull Request
gfl pr --open

# Clean up merged branches
gfl sweep feature --confirm
```

## 📋 Command Reference

```bash
$ gfl -h
GFL - GitHub Flow CLI

Usage:
  gfl [flags]
  gfl [command]

Available Commands:
  checkout    交互式的git分支切换 (alias: co)
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  hotfix      开始一个hotfix分支 (alias: hf)
  init        初始化 Github Flow 配置
  publish     发布当前分支 (alias: p)
  release     创建发布版本
  pr          创建 Pull Request (alias: rv)
  start       开始一个新功能 (alias: s)
  sweep       清理包含特定关键词的分支 (alias: clean, rm)
  sync        同步远程仓库
  tag         创建版本标签
  version     获取程序版本

Flags:
  -h, --help      help for gfl
  -v, --version   show version

Use "gfl [command] --help" for more information about a command.
```

👉 **查看完整命令文档**: [docs/commands.md](docs/commands.md)

## ✨ Features

- 🔧 **智能分支管理** - 自动化的分支创建、命名和切换
- 🚀 **快速工作流** - 简化 GitHub Flow 的每个步骤
- 🔀 **PR 创建** - 一键创建 Pull Request 并在浏览器中打开
- 📦 **版本管理** - 语义化版本控制和发布管理
- 🧹 **分支清理** - 智能清理已合并和过期的分支
- ⚙️ **灵活配置** - 支持全局和本地配置文件
- 🎯 **交互式界面** - 友好的命令行交互体验
- 🔍 **调试支持** - 详细的执行日志和调试模式

## 🏗️ Project Structure

```
gfl/
├── main.go                    # Entry point
├── cmd/                       # Command implementations
│   ├── root.go               # Root command
│   ├── start.go              # Start feature command
│   ├── publish.go            # Publish branch command
│   ├── pr.go                 # PR creation command
│   ├── release.go            # Release management
│   └── ...                   # Other commands
├── utils/                     # Utility functions
│   ├── config.go             # Configuration management
│   ├── branch.go             # Branch naming utilities
│   ├── git.go                # Git operations
│   └── ...                   # Other utilities
├── docs/                      # Documentation
│   ├── quick-start.md        # Quick start guide
│   ├── commands.md           # Complete command reference
│   ├── configuration.md      # Configuration guide
│   └── best-practices.md     # Best practices
├── .gfl.config.yml          # Global configuration
├── .gfl.config.local.yml    # Local configuration
└── README.md                # This file
```

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`gfl s amazing-feature`)
3. Commit your changes (`git commit -m 'feat: add amazing feature'`)
4. Push to the branch (`gfl p`)
5. Open a Pull Request (`gfl pr --open`)

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

**🎉 Happy coding with GFL!**

需要帮助？查看我们的[完整文档](docs/)或在 [GitHub Issues](https://github.com/your-repo/gfl/issues) 提问。
