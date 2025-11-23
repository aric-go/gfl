# GFL Start 命令技术文档

## 概述

`gfl start` 命令用于创建新的功能分支，是 GitHub Flow 工作流程的起点。该命令支持别名 `s`。

## 实现原理

### 1. 命令定义

```go
var startCmd = &cobra.Command{
    Use:     "start [feature-name]",
    Short:   "Start a new feature (alias: s)",
    Aliases: []string{"s"},
    Args:    cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        // 实现逻辑
    },
}
```

### 2. 执行流程

#### 步骤 1: 配置读取
```go
config := utils.ReadConfig()
if config == nil {
    return
}
```

#### 步骤 2: 功能名称解析
支持两种命名格式：
- **简单格式**: `new-feature` → `feature/new-feature`
- **冒号格式**: `feat:login-page` → `feat/login-page`

```go
type StartName struct {
    ActionName  string  // 功能类型
    FeatureName string  // 功能名称
}

func parseStartName(name string) *StartName {
    hasColon := str.Contains(name, ":")
    if hasColon {
        parts := str.Split(name, ":")
        return &StartName{
            ActionName:  parts[0],
            FeatureName: parts[1],
        }
    } else {
        return &StartName{
            ActionName:  "feature",
            FeatureName: name,
        }
    }
}
```

#### 步骤 3: 分支名称生成
```go
branchName := utils.GenerateBranchName(config, startName.ActionName, startName.FeatureName)
baseRemoteBranch := fmt.Sprintf("origin/%s", config.DevBaseBranch)
```

#### 步骤 4: Git 操作执行

### 3. Shell 命令执行过程

#### 命令 1: 同步远程仓库
```bash
git fetch origin
```
- **目的**: 更新远程分支引用
- **加载动画**: 显示 "正在同步..."
- **错误处理**: 失败时中止执行

#### 命令 2: 创建并切换到新分支
```bash
git checkout -b feature/aric/new-feature origin/develop
```
- **参数解释**:
  - `-b`: 创建新分支
  - `feature/aric/new-feature`: 生成的分支名称
  - `origin/develop`: 基础分支（来自配置）
- **加载动画**: 显示 "正在创建分支..."
- **成功消息**: 显示功能类型和分支名称

### 4. 完整 Shell 命令展示原理

以下是 `gfl start` 命令的完整执行过程和对应的 shell 命令：

#### 步骤 0: 初始化检查
```bash
# GFL 内部执行的检查命令
# 检查是否在 Git 仓库中
git rev-parse --is-inside-work-tree

# 检查配置文件是否存在
ls -la .gfl.config*.yml

# 检查当前分支状态
git status --porcelain
```

#### 步骤 1: 读取和解析配置
假设配置内容为：
```yaml
# .gfl.config.yml
devBaseBranch: "develop"
productionBranch: "main"
nickname: "aric"
featurePrefix: "feature"
```

```bash
# GFL 读取配置，等效的 shell 命令：
# 读取配置文件
DEV_BASE_BRANCH=$(grep "devBaseBranch" .gfl.config.yml | cut -d: -f2 | tr -d ' ')
NICKNAME=$(grep "nickname" .gfl.config.yml | cut -d: -f2 | tr -d ' ')
FEATURE_PREFIX=$(grep "featurePrefix" .gfl.config.yml | cut -d: -f2 | tr -d ' ')

echo "DEV_BASE_BRANCH: $DEV_BASE_BRANCH"    # 输出: develop
echo "NICKNAME: $NICKNAME"                  # 输出: aric
echo "FEATURE_PREFIX: $FEATURE_PREFIX"      # 输出: feature
```

#### 步骤 2: 解析功能名称
假设用户执行：`gfl start user-authentication`

```bash
# GFL 解析功能名称，等效的 shell 命令：
FEATURE_NAME="user-authentication"

# 检查是否包含冒号格式
if [[ "$FEATURE_NAME" == *:* ]]; then
    ACTION_NAME=$(echo "$FEATURE_NAME" | cut -d: -f1)
    FEATURE_NAME=$(echo "$FEATURE_NAME" | cut -d: -f2)
else
    ACTION_NAME="feature"
fi

echo "ACTION_NAME: $ACTION_NAME"      # 输出: feature
echo "FEATURE_NAME: $FEATURE_NAME"    # 输出: user-authentication
```

