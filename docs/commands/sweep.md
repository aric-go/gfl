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

### 4. 完整 Shell 命令展示原理

以下是 `gfl sweep` 命令的完整执行过程和对应的 shell 命令：

#### 步骤 0: 初始化检查
```bash
# GFL 内部执行的检查命令
# 检查是否在 Git 仓库中
git rev-parse --is-inside-work-tree

# 检查当前分支
CURRENT_BRANCH=$(git branch --show-current)

# 检查工作目录状态
git status --porcelain
```

#### 步骤 1: 参数解析和验证
假设用户执行：`gfl sweep feature --local --remote --confirm`

```bash
# GFL 解析参数，等效的 shell 命令：
KEYWORD="feature"
LOCAL_FLAG=true
REMOTE_FLAG=true
CONFIRM_FLAG=true

# 参数验证
if [[ "$LOCAL_FLAG" == false && "$REMOTE_FLAG" == false ]]; then
    echo "Error: Must specify --local or --remote"
    exit 1
fi

echo "Keyword: $KEYWORD"
echo "Local cleanup: $LOCAL_FLAG"
echo "Remote cleanup: $REMOTE_FLAG"
echo "Confirm deletion: $CONFIRM_FLAG"
```

#### 步骤 2: 本地分支清理（如果指定了 --local）
```bash
# GFL 获取本地分支列表
LOCAL_BRANCHES=$(git branch --format='%(refname:short)')

echo "Found local branches:"
echo "$LOCAL_BRANCHES"

# 匹配包含关键词的分支
MATCHED_LOCAL_BRANCHES=$(echo "$LOCAL_BRANCHES" | grep "$KEYWORD" || true)

if [[ -n "$MATCHED_LOCAL_BRANCHES" ]]; then
    echo "Matched local branches:"
    echo "$MATCHED_LOCAL_BRANCHES"

    # 排除当前分支
    FILTERED_BRANCHES=$(echo "$MATCHED_LOCAL_BRANCHES" | grep -v "^$CURRENT_BRANCH$" || true)

    if [[ "$CONFIRM_FLAG" == true && -n "$FILTERED_BRANCHES" ]]; then
        echo "Deleting local branches..."
        echo "$FILTERED_BRANCHES" | while read -r branch; do
            if [[ -n "$branch" && "$branch" != "$CURRENT_BRANCH" ]]; then
                echo "Deleting local branch: $branch"
                git branch -d "$branch" || echo "Failed to delete $branch (may not be merged)"
            fi
        done
    else
        echo "Preview mode - would delete these local branches:"
        echo "$FILTERED_BRANCHES"
    fi
else
    echo "No local branches matching '$KEYWORD' found"
fi
```

#### 步骤 3: 远程分支清理（如果指定了 --remote）
```bash
# GFL 获取远程分支列表
REMOTE_BRANCHES=$(git branch -r --format='%(refname:short)')

echo "Found remote branches:"
echo "$REMOTE_BRANCHES"

# 匹配包含关键词的远程分支
MATCHED_REMOTE_BRANCHES=$(echo "$REMOTE_BRANCHES" | grep "$KEYWORD" | grep "origin/" || true)

if [[ -n "$MATCHED_REMOTE_BRANCHES" ]]; then
    echo "Matched remote branches:"
    echo "$MATCHED_REMOTE_BRANCHES"

    # 提取分支名称（去掉 origin/ 前缀）
    FILTERED_REMOTE_BRANCHES=$(echo "$MATCHED_REMOTE_BRANCHES" | sed 's|^origin/||' || true)

    if [[ "$CONFIRM_FLAG" == true && -n "$FILTERED_REMOTE_BRANCHES" ]]; then
        echo "Deleting remote branches..."
        echo "$FILTERED_REMOTE_BRANCHES" | while read -r branch; do
            if [[ -n "$branch" ]]; then
                echo "Deleting remote branch: $branch"
                git push origin --delete "$branch" || echo "Failed to delete remote branch $branch"
            fi
        done
    else
        echo "Preview mode - would delete these remote branches:"
        echo "$FILTERED_REMOTE_BRANCHES"
    fi
else
    echo "No remote branches matching '$KEYWORD' found"
fi
```

