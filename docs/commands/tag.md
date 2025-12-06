# GFL Tag 命令技术文档

## 概述

`gfl tag` 命令用于基于最新标签生成新的版本标签，支持主版本、次版本和补丁版本的增量创建，并自动创建 GitHub Release。支持别名 `t`。

## 实现原理

### 1. 命令定义

```go
var tagCmd = &cobra.Command{
    Use:     "tag",
    Aliases: []string{"t"},
    Short:   "Generate new tag version for release branch based on latest tag",
    Run: func(cmd *cobra.Command, args []string) {
        // 实现逻辑
    },
}
```

### 2. 执行流程

#### 步骤 1: 版本计算
```go
version := utils.GetLatestVersion()
versionType, _ := cmd.Flags().GetString("type")
newVersion, err := utils.IncrementVersion(version, versionType)
```

#### 步骤 2: 版本信息显示
```go
utils.Infof(strings.GetPath("tag.previous_version"), version)
utils.Successf(strings.GetPath("tag.new_version"), newVersion)
```

#### 步骤 3: Git 操作执行
按照特定顺序执行多个 Git 命令完成版本发布

### 3. Shell 命令执行过程

#### 命令 1: 切换到发布分支
```bash
git checkout releases/release-v1.1.2
```
- **目的**: 切换到对应的发布分支
- **格式**: `releases/release-{newVersion}`
- **加载动画**: "正在切换到发布分支..."

#### 命令 2: 获取远程标签
```bash
git fetch --tags
```
- **目的**: 确保本地有所有远程标签的最新信息
- **加载动画**: "正在获取远程标签..."

#### 命令 3: 创建带注释的标签
```bash
git tag -a v1.1.2 -m 'Release-v1.1.2'
```
- **参数解析**:
  - `-a`: 创建带注释的标签
  - `v1.1.2`: 标签名称
  - `-m 'Release-v1.1.2'`: 标签消息
- **加载动画**: "正在创建版本标签..."

#### 命令 4: 推送标签到远程
```bash
git push origin v1.1.2
```
- **目的**: 将新标签推送到远程仓库
- **加载动画**: "正在推送标签..."

#### 命令 5: 创建 GitHub Release（可选）
```bash
gh release create v1.1.2 --generate-notes
```
- **条件**: 系统中安装了 GitHub CLI (`gh`)
- **功能**: 自动创建 GitHub Release 并生成变更日志
- **加载动画**: "正在创建 GitHub Release..."

## 常用参数含义

### `--type, -t`
- **类型**: `string`
- **可选值**: `major`, `minor`, `patch`
- **默认值**: `patch`
- **说明**: 指定版本递增类型
- **示例**:
  ```bash
  gfl tag --type major   # v1.1.1 → v2.0.0
  gfl tag --type minor   # v1.1.1 → v1.2.0
  gfl tag --type patch   # v1.1.1 → v1.1.2
  ```

## 版本递增规则

### 语义化版本控制 (SemVer)
格式: `MAJOR.MINOR.PATCH`

#### MAJOR (主版本)
- **条件**: 当进行不兼容的 API 修改时
- **示例**: `v1.2.3` → `v2.0.0`
- **影响**: 重置 MINOR 和 PATCH 为 0

#### MINOR (次版本)
- **条件**: 当向下兼容的功能性新增时
- **示例**: `v1.2.3` → `v1.3.0`
- **影响**: 重置 PATCH 为 0

#### PATCH (补丁版本)
- **条件**: 当向下兼容的问题修正时
- **示例**: `v1.2.3` → `v1.2.4`
- **影响**: 只递增 PATCH

## 使用场景

### 1. 补丁发布
```bash
# 修复 bug 后发布补丁版本
gfl tag --type patch
# v1.1.1 → v1.1.2
```

### 2. 功能发布
```bash
# 新功能完成后发布次版本
gfl tag --type minor
# v1.1.2 → v1.2.0
```

### 3. 重大版本发布
```bash
# 重大更新或重构后发布主版本
gfl tag --type major
# v1.2.0 → v2.0.0
```

### 4. 自动发布
```bash
# 使用默认补丁版本
gfl tag
# 等同于: gfl tag --type patch
```

## 注意事项

### 1. 前置条件
- 必须在 Git 仓库中执行
- 需要先创建对应的发布分支
- 需要网络连接推送标签
- 安装 GitHub CLI 以自动创建 Release