#### 步骤 3: 生成分支名称
```bash
# GFL 生成分支名称，等效的 shell 命令：
BRANCH_NAME="${ACTION_NAME}/${NICKNAME}/${FEATURE_NAME}"
BASE_REMOTE_BRANCH="origin/${DEV_BASE_BRANCH}"

echo "BRANCH_NAME: $BRANCH_NAME"          # 输出: feature/aric/user-authentication
echo "BASE_REMOTE_BRANCH: $BASE_REMOTE_BRANCH"  # 输出: origin/develop
```

#### 步骤 4: 执行 Git 操作
```bash
# 1. 同步远程仓库
git fetch origin

# 2. 创建并切换到新分支
git checkout -b "$BRANCH_NAME" "$BASE_REMOTE_BRANCH"

# 3. 验证创建结果
git branch --show-current
git log --oneline -1
```

#### 步骤 5: 完整执行示例
```bash
# 用户执行
$ gfl start user-authentication

# GFL 内部执行序列:
# 1. 检查 Git 仓库
$ git rev-parse --is-inside-work-tree
true

# 2. 检查配置文件
$ ls -la .gfl.config*.yml
-rw-r--r-- 1 user user 234 Dec 20 10:30 .gfl.config.yml

# 3. 读取配置
$ DEV_BASE_BRANCH=$(grep "devBaseBranch" .gfl.config.yml | cut -d: -f2 | tr -d ' ')
$ NICKNAME=$(grep "nickname" .gfl.config.local.yml | cut -d: -f2 | tr -d ' ')
$ echo "DEV_BASE_BRANCH: $DEV_BASE_BRANCH"
develop
$ echo "NICKNAME: $NICKNAME"
aric

# 4. 解析功能名称
$ FEATURE_NAME="user-authentication"
$ ACTION_NAME="feature"
$ BRANCH_NAME="feature/aric/user-authentication"

# 5. 同步远程仓库
$ git fetch origin
From github.com:user/repo
 * [new branch]      develop -> origin/develop

# 6. 创建并切换到新分支
$ git checkout -b feature/aric/user-authentication origin/develop
Branch 'feature/aric/user-authentication' set up to track remote branch 'develop' from 'origin'.
Switched to a new branch 'feature/aric/user-authentication'

# 7. 验证结果
$ git branch --show-current
feature/aric/user-authentication

# GFL 输出成功信息
# ✓ Started feature: feature/aric/user-authentication
```

#### 冒号格式示例
```bash
# 用户执行冒号格式
$ gfl start feat:login-page

# GFL 内部解析:
$ FEATURE_NAME="feat:login-page"
$ ACTION_NAME=$(echo "$FEATURE_NAME" | cut -d: -f1)  # feat
$ FEATURE_NAME=$(echo "$FEATURE_NAME" | cut -d: -f2) # login-page
$ BRANCH_NAME="feat/aric/login-page"

# 创建分支
$ git checkout -b feat/aric/login-page origin/develop
Branch 'feat/aric/login-page' set up to track remote branch 'develop' from 'origin'.
Switched to a new branch 'feat/aric/login-page'

# GFL 输出
# ✓ Started feat: feat/aric/login-page
```

#### 错误处理场景

##### 场景 1: 配置文件不存在
```bash
# 用户执行
$ gfl start new-feature

# GFL 检查配置文件
$ ls -la .gfl.config*.yml
ls: cannot access '.gfl.config*.yml': No such file or directory

# GFL 显示错误
# Error: No configuration file found. Please run 'gfl init' first.
```

##### 场景 2: 远程分支不存在
```bash
# 用户执行
$ gfl start new-feature

# GFL 同步远程仓库
$ git fetch origin
# fetch 成功，但 develop 分支不存在

# 尝试创建分支
$ git checkout -b feature/aric/new-feature origin/develop
fatal: 'origin/develop' is not a commit and a branch 'feature/aric/new-feature' cannot be created from it

# GFL 显示错误
# Error: Base branch 'origin/develop' does not exist. Please check your configuration.
```

