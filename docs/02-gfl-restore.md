# restore
> 恢复文件到未修改之前


## steps
```sh
gr() {
    if [ $# -eq 0 ]; then
        # 作用于当前目录，丢弃所有本地变更（包括已暂存的）
        git restore --source=HEAD --staged --worktree -- .
    else
        # 作用于指定路径，丢弃所有本地变更（包括已暂存的）
        git restore --source=HEAD --staged --worktree -- "$@"
    fi
}
```

## features
1. 接收文件 或者目录路径作为参数，默认作用于当前目录
2. 丢弃所有本地变更（包括已暂存的）
