# Bugfix
> 开始开发时的一些bug修复

## steps
1. 建新分支
2. 发布 bugfix 分支
3. 开发修复bug
4. 提PR，等待review
5. 合并PR

```bash
# 新建分支
gfl start fix:v1030/mobile-login-failed

# 注意：以下步骤都在新分支上进行
gfl publish

# 开发需求
# 开发完成后，提交PR
gfl pr
```