# gfl 
> 重命名分支

## steps
```sh
# step1 - 重命名本地分支
git branch -m fix/kat-9722 fix/kat-9723
# step2 - 删除远程分支
git push origin --delete fix/kat-9722
# step3 - 推送本地分支到远程分支
git push origin -u fix/kat-9723
```

## flags
- `-l`, `--local`: 本地分支
- `-r`, `--remote`: 远程分支
- 是否删除远程分支 `-d`, `--delete`: 删除远程分支

## 注意
- 编译使用 `npm run build`, 最终产物在 `dist` 目录下
