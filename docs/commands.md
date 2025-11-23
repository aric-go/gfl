# GFL 命令完整使用指南

GFL (GitHub Flow CLI) 是一个用于简化 Git 工作流程的命令行工具，基于 GitHub Flow 最佳实践。

## 安装和初始化

### 初始化配置

```bash
# 初始化 GFL 配置
gfl init

# 强制覆盖现有配置
gfl init --force

# 设置自定义昵称
gfl init --nickname myname
```

### 配置文件

GFL 使用两个配置文件：

- `.gfl.config.yml` - 全局配置（团队共享）
- `.gfl.config.local.yml` - 本地配置（个人覆盖）

#### 全局配置示例

```yaml
debug: false
devBaseBranch: dev
productionBranch: main
nickname: aric
featurePrefix: feature
fixPrefix: fix
hotfixPrefix: hotfix
```

#### 本地配置示例

```yaml
nickname: mynickname
featurePrefix: feat
fixPrefix: bugfix
hotfixPrefix: hot
```

## 核心命令

### 1. start - 开始新功能开发

创建新的功能分支并切换到该分支。

```bash
# 基本用法
gfl start feature-name

# 使用别名
gfl s feature-name
```

**功能说明：**
- 从开发基础分支（默认：dev）创建新分支
- 分支命名格式：`feature/<nickname>/<feature-name>` 或 `feature/<feature-name>`
- 自动获取最新的远程分支信息

**示例：**
```bash
$ gfl s user-authentication
✓ 已从 dev 创建分支 feature/aric/user-authentication
✓ 已切换到分支 feature/aric/user-authentication
```

### 2. checkout - 交互式分支切换

提供交互式界面选择和切换 Git 分支。

```bash
# 交互式选择分支
gfl checkout

# 使用别名
gfl co
```

**功能说明：**
- 显示所有本地和远程分支列表
- 支持键盘导航选择
- 选中后自动切换到目标分支

### 3. publish - 发布当前分支

将当前分支推送到远程仓库并设置上游分支。

```bash
# 发布当前分支
gfl publish

# 使用别名
gfl p
```

**功能说明：**
- 将当前分支推送到远程仓库
- 设置上游分支跟踪关系
- 自动处理分支名映射

### 4. pr - 创建 Pull Request

创建 GitHub Pull Request。

```bash
# 基本用法 - 推送当前分支
gfl pr

# 指定目标分支
gfl pr main

# 使用别名
gfl rv main

# 创建 PR 并在浏览器中打开
gfl pr --open
gfl pr -o

# 同步生产分支到开发分支
gfl pr --sync
```

**功能说明：**
- 自动推送当前分支到远程
- 打开 GitHub PR 创建页面
- 支持指定基础分支
- 可选择在浏览器中打开 PR 列表

### 5. release - 创建发布版本

基于最新标签创建新的发布版本。

```bash
# 创建补丁版本（默认）
gfl release

# 创建特定类型版本
gfl release --type patch    # v1.0.0 -> v1.0.1
gfl release --type minor    # v1.0.0 -> v1.1.0
gfl release --type major    # v1.0.0 -> v2.0.0

# 使用短选项
gfl release -t minor

# 创建热修复发布
gfl release --hotfix
gfl release -x
```

**功能说明：**
- 基于最新的语义版本标签创建新版本
- 自动递增版本号
- 创建发布分支 `releases/release-x.x.x`
- 支持主版本、次版本、补丁版本递增

### 6. tag - 创建版本标签

创建版本标签并推送到远程仓库。

```bash
# 创建补丁版本标签（默认）
gfl tag

# 创建特定类型版本标签
gfl tag --type patch    # v1.0.0 -> v1.0.1
gfl tag --type minor    # v1.0.0 -> v1.1.0
gfl tag --type major    # v1.0.0 -> v2.0.0

# 使用短选项
gfl tag -t minor
```

**功能说明：**
- 创建语义版本标签
- 推送标签到远程仓库
- 可选择创建 GitHub Release

### 7. hotfix - 创建热修复分支

从生产分支创建热修复分支。

```bash
# 创建热修复分支
gfl hotfix fix-critical-bug

# 使用别名
gfl hf fix-critical-bug
```

