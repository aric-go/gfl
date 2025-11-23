# GFL Publish 命令技术文档

## 概述

`gfl publish` 命令用于将当前本地分支发布到远程仓库，设置上游跟踪分支，并支持别名 `p`。

## 实现原理

### 1. 命令定义

```go
var publishCmd = &cobra.Command{
    Use:     "publish",
    Aliases: []string{"p"},
    Short:   "Publish current branch (alias: p)",
    Run: func(cmd *cobra.Command, args []string) {
        // 实现逻辑
    },
}
```

### 2. 执行流程

#### 步骤 1: 发布当前分支
执行单个 Git 命令完成分支发布操作。

### 3. Shell 命令执行过程

#### 核心命令: 发布分支
```bash
git push -u origin HEAD
```

**命令解析**:
- `git push`: 推送本地提交到远程仓库
- `-u`: 设置上游分支跟踪（--set-upstream 的缩写）
- `origin`: 远程仓库名称
- `HEAD`: 指向当前分支的引用

**命令效果**:
1. 将当前分支的所有提交推送到远程仓库
2. 在远程仓库创建同名分支
3. 设置本地分支与远程分支的跟踪关系
4. 后续可以直接使用 `git push` 进行推送

**加载动画**: 显示 "正在推送..."
**成功消息**: 显示 "分支发布成功"

### 4. 完整 Shell 命令展示原理

以下是 `gfl publish` 命令的完整执行过程和对应的 shell 命令：

#### 步骤 0: 初始化检查
```bash
# GFL 内部执行的检查命令
# 检查是否在 Git 仓库中
git rev-parse --is-inside-work-tree

# 获取当前分支名称
CURRENT_BRANCH=$(git branch --show-current)

# 检查工作目录状态
git status --porcelain

# 检查是否在 detached HEAD 状态
git symbolic-ref -q HEAD || echo "Detached HEAD state"
```

#### 步骤 1: 获取当前分支信息
```bash
# GFL 获取当前分支信息，等效的 shell 命令：
CURRENT_BRANCH=$(git rev-parse --abbrev-ref HEAD)
echo "Current branch: $CURRENT_BRANCH"

# 检查是否有未提交的更改
if [[ -n $(git status --porcelain) ]]; then
    echo "Warning: Working directory is not clean"
fi

# 检查远程状态
git ls-remote --heads origin "$CURRENT_BRANCH"
```

#### 步骤 2: 执行发布操作
```bash
# GFL 执行发布分支，等效的 shell 命令：
# 方法 1: 使用 HEAD（GFL 实际使用的方式）
git push -u origin HEAD

# 方法 2: 使用分支名称（等效方式）
git push -u origin "$CURRENT_BRANCH"

# 方法 3: 分解步骤
# 2.1 推送分支并设置上游
git push --set-upstream origin "$CURRENT_BRANCH"

# 2.2 验证推送结果
git branch -vv | grep "$CURRENT_BRANCH"
```

#### 步骤 3: 完整执行示例
```bash
# 用户执行
$ gfl publish

# GFL 内部执行序列:
# 1. 检查 Git 仓库
$ git rev-parse --is-inside-work-tree
true

# 2. 获取当前分支
$ CURRENT_BRANCH=$(git branch --show-current)
$ echo "Current branch: $CURRENT_BRANCH"
feature/aric/user-authentication

# 3. 检查分支状态
$ git status --porcelain
# （可能有未提交的更改，但不影响推送）

# 4. 执行发布
$ git push -u origin HEAD
Enumerating objects: 7, done.
Counting objects: 100% (7/7), done.
Delta compression using up to 8 threads
Compressing objects: 100% (4/4), done.
Writing objects: 100% (7/7), 1.2 KiB | 1.2 MiB/s, done.
Total 7 (delta 3), reused 0 (delta 0)
remote:
remote: Create a pull request for 'feature/aric/user-authentication' on GitHub by visiting:
remote:      https://github.com/user/repo/pull/new/feature/aric/user-authentication
remote:
To github.com:user/repo.git
 * [new branch]      HEAD -> feature/aric/user-authentication
Branch 'feature/aric/user-authentication' set up to track remote branch 'feature/aric/user-authentication' from 'origin'.

# 5. 验证结果
$ git branch -vv
* feature/aric/user-authentication a1b2c3d [origin/feature/aric/user-authentication] Add user authentication feature
  main                       d4e5f6g [origin/main] Initial commit
  develop                    e7f8g9h [origin/develop] Setup development

# GFL 输出成功信息
# ✓ Branch published successfully!
```