### 2. 分支要求
- 发布分支必须存在：`releases/release-{version}`
- 发布分支应该包含所有计划发布的代码
- 建议在发布前完成所有测试和审查

### 3. 版本号规范
- 遵循语义化版本控制 (SemVer)
- 标签名称必须以 `v` 开头
- 避免使用预发布版本号（如 `-alpha`, `-beta`）

### 4. 标签管理
- 不要删除已发布的标签
- 标签应该指向特定的提交，而不是分支
- 确保标签的唯一性

### 5. GitHub Release
- 需要 GitHub CLI 已认证
- 自动生成变更日志基于 PR 和提交信息
- 可以在 GitHub 网页上进一步编辑 Release 信息

## 工作流程集成

### 标准发布流程
```bash
# 1. 确保代码已合并到 develop
gfl sync

# 2. 创建发布分支
gfl release --type minor

# 3. 切换到发布分支进行最终测试
gfl checkout
# 选择 releases/release-v1.2.0

# 4. 运行测试
npm test

# 5. 创建标签和 Release
gfl tag --type minor

# 6. 发布到生产环境
# 部署流程...
```

### 热修复发布流程
```bash
# 1. 从生产分支创建热修复
gfl hotfix critical-bug-fix

# 2. 修复代码并测试
git add .
git commit -m "Fix critical security bug"

# 3. 发布热修复
gfl publish
gfl pr main  # 向 main 分支创建 PR

# 4. 合并后创建补丁标签
gfl tag --type patch
```

## 使用示例

### 基本使用
```bash
gfl tag
# 输出: 上一版本: v1.1.1
# 输出: 新版本: v1.1.2
# 然后执行一系列 Git 操作...
```

### 指定版本类型
```bash
# 主版本发布
gfl tag --type major
# v1.1.1 → v2.0.0

# 次版本发布
gfl tag --type minor
# v1.1.1 → v1.2.0

# 补丁版本发布
gfl tag --type patch
# v1.1.1 → v1.1.2
```

### 使用别名
```bash
gfl t --type minor
```

## 错误处理

### 常见错误及解决方案

1. **"发布分支不存在"**
   ```bash
   # 创建发布分支
   gfl release --type minor

   # 然后创建标签
   gfl tag
   ```

2. **"标签已存在"**
   ```bash
   # 检查现有标签
   git tag -l

   # 手动删除错误标签
   git tag -d v1.1.2
   git push origin :refs/tags/v1.1.2

   # 重新创建
   gfl tag
   ```

3. **"GitHub CLI 未安装"**
   ```bash
   # macOS
   brew install gh

   # Ubuntu
   sudo apt install gh

   # 认证
   gh auth login
   ```

4. **"权限不足"**
   - 确认有推送标签的权限
   - 检查 GitHub CLI 认证状态
   - 验证仓库访问权限

## 高级用法

### 1. 批量操作
```bash
# 创建发布分支并打标签
gfl release --type minor && gfl tag --type minor
```

### 2. 自定义标签信息
```bash
# 手动创建标签（更灵活）
git tag -a v1.2.0 -m "Release v1.2.0: Add user management feature"
git push origin v1.2.0
gh release create v1.2.0 --title "Version 1.2.0" --notes "New features and improvements"
```

### 3. 脚本化发布
```bash
#!/bin/bash
# 自动发布脚本
VERSION_TYPE=${1:-patch}

echo "Creating release branch..."
gfl release --type $VERSION_TYPE

echo "Creating tag..."
gfl tag --type $VERSION_TYPE

echo "Release completed successfully!"
```

## 最佳实践

### 1. 发布前检查
```bash
# 检查分支状态
git status
git log --oneline -10

# 运行测试
npm test
npm run build

# 检查版本历史
git tag -l | sort -V
```

### 2. 版本号规划
- 遵循语义化版本控制规范
- 在项目文档中说明版本策略
- 定期回顾版本发布历史

### 3. Release 信息管理
- 在 GitHub Release 中详细说明变更
- 包含升级指南和迁移说明
- 提供变更日志（CHANGELOG）

## 相关命令

- `git tag`: 原生标签命令
- `gh release`: GitHub CLI Release 命令
- `gfl release`: 创建发布分支
- `gfl sync`: 同步远程仓库
- `git checkout`: 切换分支

## 配置依赖

- 无特定配置依赖
- 依赖 Git 仓库的标签历史
- GitHub CLI 配置影响 Release 创建功能