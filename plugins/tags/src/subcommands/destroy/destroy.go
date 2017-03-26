package main

import (
	"flag"

	common "github.com/dokku/dokku/plugins/common"
)

// destroys an app image tag via command line
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
	if imageTag == "latest" {
		common.LogFail(fmt.Sprintf("You can't remove internal dokku tag (%s) for %s", imageTag, imageRepo))
	}

	dockerArgs := []string{"rmi", fmt.Sprintf("%s:%s", imageRepo, imageTag)}
	dockerCmd := common.NewShellCmdWithArgs("docker", triggerArgs)
	dockerCmd.Execute()

	triggerArgs := []string{"trigger", "tags-destroy", appName, imageTag}
	triggerCmd := common.NewShellCmdWithArgs("plugn", triggerArgs)
	triggerCmd.Execute()
}
