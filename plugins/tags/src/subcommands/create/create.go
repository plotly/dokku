package main

import (
	"errors"
    "fmt"
	"flag"
	"strconv"
    "strings"

	common "github.com/dokku/dokku/plugins/common"
    sh "github.com/codeskyblue/go-sh"
)

// creates an images tag for app via command line
func main() {
	flag.Parse()
	appName := flag.Arg(1)
	imageTag := flag.Arg(2)
	if appName == "" {
		common.LogFail("Please specify an app to run the command on")
	}

	err := common.VerifyAppName(appName)
	if err != nil {
		common.LogFail(err.Error())
	}

	imageRepo := common.GetAppImageRepo(appName)
	sourceImage := fmt.Sprintf("%s:latest", imageRepo)
	targetImage := fmt.Sprintf("%s:%s", imageRepo, imageTag)

	isTagForceAvailable, err := isTagForceAvailable()
	if err != nil {
		common.LogFail(err.Error())
	}

	tagImageArgs := []string{"tag", sourceImage, targetImage}
	if isTagForceAvailable {
		tagImageArgs = []string{"tag", "-f", sourceImage, targetImage}
	}
	tagImageCmd := common.NewShellCmdWithArgs("docker", tagImageArgs)
	tagImageCmd.Execute()

	common.LogInfoQuiet2(fmt.Sprintf("Added %s tag to %s", imageTag, imageRepo))

	triggerArgs := []string{"trigger", "tags-create", appName, imageTag}
	triggerCmd := common.NewShellCmdWithArgs("plugn", triggerArgs)
	triggerCmd.Execute()
}

// "docker tag -f" was dropped in 1.12.0
func isTagForceAvailable() (bool, error) {
	clientVersionString := dockerVersion()
	if clientVersionString == "" {
		return false, errors.New("Unable to retrieve docker version")
	}

	items := strings.Split(clientVersionString, ".")
	majorVersion, err := strconv.Atoi(items[0])
	if err != nil {
		return false, errors.New("Unable to parse docker version")
	}

	minorVersion, err := strconv.Atoi(items[1])
	if err != nil {
		return false, errors.New("Unable to parse docker version")
	}

	if majorVersion > 1 {
		return false, nil
	}
	if majorVersion == 1 && minorVersion > 11 {
		return false, nil
	}

	return true, nil
}

func dockerVersion() string {
    b, err := sh.Command("docker", "version", "-f=\"{{ .Client.Version }}\"").Output()
    if err != nil {
        common.LogFail(err.Error())
    }
    return string(b[:])
}
