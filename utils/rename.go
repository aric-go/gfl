package utils

import (
	"fmt"
	"gfl/utils/strings"
	"os/exec"
	str "strings"

	"github.com/fatih/color"
)

// RenameLocalBranch renames a local branch
func RenameLocalBranch(oldBranch string, newBranch string, confirm bool) error {
	// 检查当前分支是否是要重命名的分支
	currentBranch, err := exec.Command("git", "branch", "--show-current").Output()
	if err != nil {
		return fmt.Errorf(strings.GetPath("rename.get_current_branch_error", err))
	}

	currentBranchName := str.TrimSpace(string(currentBranch))

	// 如果当前分支就是要重命名的分支，先切换到其他分支（通常是main）
	if currentBranchName == oldBranch {
		Info(strings.GetPath("rename.switching_from_current"))
		// 尝试切换到main分支，如果main不存在则切换到master
		targetBranch := "main"
		if err := exec.Command("git", "checkout", targetBranch).Run(); err != nil {
			targetBranch = "master"
			if err := exec.Command("git", "checkout", targetBranch).Run(); err != nil {
				return fmt.Errorf(strings.GetPath("rename.switch_branch_error", targetBranch, err))
			}
		}
		Successf(strings.GetPath("rename.switched_to_branch", targetBranch))
	}

	// 执行重命名命令: git branch -m old-branch new-branch
	command := fmt.Sprintf("git branch -m %s %s", oldBranch, newBranch)
	if confirm {
		if err := RunCommandWithSpin(command, strings.GetPath("rename.renaming_local")); err != nil {
			return fmt.Errorf(strings.GetPath("rename.rename_local_error", oldBranch, newBranch, err))
		}
		Successf(strings.GetPath("rename.rename_local_success", oldBranch, newBranch))
	} else {
		LogRename(oldBranch, newBranch, "local")
	}

	return nil
}

// HandleRemoteBranch handles remote branch operations
func HandleRemoteBranch(oldBranch string, newBranch string, deleteOld bool, confirm bool) error {
	if deleteOld {
		// 删除远程旧分支
		if err := DeleteRemoteBranch(oldBranch, confirm); err != nil {
			return err
		}
	}

	// 推送新分支到远程
	return PushNewBranch(newBranch, confirm)
}

// DeleteRemoteBranch deletes a remote branch
func DeleteRemoteBranch(branch string, confirm bool) error {
	command := fmt.Sprintf("git push origin --delete %s", branch)
	if confirm {
		if err := RunCommandWithSpin(command, strings.GetPath("rename.deleting_remote")); err != nil {
			return fmt.Errorf(strings.GetPath("rename.delete_remote_error", branch, err))
		}
		Successf(strings.GetPath("rename.delete_remote_success", branch))
	} else {
		LogAction("delete", branch, "remote")
	}

	return nil
}

// PushNewBranch pushes a new branch to remote
func PushNewBranch(branch string, confirm bool) error {
	command := fmt.Sprintf("git push origin -u %s", branch)
	if confirm {
		if err := RunCommandWithSpin(command, strings.GetPath("rename.pushing_remote")); err != nil {
			return fmt.Errorf(strings.GetPath("rename.push_remote_error", branch, err))
		}
		Successf(strings.GetPath("rename.push_remote_success", branch))
	} else {
		LogAction("push", branch, "remote")
	}

	return nil
}

// LogRename logs a rename operation for dry-run mode
func LogRename(oldBranch string, newBranch string, scope string) {
	colorOld := color.RedString(oldBranch)
	colorNew := color.GreenString(newBranch)
	colorScope := color.CyanString(scope)
	msg := strings.GetPath("rename.manual_rename", colorOld, colorNew, colorScope)
	Infof(msg)
}

// LogAction logs an action for dry-run mode
func LogAction(action string, branch string, scope string) {
	colorBranch := color.GreenString(branch)
	colorAction := color.YellowString(action)
	colorScope := color.CyanString(scope)
	msg := strings.GetPath("rename.manual_action", colorAction, colorBranch, colorScope)
	Infof(msg)
}