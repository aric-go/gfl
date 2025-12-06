# GFL PR 命令技术文档

## 概述

`gfl pr` 命令用于创建 GitHub Pull Request，支持多种工作流程模式，包括同步分支、打开 PR 列表和创建新 PR。支持别名 `rv`。

## 实现原理

### 1. 命令定义

```go
var prCmd = &cobra.Command{
    Use:     "pr [base-branch]",
    Aliases: []string{"rv"},
    Short:   "Create pull request (PR)",
    Run: func(cmd *cobra.Command, args []string) {
        // 实现逻辑
    },
}
```

### 2. 执行流程

#### 步骤 1: 配置和仓库检查
```go
config := utils.ReadConfig()
if config == nil {
    return
}
repo, _ := utils.GetRepository()
```

#### 步骤 2: 参数解析和分支处理
根据命令标志决定执行不同的操作：

1. **`--sync` 模式**: 同步生产分支到开发分支
2. **`--open` 模式**: 打开 PR 列表页面
3. **默认模式**: 创建新的 Pull Request

#### 步骤 3: 参数验证和处理
- 获取当前分支名称
- 确定目标分支（配置中的开发分支或指定参数）
- 执行相应的操作

### 3. Shell 命令执行过程

#### 模式 1: 同步分支 (`--sync`)
当使用 `--sync` 标志时：

```go
if isSync {
    if !utils.SyncProductionToDev(config.ProductionBranch, config.DevBaseBranch) {
        utils.Errorf(strings.GetPath("pr.sync_failed"))
    }
    return
}
```

**内部执行的命令**:
```bash
git fetch origin
git checkout develop
git pull origin develop
git merge main
git push origin develop
```

#### 模式 2: 打开 PR 列表 (`--open`)
```go
if isOpen {
    listUrl := fmt.Sprintf("https://github.com/%s/pulls", repo)
    err := browser.OpenURL(listUrl)
    if err != nil {
        utils.Errorf(strings.GetPath("pr.browser_error"), err)
        return
    }
    return
}
```

**执行的浏览器操作**:
- 打开 URL: `https://github.com/{repo-owner}/{repo-name}/pulls`

#### 模式 3: 创建 PR（默认模式）
```go
currentBranch, err := utils.GetCurrentBranch()
var baseBranch = config.DevBaseBranch
if len(args) > 0 {
    baseBranch = args[0]
}
utils.CreatePr(baseBranch, currentBranch)
```

**内部执行的命令**:
```bash
# 获取当前分支信息
git rev-parse --abbrev-ref HEAD

# 检查分支是否已发布
git ls-remote --heads origin feature/aric/user-auth

# 创建 PR（使用 GitHub CLI）
gh pr create --base develop --head feature/aric/user-auth --fill
```

## 常用参数含义

### 位置参数: [base-branch]
- **类型**: `string`
- **可选**: 是
- **说明**: 指定 PR 的目标分支
- **默认值**: 配置文件中的 `devBaseBranch`（通常是 `develop`）
- **示例**:
  ```bash
  gfl pr main        # 向 main 分支创建 PR
  gfl pr develop     # 向 develop 分支创建 PR
  ```

### `--sync, -s`
- **类型**: `bool`
- **默认值**: `false`
- **说明**: 同步生产分支到开发分支
- **功能**: 将 main 分支的更改合并到 develop 分支
- **使用场景**: 在发布周期开始前同步分支

### `--open, -o`
- **类型**: `bool`
- **默认值**: `false`
- **说明**: 在浏览器中打开 Pull Request 列表页面
- **功能**: 快速访问现有的 PR

## 使用场景

### 1. 创建功能 PR
```bash
# 开发完成后创建 PR
gfl pr
```

### 2. 向特定分支创建 PR
```bash
# 向 main 分支创建 PR（直接发布）
gfl pr main
```

### 3. 分支同步
```bash
# 将生产分支的更改同步到开发分支
gfl pr --sync
```