#### 已发布分支的重新发布场景
```bash
# 用户执行
$ gfl publish

# GFL 检查分支状态
$ git branch -vv
* feature/aric/user-authentication a1b2c3d [origin/feature/aric/user-authentication] Add user authentication feature

# 执行推送（分支已存在）
$ git push -u origin HEAD
Everything up-to-date

# 或者有新提交时
$ git push -u origin HEAD
Enumerating objects: 3, done.
Counting objects: 100% (3/3), done.
Delta compression using up to 8 threads
Compressing objects: 100% (2/2), done.
Writing objects: 100% (2/2), 512 bytes | 512.00 KiB/s, done.
Total 2 (delta 1), reused 0 (delta 0)
To github.com:user/repo.git
   a1b2c3d..b4c5d6e  feature/aric/user-authentication -> feature/aric/user-authentication
```

#### 错误处理场景

##### 场景 1: 不是 Git 仓库
```bash
# GFL 检查仓库状态
$ git rev-parse --is-inside-work-tree
fatal: not a git repository (or any of the parent directories): .git

# GFL 显示错误
# Error: Not a Git repository. Please run 'git init' first.
```

##### 场景 2: Detached HEAD 状态
```bash
# GFL 检测到 detached HEAD
$ git symbolic-ref -q HEAD
# （没有输出，表示处于 detached HEAD）

$ git branch --show-current
# （没有输出）

# GFL 显示错误
# Error: Cannot publish from detached HEAD state. Please checkout a branch first.
```

##### 场景 3: 权限不足
```bash
# GFL 尝试推送但权限不足
$ git push -u origin HEAD
ERROR: Permission to user/repo.git denied to user.
fatal: Could not read from remote repository.

Please make sure you have the correct access rights
and the repository exists.

# GFL 显示错误
# Error: Permission denied. Please check your repository access rights.
```

##### 场景 4: 网络连接问题
```bash
# GFL 尝试推送但网络失败
$ git push -u origin HEAD
ssh: connect to host github.com port 22: Connection timed out
fatal: Could not read from remote repository.

Please make sure you have the correct access rights
and the repository exists.

# GFL 显示错误
# Error: Network connection failed. Please check your network and try again.
```

##### 场景 5: 远程仓库不存在
```bash
# GFL 尝试推送但远程仓库不存在
$ git push -u origin HEAD
ERROR: Repository not found.

fatal: Could not read from remote repository.

Please make sure you have the correct access rights
and the repository exists.

# GFL 显示错误
# Error: Remote repository not found. Please check your remote configuration.
```

#### 特殊场景：推送冲突
```bash
# GFL 检测到远程分支有不同的提交
$ git push -u origin HEAD
! [rejected]        HEAD -> feature/aric/user-authentication (non-fast-forward)
error: failed to push some refs to 'github.com:user/repo.git'
hint: Updates were rejected because the remote contains work that you do
hint: not have locally. You may want to first integrate the remote changes
hint: (e.g., 'git pull ...') before pushing again.
hint: See the 'Note about fast-forwards' in 'git push --help' for details.

# GFL 显示错误和解决建议
# Error: Push rejected. Remote branch has diverged.
# Suggestion: Run 'git pull' to integrate remote changes before pushing.
```

## 常用参数含义

此命令不接受任何参数，自动操作当前分支。

## 使用场景

### 1. 新功能分支首次发布
```bash
# 创建新功能分支
gfl start user-authentication

# 进行开发工作...
git add .
git commit -m "Add user authentication feature"

# 发布分支
gfl publish
```

### 2. 功能开发过程中的同步
```bash
# 继续开发...
git add .
git commit -m "Update authentication logic"

# 推送新提交（首次发布后可直接使用 git push）
git push

# 或者重新发布（设置上游关系）
gfl publish
```

