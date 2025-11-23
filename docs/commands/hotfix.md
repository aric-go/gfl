# GFL Hotfix 命令技术文档

## 概述

`gfl hotfix` 命令用于创建热修复分支，基于生产分支快速修复紧急问题，支持别名 `hf`。

## 实现原理

### 1. 命令定义

```go
var hotfixCmd = &cobra.Command{
    Use:     "hotfix [hotfix-name]",
    Aliases: []string{"hf"},
    Short:   "Start a hotfix branch",
    Args:    cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        // 实现逻辑
    },
}
```

### 2. 执行流程

#### 步骤 1: 配置读取和参数解析
```go
config := utils.ReadConfig()
featureName := args[0]  // 从参数中获取Hotfix名称
branchName := utils.GenerateBranchName(config, "hotfix", featureName)
```

#### 步骤 2: Git 操作执行

### 3. Shell 命令执行过程

#### 命令 1: 同步远程仓库
```bash
git fetch origin
```
- **目的**: 更新远程分支引用，确保获取最新的生产分支状态
- **加载动画**: 显示 "正在同步..."
- **错误处理**: 失败时中止执行

#### 命令 2: 创建热修复分支
```bash
git checkout -b hotfix/aric/critical-security-fix origin/main
```
- **参数解析**:
  - `-b`: 创建新分支
  - `hotfix/aric/critical-security-fix`: 生成的热修复分支名称
  - `origin/main`: 基础分支（生产分支）
- **配置影响**: `config.ProductionBranch` 决定基础分支
- **加载动画**: 显示 "正在创建热修复分支..."

## 常用参数含义

### 位置参数: [hotfix-name]
- **类型**: `string`
- **必填**: 是
- **说明**: 热修复分支的名称
- **格式**: 简洁描述要修复的问题
- **示例**:
  ```bash
  gfl hotfix memory-leak
  gfl hotfix security-vulnerability
  gfl hotfix login-bug
  gfl hotfix database-connection
  ```

## 分支命名规范

### 命名模式
```
hotfix/{nickname}/{hotfix-name}
```

### 示例
- `hotfix/aric/critical-security-fix`
- `hotfix/aric/login-page-crash`
- `hotfix/aric/payment-failure`
- `hotfix/aric/data-corruption-bug`

### 配置项影响
- `nickname`: 开发者昵称
- `hotfixPrefix`: 热修复分支前缀（默认 "hotfix"）
- `ProductionBranch`: 基础分支（通常是 "main"）

## 使用场景

### 1. 生产环境紧急问题
```bash
# 发现生产环境安全漏洞
gfl hotfix security-vulnerability

# 快速修复并发布
git add .
git commit -m "Fix critical security vulnerability"
gfl publish
gfl pr main  # 向生产分支创建 PR
```

### 2. 数据相关问题
```bash
# 发现数据损坏问题
gfl hotfix data-corruption

# 修复数据问题
# ... 修复代码 ...
git commit -m "Fix data corruption issue"
gfl publish
```

### 3. 性能问题
```bash
# 发现严重性能问题
gfl hotfix performance-degradation

# 优化性能
git commit -m "Optimize database queries to fix performance issue"
gfl publish
```

### 4. 使用别名
```bash
gfl hf critical-bug
gfl hf ui-crash
```

## 注意事项

### 1. 前置条件
- 必须在 Git 仓库中执行
- 必须先运行 `gfl init` 初始化配置
- 生产分支必须存在且可访问
- 需要网络连接同步远程仓库

### 2. 分支策略
- 基于生产分支（通常是 `main`）创建
- 不是基于开发分支，确保修复的稳定性
- 热修复分支需要快速合并回生产分支

### 3. 工作流程特殊性
- 热修复应该小而专注
- 避免在热修复分支中添加新功能
- 快速测试和发布
- 及时将修复合并回开发分支

### 4. 版本管理
- 热修复通常会创建补丁版本
- 使用 `gfl tag --type patch` 创建版本标签
- 可能需要创建新的发布分支

### 5. 团队协作
- 热修复应该由最有经验的开发者处理
- 及时通知团队成员热修复的进行
- 协调发布时间以避免冲突

## 热修复工作流程

