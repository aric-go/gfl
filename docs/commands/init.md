# GFL Init 命令技术文档

## 概述

`gfl init` 命令用于初始化 GitHub Flow 配置，创建配置文件并设置开发环境的基本参数。

## 实现原理

### 1. 命令定义和参数

```go
var initCmd = &cobra.Command{
    Use:   "init",
    Short: "Initialize Github Flow configuration",
    Run: func(cmd *cobra.Command, args []string) {
        // 实现逻辑
    },
}
```

### 2. 执行流程

#### 步骤 1: 参数解析
- **`--force/-f`**: 强制覆盖现有配置文件
- **`--nickname/-n`**: 设置 GitHub Flow 昵称（可选）

#### 步骤 2: 配置文件加载
程序通过 `embed` 包内嵌两个默认配置文件模板：
- `assets/.gfl.config.yml` - 全局配置文件模板
- `assets/.gfl.config.local.yml` - 本地配置文件模板

#### 步骤 3: YAML 解析和配置合并
```go
var gflConfigYaml utils.YamlConfig
var gflLocalConfigYaml utils.YamlConfig

_ = yaml.Unmarshal(gflConfig, &gflConfigYaml)
_ = yaml.Unmarshal(gflLocalConfig, &gflLocalConfigYaml)

if nickname != "" {
    gflLocalConfigYaml.Nickname = nickname
}
```

#### 步骤 4: 配置文件创建
1. 创建 `.gfl.config.yml`（全局配置，不会被强制覆盖）
2. 创建 `.gfl.config.local.yml`（本地配置，根据 force 标志决定是否覆盖）
3. 本地配置文件会自动添加到 `.gitignore`

### 3. Shell 命令执行过程

此命令不直接执行 shell 命令，而是通过 Go 的文件操作创建配置文件：

1. **读取内嵌资源**: `assets.ReadFile()`
2. **YAML 解析**: `yaml.Unmarshal()`
3. **文件创建**: `utils.CreateGflConfig()`
4. **Git 忽略**: 自动添加本地配置文件到 `.gitignore`

## 常用参数含义

### `--force, -f`
- **类型**: `bool`
- **默认值**: `false`
- **说明**: 强制覆盖已存在的本地配置文件
- **使用场景**: 重新初始化配置时使用

### `--nickname, -n`
- **类型**: `string`
- **默认值**: `""`
- **说明**: 设置开发者的 GitHub Flow 昵称
- **使用场景**: 在分支命名中标识开发者
- **示例**: `gfl init --nickname aric`

## 注意事项

### 1. 配置文件优先级
配置文件按以下优先级加载（高到低）：
1. 自定义配置文件（`GFL_CONFIG_FILE` 环境变量指定）
2. 本地配置文件（`.gfl.config.local.yml`）
3. 全局配置文件（`.gfl.config.yml`）
4. 默认配置

### 2. 安全考虑
- 本地配置文件会自动添加到 `.gitignore`，避免敏感信息泄露
- 全局配置文件不会被强制覆盖，保护现有配置

### 3. 文件权限
创建的配置文件遵循系统的默认 umask 设置

### 4. 错误处理
- YAML 解析失败时会显示错误信息
- 文件创建失败时会中止初始化流程
- 配置文件已存在且未使用 `--force` 时会报错

## 配置文件结构

### 全局配置文件 (.gfl.config.yml)
```yaml
debug: false
devBaseBranch: "develop"
productionBranch: "main"
nickname: ""
featurePrefix: "feature"
fixPrefix: "fix"
hotfixPrefix: "hotfix"
```

### 本地配置文件 (.gfl.config.local.yml)
```yaml
nickname: ""  # 开发者特定昵称
# 其他配置通常为空，继承自全局配置
```

## 使用示例

### 基本初始化
```bash
gfl init
```

### 带昵称的初始化
```bash
gfl init --nickname aric
```

### 强制重新初始化
```bash
gfl init --force
```

### 完整参数示例
```bash
gfl init --nickname aric --force
```

## 相关文件

- **配置文件位置**: 项目根目录
- **配置加载逻辑**: `utils/config.go`
- **默认配置模板**: `cmd/assets/`
- **字符串资源**: `utils/strings.yml`