# Release
> 版本发布流程

## steps
1. 新建 release 分支
2. 合并到 alpha/beta/main 分支，完成各个环境的测试
3. 打 tag

```bash
# 简写: gfl rls
gfl release

# 合并代码到各环境，完成发布，并提醒测试
# ----------- 以下操作，确认你是在最新的 releases/release-vx.y.z  分支上 ----------
gfl pr alpha
gfl pr beta
gfl pr main

# 打 tag，没有特别说明，会根据上个版本号自动生成
gfl tag
```