**功能说明：**
- 从生产分支（默认：main）创建分支
- 分支命名格式：`hotfix/<nickname>/<hotfix-name>`
- 用于紧急修复生产环境问题

### 8. sweep - 清理分支

清理包含特定关键词的本地和远程分支。

```bash
# 清理包含关键词的分支（默认清理本地）
gfl sweep feature

# 只清理本地分支
gfl sweep feature --local
gfl sweep feature -l

# 只清理远程分支
gfl sweep feature --remote
gfl sweep feature -r

# 同时清理本地和远程分支
gfl sweep feature --local --remote

# 自动确认操作（不询问）
gfl sweep feature --confirm
gfl sweep feature -y

# 使用别名
gfl clean feature
gfl rm feature
```

**功能说明：**
- 搜索包含指定关键词的分支
- 支持本地/远程分支清理
- 提供确认提示防止误删
- 不会删除当前所在分支

### 9. sync - 同步远程仓库

同步远程仓库并更新所有远程引用。

```bash
# 同步所有远程仓库
gfl sync
```

**功能说明：**
- 获取所有远程仓库的最新信息
- 更新远程分支引用
- 不影响本地工作分支

### 10. version - 显示版本信息

显示 GFL 工具的当前版本。

```bash
# 显示版本信息
gfl version

# 使用短选项
gfl -v
```

### 11. completion - 生成 Shell 自动补全

生成指定 Shell 的自动补全脚本。

```bash
# 生成 Bash 自动补全
gfl completion bash

# 生成 Zsh 自动补全
gfl completion zsh

# 生成 Fish 自动补全
gfl completion fish

# 生成 PowerShell 自动补全
gfl completion powershell
```

**安装自动补全：**

**Bash:**
```bash
# 永久添加到 ~/.bashrc
echo 'source <(gfl completion bash)' >> ~/.bashrc
```

**Zsh:**
```bash
# 永久添加到 ~/.zshrc
echo 'source <(gfl completion zsh)' >> ~/.zshrc
```

## 全局选项

所有命令都支持以下全局选项：

```bash
--confirm, -y      # 自动确认操作
--help, -h         # 显示帮助信息
--version, -v      # 显示版本信息
```

## 环境变量

- `GFL_CONFIG_FILE`: 指定自定义配置文件路径

```bash
# 使用自定义配置文件
export GFL_CONFIG_FILE=/path/to/custom-config.yml
gfl start new-feature
```

## 分支命名规范

GFL 遵循标准化的分支命名模式：

- **功能分支**: `feature/<nickname>/<name>` 或 `feature/<name>`
- **修复分支**: `fix/<nickname>/<name>` 或 `fix/<name>`
- **热修复分支**: `hotfix/<nickname>/<name>` 或 `hotfix/<name>`
- **发布分支**: `releases/release-x.x.x`
- **版本标签**: `v1.0.0`, `v1.0.1`, `v1.1.0` 等

## 实用技巧

### 快速开始开发流程

```bash
# 1. 开始新功能
gfl s user-profile

# 2. 开发完成后发布
gfl p

# 3. 创建 PR
gfl pr --open
```

### 紧急修复流程

```bash
# 1. 创建热修复分支
gfl hf security-patch

# 2. 修复问题后发布
gfl p

# 3. 创建紧急 PR
gfl pr main --open
```

### 版本发布流程

```bash
# 1. 创建发布版本
gfl release --type minor

# 2. 创建版本标签
gfl tag --type minor
```

### 清理工作区

```bash
# 清理所有已合并的功能分支
gfl sweep feature --confirm

# 清理所有远程已合并分支
gfl sweep --remote --confirm
```

## 故障排除

### 常见问题

1. **配置文件找不到**
   ```bash
   # 重新初始化配置
   gfl init --force
   ```

2. **分支同步失败**
   ```bash
   # 检查网络连接和仓库权限
   git remote -v

   # 手动同步后重试
   git fetch origin
   gfl start feature-name
   ```

3. **PR 创建失败**
   ```bash
   # 确保分支已发布
   gfl publish
   gfl pr
   ```

### 调试模式

在配置文件中启用调试模式：

```yaml
debug: true
```

这将输出详细的执行信息，帮助诊断问题。