### 4. 查看现有 PR
```bash
# 在浏览器中打开 PR 列表
gfl pr --open
```

### 5. 使用别名
```bash
gfl rv              # 创建 PR
gfl rv --sync       # 同步分支
gfl rv --open       # 打开 PR 列表
```

## 注意事项

### 1. 前置条件
- 必须在 Git 仓库中执行
- 当前分支必须已推送到远程仓库
- 需要安装 GitHub CLI (`gh`) 用于创建 PR
- 需要浏览器支持（用于 `--open` 模式）

### 2. 权限要求
- 需要对 GitHub 仓库的写入权限
- GitHub CLI 需要已认证
- 浏览器访问需要网络连接

### 3. 分支状态
- 当前分支不能是目标分支
- 当前分支必须在远程仓库存在
- 分支间不能存在冲突（会提示先解决冲突）

### 4. 工具依赖
- **GitHub CLI (`gh`)**: 用于创建 PR
- **浏览器**: 用于打开 PR 列表
- **网络连接**: 与 GitHub API 通信

## 工作流程集成

### 标准功能开发流程
```bash
# 1. 开始新功能
gfl start user-authentication

# 2. 开发和提交
git add .
git commit -m "Implement user authentication"

# 3. 发布分支
gfl publish

# 4. 创建 PR
gfl pr

# 5. 等待代码审查...

# 6. 合并后清理分支
gfl sweep user-authentication --local
```

### 版本发布流程
```bash
# 1. 同步生产分支到开发分支
gfl pr --sync

# 2. 创建发布分支
gfl release --type patch

# 3. 创建标签
gfl tag

# 4. 向 main 分支创建发布 PR
gfl pr main
```

## 错误处理

### 常见错误及解决方案

1. **"GitHub CLI 未安装"**
   ```bash
   # macOS
   brew install gh

   # Ubuntu
   sudo apt install gh

   # 认证
   gh auth login
   ```

2. **"分支未发布"**
   ```bash
   # 先发布分支
   gfl publish
   # 然后创建 PR
   gfl pr
   ```

3. **"权限不足"**
   ```bash
   # 检查 GitHub CLI 认证
   gh auth status

   # 重新认证
   gh auth login
   ```

4. **"PR 已存在"**
   - GitHub CLI 会自动检测现有 PR
   - 可以使用 `--open` 查看现有 PR

5. **"合并冲突"**
   ```bash
   # 同步目标分支
   git fetch origin
   git merge origin/develop

   # 解决冲突后重新发布
   gfl publish
   gfl pr
   ```

## 高级用法

### 1. 批量操作
```bash
# 同步分支后创建 PR
gfl pr --sync && gfl pr
```

### 2. 脚本化使用
```bash
#!/bin/bash
# 自动化 PR 创建脚本
gfl publish
gfl pr --open  # 打开 PR 页面进行手动检查
```

### 3. 与 CI/CD 集成
```yaml
# GitHub Actions 示例
- name: Create PR
  run: |
    gfl pr
  if: github.event_name == 'push' && startsWith(github.ref, 'refs/heads/feature/')
```

## 配置依赖

以下配置项影响 PR 命令：
- `devBaseBranch`: 默认目标分支
- `productionBranch`: 同步操作的源分支

## 最佳实践

### 1. PR 准备
```bash
# 创建 PR 前检查
git status
gfl sync
gfl publish
```

### 2. PR 标题和描述
- GitHub CLI 会自动从提交信息生成 PR 标题
- 可以手动编辑 PR 描述
- 建议包含测试步骤和相关链接

### 3. 代码审查
- 及时响应审查意见
- 保持 PR 大小适中
- 包含必要的测试和文档

## 相关命令

- `gh pr create`: 底层使用的 GitHub CLI 命令
- `gfl publish`: 发布分支
- `gfl checkout`: 切换分支
- `gfl sync`: 同步仓库
- `gfl sweep`: 清理分支