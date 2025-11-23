# GFL Sweep 命令技术文档

## 概述

`gfl sweep` 命令用于清理包含特定关键词的分支，支持清理本地分支和远程分支，别名包括 `clean` 和 `rm`。

## 实现原理

### 1. 命令定义

```go
var sweepCmd = &cobra.Command{
    Use:     "sweep [keyword]",
    Aliases: []string{"clean", "rm"},
    Short:   "Clean branches containing specific keywords",
    Args:    cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        // 实现逻辑
    },
}
```

### 2. 执行流程

#### 步骤 1: 参数解析和验证
```go
keyword := args[0]
confirm, _ := cmd.Flags().GetBool("confirm")

if !localFlag && !remoteFlag {
    utils.Error(strings.GetString("sweep", "local_remote_required"))
    return
}
```

#### 步骤 2: 分支清理操作
根据标志分别执行本地和远程分支清理：
```go
if localFlag {
    cleanLocalBranches(keyword, confirm)
}

if remoteFlag {
    cleanRemoteBranches(keyword, confirm)
}
```

### 3. Shell 命令执行过程

#### 本地分支清理命令
```bash
git branch
```
- **目的**: 获取所有本地分支列表
- **输出格式**: 每行一个分支名称，当前分支前有 `*`
- **示例输出**:
  ```
  * main
    develop
    feature/aric/user-auth
    feature/aric/payment-system
    fix/aric/login-bug
  ```

#### 本地分支删除命令
```bash
git branch -d feature/aric/user-auth
```
- **参数解析**:
  - `-d`: 安全删除（确保分支已合并）
  - `feature/aric/user-auth`: 要删除的分支名称
- **安全机制**: 如果分支未合并，删除会失败

#### 远程分支列表命令
```bash
git branch -r
```
- **目的**: 获取所有远程分支列表
- **输出格式**: `origin/分支名` 格式
- **示例输出**:
  ```
    origin/HEAD -> origin/main
    origin/main
    origin/develop
    origin/feature/aric/user-auth
    origin/fix/aric/login-bug
  ```

#### 远程分支删除命令
```bash
git push origin --delete feature/aric/user-auth
```
- **参数解析**:
  - `git push origin`: 推送到远程仓库
  - `--delete`: 删除指定分支
  - `feature/aric/user-auth`: 要删除的远程分支名

## 常用参数含义

### 位置参数: [keyword]
- **类型**: `string`
- **必填**: 是
- **说明**: 用于匹配要删除的分支的关键词
- **匹配方式**: 包含匹配（分支名称包含关键词即匹配）
- **示例**:
  ```bash
  gfl sweep user-auth      # 删除包含 "user-auth" 的分支
  gfl sweep feature        # 删除所有功能分支
  gfl sweep aric          # 删除开发者 aric 的所有分支
  ```

### `--local, -l`
- **类型**: `bool`
- **默认值**: `false`
- **说明**: 清理本地分支
- **必须配合**: `--remote` 或单独使用，但不能两个都不指定

### `--remote, -r`
- **类型**: `bool`
- **默认值**: `false`
- **说明**: 清理远程分支
- **必须配合**: `--local` 或单独使用，但不能两个都不指定

### `--confirm, -y` (全局标志)
- **类型**: `bool`
- **默认值**: `false`
- **说明**: 确认执行删除操作
- **效果**:
  - `true`: 实际删除匹配的分支
  - `false`: 只显示将要删除的分支（预览模式）

## 使用场景

### 1. 清理已合并的功能分支
```bash
# 预览要删除的功能分支
gfl sweep feature --local

# 确认删除
gfl sweep feature --local --confirm

# 同时清理本地和远程
gfl sweep feature --local --remote --confirm
```

### 2. 清理特定开发者的分支
```bash
# 清理开发者 aric 的所有分支
gfl sweep aric --local --confirm

# 只清理 aric 的功能分支
gfl sweep feature/aric --remote --confirm
```

### 3. 清理特定功能的分支
```bash
# 清理所有与用户认证相关的分支
gfl sweep auth --local --remote --confirm

# 清理测试分支
gfl sweep test --local --confirm
```

### 4. 使用别名
```bash
# 使用 clean 别名
gfl clean feature --local --confirm

# 使用 rm 别名
gfl rm hotfix --remote --confirm
```

## 注意事项

### 1. 安全考虑
- **预览模式**: 不使用 `--confirm` 时只显示匹配的分支，不实际删除
- **安全删除**: 本地分支使用 `git branch -d`（安全删除，确保已合并）
- **远程删除**: 远程分支删除是不可逆的，需要谨慎操作
- **当前分支保护**: 不能删除当前所在的分支

### 2. 前置条件
- 必须在 Git 仓库中执行
- 需要对远程仓库的写入权限（删除远程分支时）
- 工作目录应该相对干净

### 3. 分支状态
- 本地未合并的分支删除会失败（安全机制）
- 远程分支删除需要谨慎，无法恢复
- 建议先使用预览模式查看匹配的分支

### 4. 网络依赖
- 清理远程分支需要网络连接
- 远程删除操作可能需要认证

### 5. 批量操作风险
- 关键词匹配可能匹配到意外的分支
- 建议使用具体的关键词
- 重要分支建议手动删除

## 分支匹配逻辑

### 本地分支匹配
```go
for _, branch := range str.Split(string(branches), "\n") {
    branch = str.TrimSpace(branch)
    if branch == "" {
        continue
    }
    if str.Contains(branch, keyword) {
        // 匹配成功，准备删除
    }
}
```

