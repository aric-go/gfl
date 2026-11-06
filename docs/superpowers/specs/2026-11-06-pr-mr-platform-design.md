# 设计：pr 命令支持 GitLab MR（platform 配置项）

日期：2026-11-06

## 背景

PR（GitHub）与 MR（GitLab）作用相同。gfl 的 `pr` 命令目前只支持 GitHub 的 compare URL。决定共用同一个 `pr` 命令，通过新增配置项 `platform` 切换平台。

## 设计

### 1. 新配置项 `platform`

- `YamlConfig` 新增 `Platform string` + `PlatformSet bool`（`yaml:"-"`），遵循现有 `*Set` 模式
- `loadConfigFile` 增加 `v.IsSet("platform")` 检查；`mergeConfig` 增加合并分支
- 默认值 `github`，可选 `github` / `gitlab`
- 无效值：打印警告并回退 `github`

### 2. `utils/pr.go` 改造

`CreatePr(base, head)` 按 platform 生成 URL 并打开浏览器：

- github: `https://github.com/{repo}/compare/{base}...{head}?expand=1`
- gitlab: `https://gitlab.com/{repo}/-/merge_requests/new?merge_request[source_branch]={head}&merge_request[target_branch]={base}`

GitLab 的 source/target 分支参数使用 `net/url.QueryEscape` 编码（分支名含 `/` 时必须，如 `feature/aric/xxx`）。

### 3. `pr -o` 列表页适配

- github: `https://github.com/{repo}/pulls`
- gitlab: `https://gitlab.com/{repo}/-/merge_requests`

### 4. 配置展示与文档

- `cmd/config.go` 展示表格新增 `platform` 行（含来源颜色）
- `strings.yml` 中英文新增 `pr.platform_invalid`；`pr.short` 更新为提及 MR
- `CLAUDE.md` 配置字段表新增 `platform` 行

## 不做的事

- 不新建独立 `mr` 命令
- 不做远程 host 自动检测
- 不支持 Bitbucket 等其他平台

## 验证

- `npm run build` 编译通过
- `./dist/gfl config` 显示 platform 行
- 配置 `platform: gitlab` 后 `./dist/gfl pr --debug` 打开的 URL 为 GitLab MR 格式（人工核对 debug 输出中的 URL）