#### 步骤 4: 完整执行示例
```bash
# 用户执行
$ gfl sweep aric --local --remote --confirm

# GFL 内部执行序列:
# 1. 检查 Git 仓库
$ git rev-parse --is-inside-work-tree
true

# 2. 获取当前分支
$ CURRENT_BRANCH=$(git branch --show-current)
$ echo "Current branch: $CURRENT_BRANCH"
main

# 3. 获取本地分支
$ LOCAL_BRANCHES=$(git branch --format='%(refname:short)')
$ echo "Local branches:"
$ echo "$LOCAL_BRANCHES"
main
develop
feature/aric/user-auth
feature/aric/payment-system
feature/bob/dashboard
fix/aric/login-bug

# 4. 匹配本地分支
$ MATCHED_LOCAL=$(echo "$LOCAL_BRANCHES" | grep "aric")
$ echo "Matched local branches:"
$ echo "$MATCHED_LOCAL"
feature/aric/user-auth
feature/aric/payment-system
fix/aric/login-bug

# 5. 删除本地分支
$ git branch -d feature/aric/user-auth
Deleted branch feature/aric/user-auth (was a1b2c3d).

$ git branch -d feature/aric/payment-system
Deleted branch feature/aric/payment-system (was d4e5f6g).

$ git branch -d fix/aric/login-bug
Deleted branch fix/aric/login-bug (was e7f8g9h).

# 6. 获取远程分支
$ REMOTE_BRANCHES=$(git branch -r --format='%(refname:short)')
$ echo "Remote branches:"
$ echo "$REMOTE_BRANCHES"
origin/HEAD -> origin/main
origin/main
origin/develop
origin/feature/aric/user-auth
origin/feature/aric/payment-system
origin/feature/bob/dashboard
origin/fix/aric/login-bug

# 7. 匹配远程分支
$ MATCHED_REMOTE=$(echo "$REMOTE_BRANCHES" | grep "aric" | grep "origin/")
$ echo "Matched remote branches:"
$ echo "$MATCHED_REMOTE"
origin/feature/aric/user-auth
origin/feature/aric/payment-system
origin/fix/aric/login-bug

# 8. 删除远程分支
$ git push origin --delete feature/aric/user-auth
To github.com:user/repo.git
 - [deleted]         feature/aric/user-auth

$ git push origin --delete feature/aric/payment-system
To github.com:user/repo.git
 - [deleted]         feature/aric/payment-system

$ git push origin --delete fix/aric/login-bug
To github.com:user/repo.git
 - [deleted]         fix/aric/login-bug

# GFL 输出结果
# ✓ Local branches deleted: feature/aric/user-auth, feature/aric/payment-system, fix/aric/login-bug
# ✓ Remote branches deleted: feature/aric/user-auth, feature/aric/payment-system, fix/aric/login-bug
```

#### 预览模式示例
```bash
# 用户执行预览模式
$ gfl sweep feature --local

# GFL 显示将要删除的分支，但不执行删除
# Preview mode - would delete these local branches:
# feature/aric/user-auth
# feature/bob/payment-system
# feature/alice/ui

# 提示: Use --confirm to actually delete these branches.
```

#### 错误处理场景

##### 场景 1: 未指定操作类型
```bash
# 用户执行
$ gfl sweep feature

# GFL 验证参数失败
# Error: Must specify --local or --remote (or both)
# Usage: gfl sweep <keyword> [--local] [--remote] [--confirm]
```

##### 场景 2: 尝试删除当前分支
```bash
# 当前分支是 feature/aric/current-work
$ git branch --show-current
feature/aric/current-work

# 尝试清理 aric 分支
$ gfl sweep aric --local --confirm

# GFL 过滤掉当前分支
# Preview mode - would delete these local branches:
# feature/aric/old-feature
# fix/aric/bug-fix
# (Skipped: feature/aric/current-work - current branch)
```

##### 场景 3: 本地分支未合并
```bash
# 尝试删除未合并的分支
$ git branch -d feature/aric/unmerged-feature
error: The branch 'feature/aric/unmerged-feature' is not fully merged.
If you are sure you want to delete it, run 'git branch -D feature/aric/unmerged-feature'.

# GFL 显示警告
# Warning: Branch feature/aric/unmerged-feature is not merged and was not deleted.
# Use 'git branch -D' to force delete unmerged branches.
```

##### 场景 4: 远程分支删除权限不足
```bash
# 尝试删除远程分支但权限不足
$ git push origin --delete feature/aric/protected-branch
ERROR: Permission to user/repo.git denied to user.
fatal: Could not read from remote repository.

# GFL 显示错误
# Error: Permission denied. Cannot delete remote branch 'feature/aric/protected-branch'.
# Please check your repository permissions.
```

##### 场景 5: 没有匹配的分支
```bash
# 搜索不存在的关键词
$ gfl sweep nonexistent --local --remote

# GFL 输出结果
# No local branches matching 'nonexistent' found
# No remote branches matching 'nonexistent' found
# Nothing to delete.
```

#### 高级使用场景

##### 场景 1: 批量清理已合并分支
```bash
# 清理所有已合并到 develop 的功能分支
$ gfl sweep "feature/" --local --confirm

# 只清理特定开发者的分支
$ gfl sweep "feature/aric/" --local --remote --confirm
```

##### 场景 2: 清理旧的测试分支
```bash
# 清理所有测试分支
$ gfl sweep "test-" --local --remote --confirm

# 清理所有带有 "temp" 关键词的分支
$ gfl sweep temp --local --confirm
```

##### 场景 3: 定期清理脚本
```bash
#!/bin/bash
# 定期清理脚本

# 清理一个月前的功能分支
echo "Cleaning old feature branches..."
gfl sweep "feature/" --local --confirm

# 清理已合并的修复分支
echo "Cleaning merged fix branches..."
gfl sweep "fix/" --local --confirm

echo "Cleanup completed!"
```

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