##### 场景 3: 分支已存在
```bash
# 用户执行
$ gfl start existing-feature

# GFL 检测到分支已存在
$ git branch --list "feature/aric/existing-feature"
feature/aric/existing-feature

# 尝试创建分支
$ git checkout -b feature/aric/existing-feature origin/develop
fatal: A branch named 'feature/aric/existing-feature' already exists.

# GFL 显示错误
# Error: Branch 'feature/aric/existing-feature' already exists.
```

##### 场景 4: 工作目录不干净
```bash
# GFL 检查工作目录状态
$ git status --porcelain
 M src/app.js
?? new-file.js

# 虽然 git checkout -b 在有未提交更改时仍然工作，
# GFL 会显示警告
# Warning: Working directory is not clean. You may want to commit or stash changes.
```

#### 特殊配置场景
```bash
# 假设配置文件中有不同的前缀
# .gfl.config.yml
featurePrefix: "feat"
nickname: "bob"
devBaseBranch: "dev"

# 用户执行
$ gfl start payment-system

# GFL 生成分支名称
$ BRANCH_NAME="feat/bob/payment-system"
$ BASE_REMOTE_BRANCH="origin/dev"

# 创建分支
$ git checkout -b feat/bob/payment-system origin/dev
```

## 常用参数含义

### 位置参数: [feature-name]
- **类型**: `string`
- **必填**: 是
- **说明**: 功能分支的名称
- **格式**:
  - 简单格式: `new-feature`
  - 冒号格式: `feat:login-page`

## 分支命名规范

### 命名模式
```
{action}/{nickname}/{feature-name}
```

### 示例
- `feature/aric/user-authentication`
- `feat/aric/payment-integration`
- `fix/aric/login-bug`

### 配置项影响
- `nickname`: 开发者昵称
- `featurePrefix`: 功能前缀（默认 "feature"）
- `fixPrefix`: 修复前缀（默认 "fix"）

## 注意事项

### 1. 前置条件
- 必须在 Git 仓库中执行
- 必须先运行 `gfl init` 初始化配置
- 需要网络连接访问远程仓库

### 2. 分支策略
- 基于 `devBaseBranch` 配置的分支创建（通常是 `develop`）
- 新分支自动推送到远程仓库并设置上游跟踪

### 3. 错误处理
- 配置文件缺失时中止执行
- Git 命令执行失败时显示错误信息
- 网络连接问题会导致创建失败

### 4. 工作流程集成
- 此命令是 GitHub Flow 的起始点
- 创建的分支用于后续的 `publish`、`pr` 等操作

## 使用示例

### 基本使用
```bash
gfl start user-authentication
```
创建分支: `feature/aric/user-authentication`

### 使用冒号格式
```bash
gfl start feat:payment-system
```
创建分支: `feat/aric/payment-system`

### 使用别名
```bash
gfl s login-page
```
创建分支: `feature/aric/login-page`

### 创建修复分支
```bash
gfl start fix:memory-leak
```
创建分支: `fix/aric/memory-leak`

## 相关命令

- `gfl publish`: 发布当前分支
- `gfl pr`: 创建 Pull Request
- `gfl checkout`: 切换分支
- `gfl config`: 查看配置

## 配置依赖

以下配置项影响 start 命令的行为：
- `devBaseBranch`: 基础开发分支
- `nickname`: 开发者昵称
- `featurePrefix`: 功能分支前缀
- `fixPrefix`: 修复分支前缀

## 故障排除

### 常见错误及解决方案

1. **"无法找到配置文件"**
   ```bash
   gfl init
   ```

2. **"远程分支不存在"**
   ```bash
   gfl sync
   gfl start feature-name
   ```

3. **"分支已存在"**
   ```bash
   git checkout existing-branch
   # 或使用其他功能名称
   gfl start different-feature
   ```

4. **"网络连接失败"**
   - 检查网络连接
   - 确认远程仓库访问权限
   - 检查 SSH 密钥配置