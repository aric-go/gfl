# Hotfix
> 解决线上问题的临时 bug

## steps
1. 建新 hotfix分支
2. 发布 hotfix 分支
3. 开发修复bug
4. 提PR，等待review
5. 合并PR

```bash
# 新建分支
gfl hotfix online-bug-fix

# 注意：以下步骤都在新分支上进行
gfl publish

# 开发需求
# 开发完成后，提交PR到当前的预生产分支
gfl pr alpha
```