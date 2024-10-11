#!/bin/bash

# 读取当前版本
version=$(cat VERSION)

# 使用 awk 增加小版本号
new_version=$(echo "$version" | awk -F. -v OFS=. '{$NF++;print}')

# 更新版本文件
echo "$new_version" > VERSION

echo "版本已更新为: $new_version"
