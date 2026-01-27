package cmd

import (
  "fmt"
  "gfl/utils"
  "gfl/utils/strings"
  "os/exec"
  str "strings"

  "github.com/fatih/color"
  "github.com/spf13/cobra"
)

var (
  localFlag    bool
  remoteFlag   bool
  exactFlag    bool
  forceFlag    bool
)

var sweepCmd = &cobra.Command{
  Use:     "sweep [keyword]",
  Aliases: []string{"clean", "rm"},
  Short:   "Clean branches containing specific keywords (alias: clean, rm)",
  Args:    cobra.ExactArgs(1), // 需要一个关键词参数
  Run: func(cmd *cobra.Command, args []string) {
    keyword := args[0]
    // get flag confirm
    confirm, _ := cmd.Flags().GetBool("confirm")

    // 如果没有设置本地或远程标志，打印错误并返回
    if !localFlag && !remoteFlag {
      utils.Error(strings.GetPath("sweep.local_remote_required"))
      return
    }

    if localFlag {
      // 清理本地分支
      cleanLocalBranches(keyword, confirm, exactFlag, forceFlag)
    }

    if remoteFlag {
      // 清理远程分支
      cleanRemoteBranches(keyword, confirm, exactFlag)
    }

    if !confirm {
      utils.Info(strings.GetPath("sweep.skip_confirm"))
    }
  },
}

func cleanLocalBranches(keyword string, confirm bool, exactMatch bool, force bool) {
  // 获取本地分支列表
  branches, err := exec.Command("git", "branch").Output()
  if err != nil {
    utils.Errorf(strings.GetPath("sweep.local_branches_error", err))
    return
  }

  // 遍历本地分支列表并删除匹配关键词的分支
  for _, branch := range str.Split(string(branches), "\n") {
    branch = str.TrimSpace(branch) // 去除空格
    if branch == "" {
      continue // 跳过空行
    }

    // 根据精确匹配标志选择匹配方式
    var shouldDelete bool
    if exactMatch {
      shouldDelete = branch == keyword
    } else {
      shouldDelete = str.Contains(branch, keyword)
    }

    if shouldDelete {
      // 执行命令: git branch -d branch-name (安全删除) 或 git branch -D branch-name (强制删除)
      deleteFlag := "-d"
      if force {
        deleteFlag = "-D"
      }
      command := fmt.Sprintf("git branch %s %s", deleteFlag, branch)
      if confirm {
        if err := utils.RunCommandWithSpin(command, strings.GetPath("sweep.deleting_local")); err != nil {
          utils.Errorf(strings.GetPath("sweep.delete_local_error", branch, err))
        } else {
          utils.Successf(strings.GetPath("sweep.delete_local_success", branch))
        }
      } else {
        logRemove(branch, keyword)
      }
    }
  }
}

func cleanRemoteBranches(keyword string, confirm bool, exactMatch bool) {
  // 获取远程分支列表
  branches, err := exec.Command("git", "branch", "-r").Output()
  if err != nil {
    utils.Errorf(strings.GetPath("sweep.remote_branches_error", err))
    return
  }

  // 遍历远程分支列表并删除匹配关键词的分支
  for _, branch := range str.Split(string(branches), "\n") {
    branch = str.TrimSpace(branch) // 去除空格
    if branch == "" {
      continue // 跳过空行
    }

    // 提取分支名称（去掉远程名）
    remoteBranch := str.TrimPrefix(branch, "origin/")

    // 根据精确匹配标志选择匹配方式
    var shouldDelete bool
    if exactMatch {
      shouldDelete = remoteBranch == keyword
    } else {
      shouldDelete = str.Contains(branch, keyword)
    }

    if shouldDelete {
      command := fmt.Sprintf("git push origin --delete %s", remoteBranch)
      if confirm {
        if err := utils.RunCommandWithSpin(command, strings.GetPath("sweep.deleting_remote")); err != nil {
          utils.Errorf(strings.GetPath("sweep.delete_remote_error", branch, err))
        } else {
          utils.Successf(strings.GetPath("sweep.delete_remote_success", branch))
        }
      } else {
        logRemove(branch, keyword)
      }
    }
  }
}

func logRemove(branch string, keyword string) {
  colorBranch := color.GreenString(branch)
  colorKeyword := color.RedString(keyword)
  // list branches without confirm
  msg := strings.GetPath("sweep.manual_delete", colorBranch, colorKeyword)
  utils.Infof(msg)
}

func init() {
  sweepCmd.Flags().BoolVarP(&localFlag, "local", "l", false, strings.GetPath("sweep.local_flag"))
  sweepCmd.Flags().BoolVarP(&remoteFlag, "remote", "r", false, strings.GetPath("sweep.remote_flag"))
  sweepCmd.Flags().BoolVarP(&exactFlag, "exact", "e", false, strings.GetPath("sweep.exact_flag"))
  sweepCmd.Flags().BoolVarP(&forceFlag, "force", "f", false, strings.GetPath("sweep.force_flag"))
  rootCmd.AddCommand(sweepCmd)
}
