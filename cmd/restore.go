package cmd

import (
  "gfl/utils"
  "gfl/utils/strings"

  "github.com/spf13/cobra"
)

var restoreCmd = &cobra.Command{
  Use:     "restore [path...]",
  Aliases: []string{"r"},
  Short:   "Restore files to unmodified state", // Will be updated after strings load
  Args:    cobra.MinimumNArgs(0),               // 可以接受 0 个或多个参数
  Run: func(cmd *cobra.Command, args []string) {
    // get flag confirm
    confirm, _ := cmd.Flags().GetBool("confirm")

    if len(args) == 0 {
      // 没有参数时，作用于当前目录
      utils.RestorePath(".", confirm)
    } else {
      // 有参数时，作用于指定的路径
      for _, path := range args {
        utils.RestorePath(path, confirm)
      }
    }

    if !confirm {
      utils.Info(strings.GetString("restore", "skip_confirm"))
    }
  },
}

func init() {
  rootCmd.AddCommand(restoreCmd)
}