### 标准热修复流程
```bash
# 1. 创建热修复分支
gfl hotfix critical-security-fix

# 2. 快速修复问题
vim security.patch
git add security.patch
git commit -m "Fix critical security vulnerability in authentication"

# 3. 测试修复（在测试环境）
# 部署到测试环境并验证

# 4. 发布热修复
gfl publish

# 5. 创建到生产分支的 PR
gfl pr main

# 6. 代码审查和紧急合并
# 团队负责人快速审查并合并

# 7. 创建补丁版本
gfl tag --type patch

# 8. 将修复合并回开发分支
git checkout develop
git cherry-pick <hotfix-commit-hash>
git push origin develop

# 9. 清理热修复分支
gfl sweep critical-security-fix --local
```

### 紧急发布流程
```bash
# 1. 立即创建热修复分支
gfl hotfix emergency-fix

# 2. 最小化修复
git commit -m "Emergency fix: prevent system crash"

# 3. 直接发布到生产
gfl publish
git push origin main --force  # 紧急情况下使用

# 4. 事后补流程
gfl tag --type patch
gfl pr main  # 记录修复过程
```

## 使用示例

### 基本使用
```bash
gfl hotfix login-page-crash
# 创建分支: hotfix/aric/login-page-crash
# 输出: 热修复分支 hotfix/aric/login-page-crash 创建成功
```

### 使用别名
```bash
gfl hf memory-leak
# 创建分支: hotfix/aric/memory-leak
```

### 完整热修复流程
```bash
# 1. 发现生产问题
# 用户报告登录页面崩溃

# 2. 创建热修复分支
gfl hotfix login-page-crash
# 输出: 正在同步...
# 输出: 正在创建热修复分支...
# 输出: 执行命令: git checkout -b hotfix/aric/login-page-crash origin/main
# 输出: 热修复分支 hotfix/aric/login-page-crash 创建成功

# 3. 修复代码
vim login.js
git add login.js
git commit -m "Fix login page crash on invalid input"

# 4. 快速测试
npm test

# 5. 发布修复
gfl publish
# 输出: 正在推送...
# 输出: 分支发布成功

# 6. 创建紧急 PR
gfl pr main

# 7. 合并后创建补丁版本
gfl tag --type patch
# v1.2.3 → v1.2.4

# 8. 清理分支
gfl sweep login-page-crash --local
```

## 错误处理

### 常见错误及解决方案

1. **"生产分支不存在"**
   ```bash
   # 检查生产分支配置
   gfl config

   # 同步远程分支
   gfl sync

   # 检查分支是否存在
   git branch -r | grep main
   ```

2. **"权限不足"**
   ```bash
   # 检查仓库权限
   git remote show origin

   # 配置认证
   ssh -T git@github.com
   ```

3. **"分支已存在"**
   ```bash
   # 切换到现有分支
   git checkout hotfix/aric/existing-fix

   # 或使用不同的名称
   gfl hotfix login-crash-v2
   ```

4. **"工作目录不干净"**
   ```bash
   # 提交当前更改
   git add .
   git commit -m "Save current work"

   # 或暂存更改
   git stash
   gfl hotfix new-fix
   git stash pop
   ```

## 最佳实践

### 1. 热修复原则
- **快速**: 最小化修复时间
- **专注**: 只修复一个具体问题
- **安全**: 不引入新的风险
- **测试**: 在测试环境充分验证

### 2. 代码质量
- 保持代码简洁明了
- 添加必要的注释说明修复原因
- 包含回归测试（如果可能）

### 3. 文档记录
- 在 PR 中详细描述问题和修复
- 更新相关文档
- 记录在变更日志中

### 4. 团队沟通
- 及时通知所有相关人员
- 在团队聊天中记录修复过程
- 安排事后回顾会议

## 相关命令

- `gfl start`: 创建功能分支
- `gfl publish`: 发布分支
- `gfl pr`: 创建 Pull Request
- `gfl tag`: 创建版本标签
- `glf sweep`: 清理分支
- `gfl checkout`: 切换分支

## 配置依赖

以下配置项影响 hotfix 命令：
- `ProductionBranch`: 热修复的基础分支
- `nickname`: 开发者昵称
- `hotfixPrefix`: 热修复分支前缀

## 与其他命令的协作

### 与 Tag 命令协作
```bash
# 热修复发布后创建补丁版本
gfl hotfix critical-bug
# ... 修复代码 ...
gfl publish
gfl tag --type patch
```

### 与 Release 命令协作
```bash
# 热修复后可能需要创建新的发布分支
gfl hotfix production-bug
# ... 修复并发布 ...
gfl release --type hotfix
```

### 与 Sweep 命令协作
```bash
# 清理已修复的热修复分支
gfl sweep critical-fix --local --confirm
```