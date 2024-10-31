package utils

func GetLatestReleaseBranch() string {
	version := GetLatestVersion()
	return "releases/releases-" + version
}
