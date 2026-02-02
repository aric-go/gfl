package cmd

import (
	"fmt"
	"gfl/utils"
	"gfl/utils/strings"
	str "strings"
	"github.com/afeiship/go-box"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:     "info",
	Aliases: []string{"i"},
	Short:   "Display repository info (alias: i)", // Will be updated after strings load
	Run: func(cmd *cobra.Command, args []string) {
		displayInfo()
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}

func displayInfo() {
	// Get branch information
	info, err := utils.GetBranchInfo()
	if err != nil {
		utils.Errorf("Failed to get repository info: %v", err)
		return
	}

	// Build info lines for display
	lines := buildInfoLines(info)

	// Display using ASCII box
	box.PrintASCIIBox(lines)
	fmt.Println()
}

func buildInfoLines(info *utils.BranchInfo) []string {
	lines := []string{}

	// Current Branch
	lines = append(lines, fmt.Sprintf("📍 %s: %s",
		strings.GetPath("info.current_branch"),
		info.CurrentBranch))

	// Tracking Branch
	trackingLabel := strings.GetPath("info.not_tracking")
	if info.TrackingBranch != "" {
		trackingLabel = info.TrackingBranch
	}
	lines = append(lines, fmt.Sprintf("🔗 %s: %s",
		strings.GetPath("info.tracking"),
		trackingLabel))

	// Position (ahead/behind)
	positionLabel := "-"
	if info.TrackingBranch != "" {
		if info.AheadCommits == 0 && info.BehindCommits == 0 {
			positionLabel = strings.GetPath("info.up_to_date")
		} else {
			parts := []string{}
			if info.AheadCommits > 0 {
				parts = append(parts, fmt.Sprintf(strings.GetPath("info.ahead"), info.AheadCommits))
			}
			if info.BehindCommits > 0 {
				parts = append(parts, fmt.Sprintf(strings.GetPath("info.behind"), info.BehindCommits))
			}
			if len(parts) > 0 {
				positionLabel = str.Join(parts, ", ")
			}
		}
	}
	lines = append(lines, fmt.Sprintf("📊 %s: %s",
		strings.GetPath("info.position"),
		positionLabel))

	// Working Directory Status
	workingDirLabel := strings.GetPath("info.working_dir_clean")
	if !info.WorkingDirClean {
		workingDirLabel = strings.GetPath("info.working_dir_dirty")
	}
	lines = append(lines, fmt.Sprintf("✅ %s: %s",
		strings.GetPath("info.working_dir"),
		workingDirLabel))

	// Remote Repository
	remoteLabel := info.RemoteURL
	if remoteLabel == "" {
		remoteLabel = strings.GetPath("info.not_configured")
	}
	lines = append(lines, fmt.Sprintf("🌐 %s: %s",
		strings.GetPath("info.remote"),
		remoteLabel))

	return lines
}
