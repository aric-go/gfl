# GFL 快速开始

本指南将帮助您快速上手 GFL (GitHub Flow CLI) 工具。

## 安装

### 从源码安装

```bash
git clone https://github.com/your-repo/gfl.git
cd gfl
go install
```

### 使用预编译二进制文件

从 [Releases](https://github.com/your-repo/gfl/releases) 页面下载适合您系统的二进制文件。

## 初始化

1. **进入您的项目目录**
   ```bash
   cd your-project
   ```

2. **初始化 GFL 配置**
   ```bash
   gfl init
   ```

3. **配置您的昵称（可选）**
   ```bash
   gfl init --nickname yourname
   ```

## 基本工作流程

### 1. 开始新功能开发

```bash
# 创建新功能分支
gfl start user-authentication

# 这将创建名为 feature/yourname/user-authentication 的分支
# 并自动切换到该分支
```

### 2. 开发和提交

```bash
# 进行您的开发工作
# 添加文件、提交更改
git add .
git commit -m "feat: add user authentication"
```

### 3. 发布分支

```bash
# 将分支推送到远程仓库
gfl publish

# 或者使用别名
gfl p
```

### 4. 创建 Pull Request

```bash
# 创建 PR 并在浏览器中打开
gfl pr --open

# 或者使用别名
gfl rv -o
```

## 常用命令速查

| 命令 | 功能 | 示例 |
|------|------|------|
| `gfl start` | 创建功能分支 | `gfl s login` |
| `gfl publish` | 发布当前分支 | `gfl p` |
| `gfl pr` | 创建 Pull Request | `gfl rv --open` |
| `gfl checkout` | 交互式切换分支 | `gfl co` |
| `gfl hotfix` | 创建热修复分支 | `gfl hf bug-fix` |
| `gfl release` | 创建发布版本 | `gfl release -t minor` |
| `gfl tag` | 创建版本标签 | `gfl tag -t patch` |
| `gfl sweep` | 清理分支 | `gfl sweep feature -y` |

## 下一步

- 查看 [完整命令文档](commands.md) 了解所有可用命令
- 阅读 [配置指南](configuration.md) 自定义 GFL 行为
- 了解 [最佳实践](best-practices.md) 提升工作效率

## 遇到问题？

如果遇到任何问题，请：
1. 查看命令帮助：`gfl help [command]`
2. 启用调试模式：在配置文件中设置 `debug: true`
3. 提交 Issue：[GitHub Issues](https://github.com/your-repo/gfl/issues)