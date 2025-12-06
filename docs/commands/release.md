# GFL Release 命令技术文档

## 概述

`gfl release` 命令用于基于最新标签生成新的发布版本，创建发布分支并推送到远程仓库，支持别名 `rls`。

## 实现原理

### 1. 命令定义

```go
var releaseCmd = &cobra.Command{
    Use:     "release",
    Aliases: []string{"rls"},
    Short:   "Generate new release version based on latest tag",
    Run: func(cmd *cobra.Command, args []string) {
        // 实现逻辑
    },
}
```

### 2. 执行流程

#### 步骤 1: 版本计算和参数解析
```go
version := utils.GetLatestVersion()
versionType, _ := cmd.Flags().GetString("type")
hotfix, _ := cmd.Flags().GetBool("hotfix")
newVersion, err := utils.IncrementVersion(version, versionType)
```

#### 步骤 2: 版本信息显示
```go
utils.Infof(strings.GetPath("release.previous_version"), version)
utils.Successf(strings.GetPath("release.new_version"), newVersion)
```

#### 步骤 3: 基础分支选择
```go
config := utils.ReadConfig()
remoteBranch := config.DevBaseBranch
if hotfix {
    remoteBranch = config.ProductionBranch
}
```

#### 步骤 4: Git 操作执行

### 3. Shell 命令执行过程

#### 命令 1: 同步远程仓库
```bash
git fetch origin
```
- **目的**: 获取远程仓库的最新状态
- **加载动画**: 显示 "正在获取远程分支..."
- **错误处理**: 失败时中止执行并显示错误

#### 命令 2: 创建发布分支
```bash
git checkout -b releases/release-v1.2.0 origin/develop
```
- **参数解析**:
  - `-b`: 创建新分支
  - `releases/release-v1.2.0`: 生成的发布分支名称
  - `origin/develop`: 基础分支（默认为开发分支）
- **加载动画**: 显示 "正在创建发布分支..."

#### 命令 3: 推送发布分支
```bash
git push -u origin releases/release-v1.2.0
```
- **目的**: 将新创建的发布分支推送到远程仓库
- **参数解析**:
  - `-u`: 设置上游分支跟踪
  - `origin`: 远程仓库名称
  - `releases/release-v1.2.0`: 发布分支名称
- **加载动画**: 显示 "正在推送发布分支..."

## 常用参数含义

### `--type, -t`
- **类型**: `string`
- **可选值**: `major`, `minor`, `patch`
- **默认值**: `patch`
- **说明**: 指定版本递增类型
- **示例**:
  ```bash
  gfl release --type major    # v1.1.1 → v2.0.0
  gfl release --type minor    # v1.1.1 → v1.2.0
  gfl release --type patch    # v1.1.1 → v1.1.2
  ```

### `--hotfix, -x`
- **类型**: `bool`
- **默认值**: `false`
- **说明**: 基于生产分支创建发布分支（用于热修复发布）
- **效果**: 将基础分支从 `develop` 改为 `main`
- **使用场景**: 在生产版本出现问题时创建热修复发布
- **示例**:
  ```bash
  gfl release --hotfix        # 基于 main 分支创建
  gfl release --type patch --hotfix  # 补丁版本 + 基于生产分支
  ```

## 版本递增规则

### 语义化版本控制 (SemVer)
格式: `MAJOR.MINOR.PATCH`

#### MAJOR (主版本) - `--type major`
- **条件**: 不兼容的 API 修改
- **示例**: `v1.2.3` → `v2.0.0`
- **影响**: 重置 MINOR 和 PATCH 为 0
- **分支名称**: `releases/release-v2.0.0`

#### MINOR (次版本) - `--type minor`
- **条件**: 向下兼容的功能性新增
- **示例**: `v1.2.3` → `v1.3.0`
- **影响**: 重置 PATCH 为 0
- **分支名称**: `releases/release-v1.3.0`

#### PATCH (补丁版本) - `--type patch`
- **条件**: 向下兼容的问题修正
- **示例**: `v1.2.3` → `v1.2.4`
- **影响**: 只递增 PATCH
- **分支名称**: `releases/release-v1.2.4`

