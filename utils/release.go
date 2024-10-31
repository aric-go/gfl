package utils

func GetLatestReleaseBranch() string {
	version := GetLatestVersion()
	return "releases/release-" + version
}