### 3. 协作开发中的分支共享
```bash
# 发布分支供团队成员查看
gfl publish

# 团队成员可以查看和拉取
git checkout -b feature/aric/user-authentication origin/feature/aric/user-authentication
```

## 注意事项

### 1. 前置条件
- 必须在 Git 仓库中执行
- 当前必须在某个分支上（不能在 detached HEAD 状态）
- 需要远程仓库的写入权限
- 网络连接正常

### 2. 分支状态检查
- 当前分支有未提交的更改时，推送的是已提交的内容
- 未提交的更改不会被推送
- 建议在发布前先提交所有更改

### 3. 远程分支冲突
- 如果远程已存在同名分支且有冲突，推送会失败
- 需要先拉取远程更改并解决冲突
- 可能需要强制推送（不建议在生产环境中使用）

### 4. 权限要求
- 需要对远程仓库的写入权限
- 可能需要配置 SSH 密钥或个人访问令牌

### 5. 分支命名规范
- 遵循项目的分支命名约定
- 避免使用特殊字符和空格
- 建议使用有意义的描述性名称

## 工作流程集成

### 典型的 GitHub Flow 工作流
```bash
# 1. 开始新功能
gfl start user-profile

# 2. 开发功能
# ... 进行编码工作 ...
git add .
git commit -m "Implement user profile page"

# 3. 发布分支
gfl publish

# 4. 创建 Pull Request
gfl pr

# 5. 代码审查和合并
# ... 团队审查 ...

# 6. 清理分支
gfl sweep user-profile --local
```

### 与其他命令的协作
- `gfl start`: 创建需要发布的新分支
- `gfl pr`: 发布后创建 Pull Request
- `gfl sync`: 同步远程仓库状态
- `gfl checkout`: 切换到其他分支

## 使用示例

### 基本使用
```bash
gfl publish
```

### 使用别名
```bash
gfl p
```

### 完整工作流示例
```bash
# 1. 创建功能分支
gfl start payment-integration

# 2. 进行开发
echo "Payment code" > payment.go
git add payment.go
git commit -m "Add payment integration"

# 3. 发布分支
gfl publish
# 输出: 正在推送...
# 输出: 分支发布成功

# 4. 创建 PR
gfl pr
```

## 错误处理

### 常见错误及解决方案

1. **"权限被拒绝"**
   ```bash
   # 检查 SSH 配置
   ssh -T git@github.com

   # 或配置 HTTPS 认证
   git remote set-url origin https://github.com/user/repo.git
   ```

2. **"远程分支已存在"**
   ```bash
   # 强制推送（谨慎使用）
   git push -f origin HEAD

   # 或先拉取然后推送
   git pull origin HEAD
   gfl publish
   ```

3. **"无 upstream 分支"**
   ```bash
   # 手动设置上游分支
   git push --set-upstream origin feature-name

   # 或使用 gfl publish 自动设置
   gfl publish
   ```

4. **"网络连接失败"**
   - 检查网络连接
   - 确认远程仓库 URL 正确
   - 检查防火墙和代理设置

## 性能考虑

### 1. 推送数据量
- 只推送当前分支与远程分支的差异
- 首次推送可能包含更多提交
- 后续推送只传输增量更改

### 2. 网络依赖
- 需要稳定的网络连接
- 大项目推送可能需要较长时间
- 可以使用 `--verbose` 查看详细进度

### 3. 并发安全
- 避免多人同时推送同一分支
- 建议使用分支策略避免冲突
- 及时同步远程更改

## 最佳实践

### 1. 发布前检查
```bash
# 检查分支状态
git status

# 检查将要推送的提交
git log origin/HEAD..HEAD

# 检查远程状态
gfl sync
```

### 2. 频繁发布
- 小步快跑，频繁发布
- 每个功能模块完成后立即发布
- 便于代码审查和早期反馈

### 3. 分支管理
- 及时清理已合并的分支
- 使用有意义的分支名称
- 遵循团队约定的工作流程

## 相关命令

- `git push`: 原生推送命令
- `git push -u origin HEAD`: 底层实际执行命令
- `gfl pr`: 创建 Pull Request
- `gfl sync`: 同步远程仓库
- `gfl start`: 创建新分支