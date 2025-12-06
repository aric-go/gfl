# gfl bugfix

开始一个bug修复分支，专门用于修复代码中的问题。

## 用法

```bash
gfl bugfix [bug-name]
gfl fix [bug-name]     # 别名，更简洁的写法
```

## 参数

- `bug-name`: bug修复的描述名称（必需）

## 功能

该命令会：
1. 同步远程仓库
2. 基于配置的开发基础分支（默认为 `develop`）创建新的修复分支
3. 分支命名格式为 `fix/[nickname]/[bug-name]` 或 `fix/[bug-name]`（取决于是否配置了昵称）

## 示例

### 创建基本的bug修复分支
```bash
gfl bugfix memory-leak
# 或者使用别名
gfl fix memory-leak
```
创建分支: `fix/memory-leak`

### 带昵称的bug修复分支
```bash
gfl bugfix login-validation-error
# 或者使用别名
gfl fix login-validation-error
```
创建分支: `fix/aric/login-validation-error`（如果昵称配置为 "aric"）

### 复杂的bug修复分支名称
```bash
gfl bugfix user-profile-image-upload-failure
# 或者使用别名
gfl fix user-profile-image-upload-failure
```
创建分支: `fix/aric/user-profile-image-upload-failure`

## 分支命名规则

- **有昵称**: `fix/{nickname}/{bug-name}`
- **无昵称**: `fix/{bug-name}`

例如：
- `fix/aric/memory-leak` - 有昵称的情况
- `fix/memory-leak` - 无昵称的情况

## 相关命令

- `gfl publish`: 发布当前分支
- `gfl pr`: 创建 Pull Request
- `gfl checkout`: 切换分支
- `gfl start`: 开始新功能分支（也可以使用 `gfl start fix:xxx`）
- `gfl hotfix`: 开始热修复分支

## 与其他命令的区别

| 命令 | 用途 | 分支前缀 |
|------|------|---------|
| `gfl start` | 新功能开发 | `feature` |
| `gfl bugfix` | bug修复 | `fix` |
| `gfl hotfix` | 紧急热修复 | `hotfix` |

## 配置依赖

以下配置项影响 bugfix 命令的行为：
- `devBaseBranch`: 基础开发分支（默认: `develop`）
- `nickname`: 开发者昵称（可选）
- `fixPrefix`: 修复分支前缀（默认: `fix`）

## 工作流程

1. **识别bug**: 确定需要修复的问题
2. **创建分支**: 使用 `gfl bugfix [description]` 创建修复分支
3. **修复问题**: 在新分支中进行代码修复
4. **测试验证**: 确保修复正确且没有引入新问题
5. **发布分支**: 使用 `gfl publish` 推送到远程仓库
6. **创建PR**: 使用 `gfl pr` 创建代码审查请求
7. **代码合并**: 审查通过后合并到目标分支

## 最佳实践

### 分支命名建议
- 使用描述性名称：`memory-leak` 而不是 `bug1`
- 使用小写字母和连字符：`user-authentication-fix`
- 避免特殊字符和空格

### 与 `gfl start fix:xxx` 的选择
- **gfl bugfix**: 专门用于bug修复，语义更清晰
- **gfl start fix:xxx**: 更通用的语法，可以处理所有分支类型

## 示例工作流

```bash
# 创建bug修复分支（两种方式都可以）
gfl bugfix user-login-validation
# 或者使用别名
gfl fix user-login-validation

# 进行代码修复...
# git add .
# git commit -m "Fix user login validation issue"

# 发布并创建PR
gfl publish
gfl pr
```