## 使用场景

### 1. 功能发布
```bash
# 新功能完成后发布次版本
gfl release --type minor
# 创建: releases/release-v1.2.0
```

### 2. 重大版本发布
```bash
# 重大更新后发布主版本
gfl release --type major
# 创建: releases/release-v2.0.0
```

### 3. 热修复发布
```bash
# 热修复完成后发布补丁版本
gfl release --type patch --hotfix
# 基于 main 分支创建: releases/release-v1.1.4
```

### 4. 自动补丁发布
```bash
# 使用默认补丁类型
gfl release
# 等同于: gfl release --type patch
# 创建: releases/release-v1.1.4
```

### 5. 使用别名
```bash
gfl rls --type minor
gfl rls --hotfix
```

## 注意事项

### 1. 前置条件
- 必须在 Git 仓库中执行
- 必须先运行 `gfl init` 初始化配置
- 基础分支（develop 或 main）必须存在
- 需要网络连接推送发布分支

### 2. 分支策略
- **常规发布**: 基于 `develop` 分支
- **热修复发布**: 基于 `main` 分支（使用 `--hotfix`）
- 发布分支名称格式：`releases/release-{version}`

### 3. 发布流程
- 此命令只创建发布分支，不创建标签
- 完整发布流程需要后续使用 `gfl tag` 命令
- 发布分支用于最终测试和集成

### 4. 版本规划
- 遵循语义化版本控制规范
- 避免频繁的主版本发布
- 补丁版本应包含重要修复

### 5. 团队协作
- 发布前应完成代码审查和测试
- 及时通知团队发布计划
- 协调部署时间和回滚计划

## 发布工作流程

### 标准发布流程
```bash
# 1. 确保开发分支稳定
gfl sync
git checkout develop
git pull origin develop

# 2. 运行测试套件
npm test
npm run build

# 3. 创建发布分支
gfl release --type minor
# 输出: 上一版本: v1.1.1
# 输出: 新版本: v1.2.0
# 输出: 正在获取远程分支...
# 输出: 正在创建发布分支...
# 输出: 正在推送发布分支...

# 4. 切换到发布分支进行最终测试
gfl checkout
# 选择: releases/release-v1.2.0

# 5. 在发布分支上进行最终测试和修复
# （如果需要）
git add .
git commit -m "Final adjustments for v1.2.0 release"
git push origin releases/release-v1.2.0

# 6. 创建版本标签
gfl tag --type minor

# 7. 部署到生产环境
# 部署流程...

# 8. 合并发布分支到生产分支
git checkout main
git merge releases/release-v1.2.0
git push origin main

# 9. 合并发布分支回开发分支
git checkout develop
git merge releases/release-v1.2.0
git push origin develop

# 10. 清理发布分支（可选）
gfl sweep releases/release-v1.2.0 --local
```

### 热修复发布流程
```bash
# 1. 创建热修复分支
gfl hotfix critical-bug-fix

# 2. 修复问题
git add .
git commit -m "Fix critical bug in authentication"
gfl publish

# 3. 创建热修复发布分支
gfl release --type patch --hotfix
# 基于 main 分支创建: releases/release-v1.1.4

# 4. 创建补丁标签
gfl tag --type patch

# 5. 部署热修复
# 部署流程...

# 6. 将修复合并回开发分支
git checkout develop
git cherry-pick <hotfix-commit-hash>
git push origin develop
```

## 使用示例

### 基本使用
```bash
gfl release
# 输出: 上一版本: v1.1.1
# 输出: 新版本: v1.1.2
# 然后执行 Git 操作创建发布分支
```

### 指定版本类型
```bash
# 主版本发布
gfl release --type major
# v1.1.1 → v2.0.0
# 创建: releases/release-v2.0.0

# 次版本发布
gfl release --type minor
# v1.1.1 → v1.2.0
# 创建: releases/release-v1.2.0

# 补丁版本发布
gfl release --type patch
# v1.1.1 → v1.1.2
# 创建: releases/release-v1.1.2
```

