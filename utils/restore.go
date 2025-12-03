package utils

import (
	"fmt"
	"gfl/utils/strings"
	"os"
	"os/exec"
	"path/filepath"
	str "strings"
)

// RestorePath restores a file or directory to its HEAD state
func RestorePath(path string, confirm bool) error {
	// 检查路径是否存在
	if _, err := os.Stat(path); os.IsNotExist(err) {
		Errorf(strings.GetString("restore", "path_not_exist", path))
		return err
	}

	// 检查是否是 git 仓库
	if !isGitRepository() {
		Error(strings.GetString("restore", "not_git_repo"))
		return fmt.Errorf("not a git repository")
	}

	// 构建 git restore 命令
	command := fmt.Sprintf("git restore --source=HEAD --staged --worktree -- %s", path)

	if confirm {
		// 执行实际的恢复操作
		if err := RunCommandWithSpin(command, strings.GetString("restore", "restoring", path)); err != nil {
			Errorf(strings.GetString("restore", "restore_error", path, err))
			return err
		}
		Successf(strings.GetString("restore", "restore_success", path))
	} else {
		// Dry-run 模式，显示将要执行的操作
		LogRestore(path)
	}

	return nil
}

// isGitRepository checks if the current directory is a git repository
func isGitRepository() bool {
	_, err := exec.Command("git", "rev-parse", "--git-dir").Output()
	return err == nil
}

// LogRestore logs what would be restored in dry-run mode
func LogRestore(path string) {
	// 检查路径是否有变化
	hasChanges, err := checkPathChanges(path)
	if err != nil {
		Errorf(strings.GetString("restore", "check_changes_error", path, err))
		return
	}

	if !hasChanges {
		Infof(strings.GetString("restore", "no_changes", path))
		return
	}

	absPath, _ := filepath.Abs(path)
	Infof(strings.GetString("restore", "would_restore", absPath))
}

// checkPathChanges checks if a path has any changes compared to HEAD
func checkPathChanges(path string) (bool, error) {
	// 检查工作区是否有变化
	worktreeCmd := exec.Command("git", "diff", "--quiet", "--", path)
	if worktreeCmd.Run() != nil {
		return true, nil // 有工作区变化
	}

	// 检查暂存区是否有变化
	stagedCmd := exec.Command("git", "diff", "--cached", "--quiet", "--", path)
	if stagedCmd.Run() != nil {
		return true, nil // 有暂存区变化
	}

	// 检查是否是未被跟踪的新文件
	lsFilesCmd := exec.Command("git", "ls-files", "--others", "--exclude-standard", "--", path)
	output, err := lsFilesCmd.Output()
	if err == nil && len(str.TrimSpace(string(output))) > 0 {
		return true, nil // 是未跟踪的新文件
	}

	return false, nil // 没有变化
}