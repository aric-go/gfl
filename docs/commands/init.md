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

### 4. 完整 Shell 命令展示原理

以下是 `gfl init` 命令的完整执行过程和等效的 shell 命令：

#### 步骤 0: 初始化检查
```bash
# GFL 内部执行的检查命令
# 检查是否在 Git 仓库中
git rev-parse --is-inside-work-tree

# 检查现有配置文件
ls -la .gfl.config*
```

#### 步骤 1: 读取内嵌配置模板
```bash
# GFL 使用 Go embed 包读取，等效的 shell 命令：
# 读取全局配置模板
cat > /tmp/global-config.yml << 'EOF'
debug: false
devBaseBranch: "develop"
productionBranch: "main"
nickname: ""
featurePrefix: "feature"
fixPrefix: "fix"
hotfixPrefix: "hotfix"
EOF

# 读取本地配置模板
cat > /tmp/local-config.yml << 'EOF'
nickname: ""
debug: false
EOF
```

#### 步骤 2: 处理用户参数
假设用户执行：`gfl init --nickname aric`

```bash
# GFL 内部处理昵称参数，等效的 shell 命令：
# 更新本地配置模板中的昵称
sed "s/nickname: \"\"/nickname: \"aric\"/" /tmp/local-config.yml > /tmp/local-config-with-nickname.yml

# 验证 YAML 语法
python -c "import yaml; yaml.safe_load(open('/tmp/local-config-with-nickname.yml'))" || echo "YAML 语法错误"
```

#### 步骤 3: 创建配置文件
```bash
# 创建全局配置文件
cat > .gfl.config.yml << 'EOF'
debug: false
devBaseBranch: "develop"
productionBranch: "main"
nickname: ""
featurePrefix: "feature"
fixPrefix: "fix"
hotfixPrefix: "hotfix"
EOF

# 创建本地配置文件（带昵称）
cat > .gfl.config.local.yml << 'EOF'
nickname: "aric"
debug: false
EOF

# 设置文件权限（遵循 umask）
chmod 644 .gfl.config.yml
chmod 644 .gfl.config.local.yml
```

#### 步骤 4: Git 忽略配置
```bash
# 检查 .gitignore 是否已存在本地配置文件
if ! grep -q ".gfl.config.local.yml" .gitignore 2>/dev/null; then
    # 添加本地配置文件到 .gitignore
    echo "" >> .gitignore
    echo "# GFL local configuration" >> .gitignore
    echo ".gfl.config.local.yml" >> .gitignore
fi
```

#### 步骤 5: 完整执行示例
```bash
# 用户执行
$ gfl init --nickname aric

# GFL 内部执行序列:
# 1. 检查 Git 仓库
$ git rev-parse --is-inside-work-tree
true

# 2. 检查现有配置文件
$ ls -la .gfl.config*
ls: cannot access '.gfl.config*': No such file or directory

# 3. 创建全局配置文件
$ cat > .gfl.config.yml << 'EOF'
debug: false
devBaseBranch: "develop"
productionBranch: "main"
nickname: ""
featurePrefix: "feature"
fixPrefix: "fix"
hotfixPrefix: "hotfix"
EOF

# 4. 创建本地配置文件（带昵称）
$ cat > .gfl.config.local.yml << 'EOF'
nickname: "aric"
debug: false
EOF

# 5. 设置文件权限
$ chmod 644 .gfl.config.yml
$ chmod 644 .gfl.config.local.yml

# 6. 更新 .gitignore
$ echo "" >> .gitignore
$ echo "# GFL local configuration" >> .gitignore
$ echo ".gfl.config.local.yml" >> .gitignore

# 7. 验证配置文件
$ cat .gfl.config.yml
debug: false
devBaseBranch: "develop"
productionBranch: "main"
nickname: ""
featurePrefix: "feature"
fixPrefix: "fix"
hotfixPrefix: "hotfix"

$ cat .gfl.config.local.yml
nickname: "aric"
debug: false

# GFL 输出成功信息
# Configuration initialized successfully!
# Global config: .gfl.config.yml
# Local config: .gfl.config.local.yml
```

#### 强制重新初始化场景
```bash
# 用户执行
$ gfl init --force --nickname bob

# GFL 内部执行序列:
# 1. 检查现有配置文件
$ ls -la .gfl.config*
-rw-r--r-- 1 user user 234 Dec 20 10:30 .gfl.config.yml
-rw-r--r-- 1 user user 45  Dec 20 10:30 .gfl.config.local.yml

# 2. 由于使用了 --force，覆盖本地配置文件
$ cat > .gfl.config.local.yml << 'EOF'
nickname: "bob"
debug: false
EOF

# 3. 全局配置文件不被强制覆盖（保护机制）
$ echo "Global config file preserved (not forced)"

# 4. 输出结果
# Local configuration file overwritten successfully!
# Global configuration file preserved.
```

#### 错误处理场景

##### 场景 1: 配置文件已存在（未使用 --force）
```bash
# 用户执行
$ gfl init --nickname new-user

# GFL 检测到现有文件
$ ls -la .gfl.config.local.yml
-rw-r--r-- 1 user user 45 Dec 20 10:30 .gfl.config.local.yml

# GFL 显示错误
# Error: Local configuration file already exists: .gfl.config.local.yml
# Use --force to overwrite existing configuration.
```

##### 场景 2: 不是 Git 仓库
```bash
# GFL 检查仓库状态
$ git rev-parse --is-inside-work-tree
fatal: not a git repository (or any of the parent directories): .git

# GFL 显示错误
# Error: Not a Git repository. Please run 'git init' first.
```

##### 场景 3: 权限不足
```bash
# GFL 尝试创建文件但权限不足
$ cat > .gfl.config.yml << 'EOF'
debug: false
EOF
bash: .gfl.config.yml: Permission denied

# GFL 显示错误
# Error: Cannot create configuration files. Permission denied.
```

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