### 热修复发布
```bash
gfl release --hotfix
# 基于 main 分支创建发布分支
# 输出: 上一版本: v1.1.1
# 输出: 新版本: v1.1.2
# 创建: releases/release-v1.1.2
```

### 组合使用
```bash
# 热修复补丁发布
gfl release --type patch --hotfix
# 基于 main 分支创建补丁发布分支

# 使用别名
gfl rls --type minor
```

## 错误处理

### 常见错误及解决方案

1. **"版本号解析失败"**
   ```bash
   # 检查现有标签
   git tag -l | sort -V

   # 手动设置初始标签
   git tag v1.0.0
   git push origin v1.0.0
   ```

2. **"基础分支不存在"**
   ```bash
   # 检查开发分支
   git branch -r | grep develop

   # 同步远程分支
   gfl sync

   # 创建开发分支（如果不存在）
   git checkout -b develop origin/main
   git push origin develop
   ```

3. **"发布分支已存在"**
   ```bash
   # 切换到现有发布分支
   git checkout releases/release-v1.2.0

   # 或删除现有分支重新创建
   git branch -D releases/release-v1.2.0
   gfl release --type minor
   ```

4. **"权限不足"**
   ```bash
   # 检查仓库权限
   git remote show origin

   # 配置认证
   ssh -T git@github.com
   ```

5. **"网络连接失败"**
   - 检查网络连接
   - 确认远程仓库 URL 正确
   - 检查防火墙和代理设置

## 最佳实践

### 1. 发布前准备
```bash
# 检查代码状态
git status
git log --oneline origin/develop..HEAD

# 运行完整测试
npm test
npm run lint
npm run build

# 检查版本历史
git tag -l | sort -V | tail -5
```

### 2. 版本号策略
- 遵循语义化版本控制
- 在文档中说明版本策略
- 定期回顾发布历史
- 避免跳跃式版本增长

### 3. 发布分支管理
- 保持发布分支的简洁性
- 避免在发布分支中添加新功能
- 及时清理已合并的发布分支
- 使用有意义的分支名称

### 4. 团队协作
- 制定发布时间表
- 建立发布审查流程
- 准备回滚计划
- 及时沟通发布状态

## 高级用法

### 1. 批量操作
```bash
# 创建发布分支并打标签
gfl release --type minor && gfl tag --type minor
```

### 2. 脚本化发布
```bash
#!/bin/bash
# 自动发布脚本
RELEASE_TYPE=${1:-minor}

echo "准备发布 $RELEASE_TYPE 版本..."

# 检查工作目录状态
if [ -n "$(git status --porcelain)" ]; then
    echo "错误: 工作目录不干净，请先提交所有更改"
    exit 1
fi

# 运行测试
npm test
if [ $? -ne 0 ]; then
    echo "错误: 测试失败"
    exit 1
fi

# 创建发布分支
gfl release --type $RELEASE_TYPE

echo "发布分支创建完成，请进行最终测试和部署"
```

### 3. 与 CI/CD 集成
```yaml
# GitHub Actions 示例
name: Create Release
on:
  workflow_dispatch:
    inputs:
      version_type:
        required: true
        type: choice
        options: [patch, minor, major]

jobs:
  release:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - name: Setup GFL
        run: |
          wget -qO- https://github.com/user/gfl/releases/latest/download/gfl-linux-amd64.tar.gz | tar xz
          sudo mv gfl /usr/local/bin/

      - name: Create Release Branch
        run: gfl release --type ${{ github.event.inputs.version_type }}
```

## 相关命令

- `gfl tag`: 创建版本标签
- `gfl hotfix`: 创建热修复分支
- `gfl sync`: 同步远程仓库
- `gfl checkout`: 切换分支
- `gfl sweep`: 清理分支
- `git tag`: 原生标签命令

## 配置依赖

以下配置项影响 release 命令：
- `DevBaseBranch`: 常规发布的基础分支（通常是 "develop"）
- `ProductionBranch`: 热修复发布的基础分支（通常是 "main"）