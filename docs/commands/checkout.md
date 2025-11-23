# GFL Checkout 命令技术文档

## 概述

`gfl checkout` 命令提供交互式的 Git 分支切换功能，支持别名 `co`。它允许用户通过选择列表快速切换本地分支。

## 实现原理

### 1. 命令定义

```go
var checkoutCmd = &cobra.Command{
    Use:     "checkout",
    Aliases: []string{"co"},
    Short:   "Interactive git branch switching (alias: co)",
    Run: func(cmd *cobra.Command, args []string) {
        branches := utils.GetLocalBranches()
        utils.BuildCommandList(branches)
    },
}
```

### 2. 执行流程

#### 步骤 1: 获取本地分支列表
调用 `utils.GetLocalBranches()` 获取所有本地分支：
```go
// 伪代码示例
func GetLocalBranches() []string {
    // 执行: git branch --format='%(refname:short)'
    // 解析输出并返回分支名称列表
}
```

#### 步骤 2: 构建交互式选择列表
调用 `utils.BuildCommandList(branches)` 创建交互式界面：
```go
// 伪代码示例
func BuildCommandList(branches []string) {
    // 1. 显示分支列表
    // 2. 提供选择界面
    // 3. 处理用户选择
    // 4. 执行 git checkout 命令
}
```

### 3. Shell 命令执行过程

#### 命令 1: 获取分支列表
```bash
git branch --format='%(refname:short)'
```
- **目的**: 获取本地分支的简洁名称列表
- **输出**: 每行一个分支名称
- **示例输出**:
  ```
  main
  develop
  feature/aric/user-auth
  fix/aric/login-bug
  ```

#### 命令 2: 切换分支（用户选择后）
```bash
git checkout <selected-branch>
```
- **目的**: 切换到用户选择的分支
- **触发条件**: 用户在交互界面中选择分支后
- **参数**: `<selected-branch>` 为用户选择的分支名称

## 交互式界面特性

### 1. 分支列表显示
- 按字母顺序排列
- 当前分支通常有特殊标记（如 * 或颜色）
- 支持彩色输出增强可读性

### 2. 选择机制
- 支持键盘导航（上下箭头）
- 支持数字快速选择
- Enter 确认选择
- Esc 或 Ctrl+C 取消操作

### 3. 视觉反馈
- 选中项高亮显示
- 操作结果实时反馈
- 错误信息友好提示

## 常用参数含义

此命令不接受任何参数，完全通过交互式界面操作。

## 使用场景

### 1. 多分支项目管理
当项目中有大量功能分支时，使用交互式选择比记忆分支名称更高效：
```bash
gfl co
# 显示所有可用分支供选择
```

### 2. 快速分支切换
避免输入完整的分支名称：
```bash
# 传统方式
git checkout feature/aric/user-authentication-system

# GFL 方式
gfl co  # 然后从列表选择
```

### 3. 分支探索
在不确定所有分支名称时，可以快速查看可用分支：
```bash
gfl co  # 查看所有分支，然后取消选择
```

## 注意事项

### 1. 前置条件
- 必须在 Git 仓库中执行
- 需要有本地分支存在
- 终端需要支持交互式操作

### 2. 工作目录状态
- 如果有未提交的更改，切换分支可能失败
- 建议在切换前提交或暂存更改

### 3. 分支状态
- 只能切换到本地存在的分支
- 如果需要切换远程分支，先使用 `gfl sync` 或 `git fetch`

### 4. 交互依赖
- 不支持脚本化或自动化使用
- 需要用户交互输入

## 扩展功能

### 集成建议
虽然当前实现较简单，但可以考虑以下增强：

1. **分支过滤**
   ```bash
   gfl co --filter "feature/*"
   gfl co --exclude "main"
   ```

2. **远程分支支持**
   ```bash
   gfl co --remote  # 显示并可选择远程分支
   ```

3. **最近使用分支**
   ```bash
   gfl co --recent  # 显示最近使用的分支
   ```

## 使用示例

### 基本使用
```bash
gfl checkout
# 或使用别名
gfl co
```

### 典型工作流
```bash
# 1. 查看并切换到功能分支
gfl co
# [从列表中选择 feature/aric/user-auth]

# 2. 进行开发工作...

# 3. 切换到另一个分支
gfl co
# [从列表中选择 fix/aric/login-bug]
```

## 相关命令

- `gfl start`: 创建新分支
- `gfl publish`: 发布当前分支
- `gfl sweep`: 清理分支
- `git branch`: 原生 Git 分支命令

## 故障排除

### 常见问题

1. **"没有本地分支"**
   ```bash
   # 检查远程分支
   git branch -r
   # 同步远程分支
   gfl sync
   ```

2. **"无法切换分支（有未提交更改）"**
   ```bash
   # 提交更改
   git add .
   git commit -m "Save work"

   # 或暂存更改
   git stash
   gfl co
   git stash pop
   ```

3. **"交互界面无法显示"**
   - 确认终端支持交互式操作
   - 检查终端环境变量（TERM）
   - 尝试使用标准 git checkout 命令

### 替代方案
当交互式界面不可用时，可以使用原生 Git 命令：
```bash
# 查看所有分支
git branch

# 切换分支
git checkout <branch-name>
```

## 性能考虑

### 1. 大量分支时的性能
- 分支数量很多时，列表可能很长
- 可以考虑分页或搜索功能

### 2. 网络依赖
- 不需要网络连接（只操作本地分支）
- 响应速度快于远程分支操作

### 3. 内存使用
- 内存占用较小，主要存储分支名称列表
- 适合在资源受限的环境中使用