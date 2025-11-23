# GFL 最佳实践

本文档提供了使用 GFL 的最佳实践和建议，帮助团队提高开发效率。

## 团队协作

### 统一配置规范

1. **创建团队标准配置**
   ```yaml
   # .gfl.config.yml
   devBaseBranch: develop
   productionBranch: main
   nickname: ""  # 团队成员在本地配置中设置
   featurePrefix: feature
   fixPrefix: fix
   hotfixPrefix: hotfix
   ```

2. **个人配置模板**
   ```yaml
   # .gfl.config.local.yml.template
   nickname: your-github-username
   debug: false
   ```

3. **版本控制**
   - 提交 `.gfl.config.yml` 到版本控制
   - 将 `.gfl.config.local.yml` 添加到 `.gitignore`
   - 为新团队成员提供配置模板

### 分支命名策略

#### 推荐的命名格式

1. **功能分支**
   ```bash
   # 包含功能描述
   gfl s user-authentication

   # 包含 JIRA 票号
   gfl s PROJ-123-user-authentication
   ```

2. **热修复分支**
   ```bash
   # 描述紧急修复内容
   gfl hf security-vulnerability-fix

   # 包含版本号
   gfl hf v1.2.1-memory-leak
   ```

3. **Bug 修复分支**
   ```bash
   # 描述具体问题
   gfl s fix-login-validation-error
   ```

#### 分支命名最佳实践

- 使用描述性的名称
- 避免特殊字符（除连字符外）
- 保持简洁但信息完整
- 使用英文命名（便于协作）
- 包含相关票号（如 JIRA、GitHub Issues）

## 工作流程

### 功能开发流程

```bash
# 1. 同步最新代码
gfl sync

# 2. 开始新功能
gfl s feature-name

# 3. 开发过程中定期发布
gfl p

# 4. 功能完成，创建 PR
gfl pr --open

# 5. 代码审查通过后合并
# 6. 清理本地分支
gfl sweep feature-name --local --confirm
```

### 紧急修复流程

```bash
# 1. 创建热修复分支
gfl hf critical-security-fix

# 2. 快速修复问题
# 3. 发布修复
gfl p

# 4. 创建紧急 PR
gfl pr main --open

# 5. 合并后创建补丁版本
gfl release --hotfix
gfl tag --type patch
```

### 版本发布流程

```bash
# 1. 确保所有功能已合并
# 2. 创建发布版本
gfl release --type minor

# 3. 创建版本标签
gfl tag --type minor

# 4. 清理已完成的开发分支
gfl sweep feature --confirm
```

## 代码质量管理

### 提交信息规范

配合 GFL 使用良好的 Git 提交规范：

```bash
# 功能开发
git commit -m "feat: add user authentication system"

# Bug 修复
git commit -m "fix: resolve login validation error"

# 热修复
git commit -m "hotfix: patch security vulnerability"

# 文档更新
git commit -m "docs: update API documentation"
```

### Pull Request 规范

1. **PR 标题格式**
   - 功能：`feat: 添加用户认证功能`
   - 修复：`fix: 修复登录验证错误`
   - 热修复：`hotfix: 修复安全漏洞`

2. **PR 描述模板**
   ```markdown
   ## 变更类型
   - [ ] 新功能
   - [ ] Bug 修复
   - [ ] 热修复
   - [ ] 文档更新

   ## 变更描述
   简要描述本次变更的内容

   ## 测试
   - [ ] 单元测试通过
   - [ ] 集成测试通过
   - [ ] 手动测试完成

   ## 检查清单
   - [ ] 代码符合团队规范
   - [ ] 已更新相关文档
   - [ ] 测试覆盖率达标
   ```

## 自动化集成

### Shell 自动补全

```bash
# 安装自动补全
echo 'source <(gfl completion bash)' >> ~/.bashrc
# 或
echo 'source <(gfl completion zsh)' >> ~/.zshrc
```

### Git Hooks 配置

使用 Git Hooks 确保 GFL 使用规范：

```bash
# pre-commit hook
#!/bin/bash
# 确保分支名称符合规范
current_branch=$(git branch --show-current)
if [[ ! $current_branch =~ ^(feature|fix|hotfix)/ ]]; then
    echo "错误：分支名称不符合规范，请使用 gfl start 创建分支"
    exit 1
fi
```

### CI/CD 集成

在 CI/CD 流水线中使用 GFL：

```yaml
# GitHub Actions 示例
- name: Setup GFL
  run: |
    curl -L https://github.com/your-repo/gfl/releases/latest/download/gfl-linux-amd64 -o gfl
    chmod +x gfl
    sudo mv gfl /usr/local/bin/

- name: Release
  run: |
    gfl tag --type ${{ github.event.inputs.version_type }}
    gfl release --type ${{ github.event.inputs.version_type }}
```

## 性能优化

### 大型仓库优化

```bash
# 定期同步远程仓库
gfl sync

# 在网络良好环境下操作以确保获取最新代码
```

### 分支清理策略

```bash
# 定期清理已合并的功能分支
gfl sweep feature --confirm

# 清理远程已合并分支
gfl sweep --remote --confirm

# 清理旧的发布分支
gfl sweep release --confirm
```

## 安全考虑

### 敏感信息保护

1. **配置文件安全**
   ```bash
   # 设置正确的文件权限
   chmod 644 .gfl.config.yml
   chmod 600 .gfl.config.local.yml
   ```

2. **不提交敏感信息**
   ```bash
   # .gitignore
   .gfl.config.local.yml
   *.key
   *.pem
   ```

### 访问控制

确保团队成员具有适当的 Git 仓库权限：

- 开发分支：读写权限
- 生产分支：受保护的写入权限
- 发布权限：仅限核心团队成员

## 监控和日志

### 调试模式

启用调试模式排查问题：

```yaml
# .gfl.config.local.yml
debug: true
```

### 操作日志

```bash
# 查看最近的操作历史
gfl history

# 查看特定命令的执行详情
gfl history --command start
```

## 故障排除

### 常见问题解决方案

1. **分支冲突**
   ```bash
   # 同步最新代码
   gfl sync
   # 重新基于最新代码
   git rebase dev
   ```

2. **远程分支不存在**
   ```bash
   # 发布当前分支
   gfl publish
   # 然后再创建 PR
   gfl pr
   ```

3. **配置问题**
   ```bash
   # 重置配置
   gfl init --force
   # 验证配置
   gfl config validate
   ```

## 团队培训

### 新成员入门流程

1. **环境准备**
   ```bash
   # 安装 GFL
   go install github.com/your-repo/gfl@latest

   # 初始化配置
   gfl init --nickname new-member
   ```

2. **培训内容**
   - 基本命令使用
   - 团队工作流程
   - 分支命名规范
   - PR 创建流程

3. **参考资料**
   - [快速开始指南](quick-start.md)
   - [完整命令文档](commands.md)
   - 团队内部 Wiki

### 定期回顾

定期回顾和优化 GFL 使用流程：

- 每月检查配置文件是否需要更新
- 收集团队反馈，改进工作流程
- 更新最佳实践文档
- 培训新团队成员

通过遵循这些最佳实践，您的团队可以充分发挥 GFL 的潜力，提高开发效率和代码质量。