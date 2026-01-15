# gfl forward

将 main 分支的代码 forward（转发）到 dev 分支，通过创建 Pull Request 的方式实现。

## 用法

```bash
gfl forward [flags]
gfl fwd [flags]     # 别名，更简洁的写法
```

## 选项

- `-t, --title string`: PR 标题（可选，默认: "Sync main to dev"）
- `-b, --body string`: PR 描述（可选）

## 功能

该命令会：
1. 基于远程分支创建从 `origin/main` 到 `origin/dev` 的 Pull Request
2. 使用默认或自定义的 PR 标题和描述
3. 通过 GitHub CLI (gh) 创建 PR

## 前置条件

- 已安装 [GitHub CLI](https://cli.github.com/)
- 配置文件中 `devBaseBranch` 和 `productionBranch` 不能相同
- 已通过 `gh auth login` 完成 GitHub 认证

## 示例

### 基本用法（使用默认标题和描述）
```bash
gfl forward
# 或使用别名
gfl fwd
```
创建标题为 "Sync main to dev" 的 PR

### 自定义 PR 标题
```bash
gfl forward --title "Hotfix sync to dev"
# 或使用简写
gfl fwd -t "Hotfix sync to dev"
```

### 自定义标题和描述
```bash
gfl forward --title "Sync release changes" --body "Forwarding release branch changes to dev"
# 或使用简写
gfl fwd -t "Sync release changes" -b "Forwarding release branch changes to dev"
```

## 默认值

- **默认标题**: `Sync {productionBranch} to {devBaseBranch}`
- **默认描述**: `Forwarding changes from `{productionBranch}` to `{devBaseBranch}`.`

## 配置依赖

以下配置项影响 forward 命令的行为：
- `devBaseBranch`: 目标开发分支（默认: `dev`）
- `productionBranch`: 源生产分支（默认: `main`）

## 相关命令

- `gfl sync`: 同步远程仓库到本地
- `gfl pr`: 创建 Pull Request
- `gfl release`: 创建 Release

## 使用场景

### 将已发布的代码同步到开发分支
当 main 分支的代码已经发布并打上 tag 后，需要将这些更改同步回 dev 分支：

```bash
# 1. 确保远程分支已同步
gfl sync

# 2. 创建 forward PR
gfl forward

# 3. 审查并合并 PR
```

### 将热修复同步到开发分支
当使用 `gfl hotfix` 修复生产环境问题后，需要将修复内容同步到开发分支：

```bash
# 完成热修复并合并到 main 后
gfl forward --title "Sync hotfix to dev"
```

## 错误处理

### 开发分支和生产分支相同
如果配置文件中 `devBaseBranch` 和 `productionBranch` 相同，会提示：
```
ERROR: 开发分支和发布分支不能相同，请检查配置文件
```

### GitHub CLI 未安装或未认证
如果 gh CLI 未安装或未完成认证，PR 创建会失败并显示错误信息。

## 工作流程

```bash
# 典型的 forward 工作流
gfl sync                    # 1. 同步远程分支
gfl forward                 # 2. 创建 forward PR
# 在 GitHub 上审查并合并 PR
```
