package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"

	"gfl/utils/strings"
)

// rootVersionPattern matches the rootCmd Version field in cmd/root.go
var rootVersionPattern = regexp.MustCompile(`Version:\s*"[^"]*"`)

// ResolveBumpArgs normalizes bump arguments: type falls back to "patch";
// prerelease-like types require a preid, and release types with a preid are
// mapped to their prerelease variant (eg. patch + beta -> prepatch).
func ResolveBumpArgs(flagType, flagPreid string) (string, string, error) {
	bumpType := flagType
	if bumpType == "" {
		bumpType = "patch"
	}

	preid := flagPreid
	if preid == "" {
		// prerelease-like types require a preid to produce e.g. 1.0.0-beta.1
		switch bumpType {
		case "premajor", "preminor", "prepatch", "prerelease":
			return "", "", fmt.Errorf("type %q requires --preid", bumpType)
		}
	} else {
		// npm ignores --preid for release types; map to the prerelease variant
		// so `patch --preid beta` produces 1.0.1-beta.0 instead of a plain 1.0.1
		switch bumpType {
		case "major":
			bumpType = "premajor"
		case "minor":
			bumpType = "preminor"
		case "patch":
			bumpType = "prepatch"
		}
	}

	return bumpType, preid, nil
}

// ReadPkgVersion returns the version field of package.json in the current directory.
func ReadPkgVersion() (string, error) {
	data, err := os.ReadFile("package.json")
	if err != nil {
		return "", fmt.Errorf("failed to read package.json: %w", err)
	}
	var pkg struct {
		Version string `json:"version"`
	}
	if err := json.Unmarshal(data, &pkg); err != nil {
		return "", fmt.Errorf("failed to parse package.json: %w", err)
	}
	if pkg.Version == "" {
		return "", fmt.Errorf("version not found in package.json")
	}
	return pkg.Version, nil
}

// syncRootVersion writes the new version into the Version field of cmd/root.go.
func syncRootVersion(version string) error {
	const path = "cmd/root.go"
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read %s: %w", path, err)
	}
	if !rootVersionPattern.Match(data) {
		return fmt.Errorf("version field not found in %s", path)
	}
	updated := rootVersionPattern.ReplaceAll(data, fmt.Appendf(nil, `Version: "%s"`, version))
	if err := os.WriteFile(path, updated, 0644); err != nil {
		return fmt.Errorf("failed to write %s: %w", path, err)
	}
	return nil
}

// RunBump bumps the version via pnpm, syncs cmd/root.go, then commits both files (no push).
func RunBump(bumpType, preid string) error {
	args := []string{"version", bumpType, "--no-git-tag-version"}
	if preid != "" {
		args = append(args, "--preid", preid)
	}
	if err := RunCommandWithArgs("pnpm", args, strings.GetPath("bump.bumping")); err != nil {
		return fmt.Errorf("version bump failed: %w", err)
	}

	version, err := ReadPkgVersion()
	if err != nil {
		return err
	}
	if err := syncRootVersion(version); err != nil {
		return err
	}
	Successf(strings.GetPath("bump.updated", version))

	if err := RunCommandWithArgs("git", []string{"add", "package.json", "cmd/root.go"}, strings.GetPath("bump.committing")); err != nil {
		return err
	}
	msg := fmt.Sprintf("chore(release): v%s [gfl]", version)
	if err := RunCommandWithArgs("git", []string{"commit", "-m", msg}, strings.GetPath("bump.committing")); err != nil {
		return err
	}
	return nil
}
