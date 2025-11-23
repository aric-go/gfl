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