### 远程分支匹配
```go
for _, branch := range str.Split(string(branches), "\n") {
    branch = str.TrimSpace(branch)
    if str.Contains(branch, keyword) {
        remoteBranch := str.TrimPrefix(branch, "origin/")
        // 删除远程分支
    }
}
```

### 匹配示例
```bash
# 关键词: "feature"
# 匹配的分支:
#   - feature/aric/user-auth
#   - feature/user-payment
#   - hotfix/feature-bug  # 也会匹配（包含 "feature"）

# 关键词: "aric"
# 匹配的分支:
#   - feature/aric/user-auth
#   - fix/aric/login-bug
#   - hotfix/aric/critical-fix
```

## 使用示例

### 基本使用
```bash
# 预览要删除的功能分支
gfl sweep feature --local
# 输出: 将删除: feature/aric/user-auth
# 输出: 将删除: feature/aric/payment-system
# 输出: 跳过确认操作

# 确认删除
gfl sweep feature --local --confirm
# 输出: 正在删除本地分支...
# 输出: 分支 feature/aric/user-auth 删除成功
# 输出: 分支 feature/aric/payment-system 删除成功
```

### 清理远程分支
```bash
# 预览远程分支删除
gfl sweep feature --remote
# 输出: 将删除: origin/feature/aric/user-auth

# 确认删除远程分支
gfl sweep feature --remote --confirm
# 输出: 正在删除远程分支...
# 输出: 分支 origin/feature/aric/user-auth 删除成功
```

### 同时清理本地和远程
```bash
gfl sweep completed --local --remote --confirm
# 清理所有包含 "completed" 的本地和远程分支
```

### 使用别名
```bash
gfl clean test --local --confirm      # 等同于 gfl sweep
gfl rm old-feature --remote --confirm  # 等同于 gfl sweep
```

## 错误处理

### 常见错误及解决方案

1. **"必须指定 --local 或 --remote"**
   ```bash
   # 错误用法
   gfl sweep feature

   # 正确用法
   gfl sweep feature --local
   gfl sweep feature --remote
   gfl sweep feature --local --remote
   ```

2. **"无法删除当前分支"**
   ```bash
   # 切换到其他分支
   git checkout main
   gfl sweep feature --local --confirm
   ```

3. **"分支未合并，删除失败"**
   ```bash
   # 强制删除（谨慎使用）
   git branch -D branch-name

   # 或先合并再删除
   git merge branch-name
   gfl sweep feature --local --confirm
   ```

4. **"远程分支删除权限不足"**
   ```bash
   # 检查权限
   git remote show origin

   # 配置认证
   ssh -T git@github.com
   ```

5. **"匹配到重要分支"**
   ```bash
   # 使用更具体的关键词
   gfl sweep "feature/aric/old" --local --confirm
   # 而不是
   gfl sweep feature --local --confirm
   ```

## 最佳实践

### 1. 安全删除流程
```bash
# 1. 预览匹配的分支
gfl sweep feature --local

# 2. 仔细检查匹配结果
# 确认没有匹配到重要分支

# 3. 确认删除
gfl sweep feature --local --confirm
```

### 2. 定期清理策略
```bash
# 每周清理已合并的功能分支
gfl sweep "feature/" --local --remote --confirm

# 每月清理修复分支
gfl sweep "fix/" --local --confirm
```

### 3. 关键词选择策略
```bash
# 使用具体的关键词
gfl sweep "completed-feature" --local --confirm

# 避免使用过于宽泛的关键词
# 不要这样: gfl sweep f --local --confirm

# 使用开发者名称
gfl sweep "aric/merged" --local --confirm
```

### 4. 团队协作建议
- 在团队中建立分支清理约定
- 定期清理已合并的分支
- 重要分支添加保护标记
- 使用分支前缀规范命名

## 高级用法

### 1. 条件清理脚本
```bash
#!/bin/bash
# 清理已超过30天的功能分支

for branch in $(git branch --format='%(refname:short)' | grep "feature/"); do
    last_commit=$(git log -1 --format='%ct' $branch)
    current_time=$(date +%s)
    age=$((current_time - last_commit))

    if [ $age -gt $((30 * 24 * 3600)) ]; then
        echo "清理30天未更新的分支: $branch"
        git branch -d $branch
    fi
done
```

### 2. 批量操作
```bash
# 一次性清理多种类型的分支
for keyword in "completed" "old" "test" "demo"; do
    gfl sweep $keyword --local --confirm
done
```

### 3. 与 CI/CD 集成
```yaml
# GitHub Actions 定期清理
name: Cleanup Old Branches
on:
  schedule:
    - cron: '0 0 * * 0'  # 每周日午夜

jobs:
  cleanup:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - name: Cleanup completed features
        run: gfl sweep completed --local --confirm
```

## 相关命令

- `git branch`: 原生分支管理命令
- `git branch -d`: 安全删除分支
- `git push origin --delete`: 删除远程分支
- `gfl checkout`: 切换分支
- `gfl start`: 创建新分支

## 配置依赖

- 无特定配置依赖
- 分支清理基于 Git 仓库状态

## 与其他命令的协作

### 与 Prune 操作结合
```bash
# GFL sweep + Git prune 完整清理
gfl sweep completed --local --confirm
git remote prune origin
```

### 与 Checkout 结合
```bash
# 切换分支并清理
git checkout main
gfl sweep feature --local --confirm
```