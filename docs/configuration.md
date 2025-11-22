# GFL 配置指南

GFL 使用灵活的配置系统，支持全局和本地配置文件，以及环境变量。

## 配置文件

### 全局配置 (.gfl.config.yml)

全局配置文件用于团队共享的设置，应该提交到版本控制系统。

```yaml
# 开启调试模式
debug: false

# 开发基础分支
devBaseBranch: dev

# 生产分支
productionBranch: main

# 开发者昵称
nickname: aric

# 分支类型前缀
featurePrefix: feature
fixPrefix: fix
hotfixPrefix: hotfix
```

### 本地配置 (.gfl.config.local.yml)

本地配置文件用于个人设置，应该添加到 `.gitignore` 中。

```yaml
# 覆盖全局昵称设置
nickname: mynickname

# 自定义分支前缀
featurePrefix: feat
fixPrefix: bugfix
hotfixPrefix: hot
```

## 配置选项详解

### 基础配置

| 选项 | 类型 | 默认值 | 说明 |
|------|------|--------|------|
| `debug` | boolean | false | 是否启用调试模式 |
| `devBaseBranch` | string | dev | 开发基础分支名 |
| `productionBranch` | string | main | 生产分支名 |
| `nickname` | string | aric | 开发者昵称 |

### 分支前缀配置

| 选项 | 类型 | 默认值 | 说明 |
|------|------|--------|------|
| `featurePrefix` | string | feature | 功能分支前缀 |
| `fixPrefix` | string | fix | 修复分支前缀 |
| `hotfixPrefix` | string | hotfix | 热修复分支前缀 |

## 环境变量

GFL 支持通过环境变量覆盖配置：

### GFL_CONFIG_FILE

指定自定义配置文件路径：

```bash
export GFL_CONFIG_FILE=/path/to/custom-config.yml
gfl start new-feature
```

### 优先级

配置的优先级顺序（从高到低）：

1. 命令行参数
2. 环境变量
3. 本地配置文件 (.gfl.config.local.yml)
4. 全局配置文件 (.gfl.config.yml)
5. 默认值

## 配置示例

### 团队配置示例

```yaml
# .gfl.config.yml
debug: false
devBaseBranch: develop
productionBranch: master
nickname: team-member
featurePrefix: feature
fixPrefix: bugfix
hotfixPrefix: hotfix
```

### 个人配置示例

```yaml
# .gfl.config.local.yml
nickname: john-doe
debug: true
```

### 自定义分支命名

```yaml
# .gfl.config.local.yml
featurePrefix: feat
fixPrefix: fix
hotfixPrefix: emergency
```

这将产生以下分支命名：

- 功能分支：`feat/john-doe/user-login`
- 修复分支：`fix/john-doe/display-issue`
- 热修复分支：`emergency/john-doe/critical-bug`

## 配置管理

### 初始化配置

```bash
# 初始化默认配置
gfl init

# 强制覆盖现有配置
gfl init --force

# 初始化并设置昵称
gfl init --nickname myname
```

### 验证配置

```bash
# 查看当前配置
gfl config show

# 验证配置文件
gfl config validate
```

### 重置配置

```bash
# 重置为默认配置
gfl config reset

# 重置特定选项
gfl config reset --option nickname
```

## 高级配置

### 多项目配置

对于多个项目使用不同配置：

```bash
# 项目 A
cd project-a
gfl init --nickname team-a-dev

# 项目 B
cd project-b
gfl init --nickname team-b-dev
```

### 分支命名策略

#### 策略 1：包含团队信息

```yaml
nickname: frontend-team
featurePrefix: feature/fe
```

结果：`feature/fe/frontend-team/dashboard-ui`

#### 策略 2：JIRA 集成

```yaml
nickname: ""
featurePrefix: feature
```

结果：`feature/PROJ-123-user-profile`

#### 策略 3：简化命名

```yaml
nickname: ""
featurePrefix: ""
```

结果：`user-profile`

## 调试配置

启用调试模式查看详细执行信息：

```yaml
debug: true
```

调试输出示例：

```bash
$ gfl start new-feature
[DEBUG] Loading config from .gfl.config.yml
[DEBUG] Loading local config from .gfl.config.local.yml
[DEBUG] Current branch: main
[DEBUG] Dev base branch: dev
[DEBUG] Creating branch: feature/aric/new-feature
✓ 已从 dev 创建分支 feature/aric/new-feature
```

## 配置最佳实践

1. **团队配置**：将 `.gfl.config.yml` 提交到版本控制
2. **个人配置**：将 `.gfl.config.local.yml` 添加到 `.gitignore`
3. **一致性**：团队使用相同的分支命名规范
4. **文档化**：在项目 README 中说明配置规范
5. **备份**：定期备份重要配置文件

## 故障排除

### 常见配置问题

1. **配置文件未找到**
   ```bash
   # 重新初始化配置
   gfl init --force
   ```

2. **配置权限问题**
   ```bash
   # 检查文件权限
   ls -la .gfl.config*.yml

   # 修正权限
   chmod 644 .gfl.config.yml
   chmod 600 .gfl.config.local.yml
   ```

3. **配置格式错误**
   ```bash
   # 验证 YAML 格式
   python -c "import yaml; yaml.safe_load(open('.gfl.config.yml'))"
   ```