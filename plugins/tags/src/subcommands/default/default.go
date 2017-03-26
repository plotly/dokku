package main

import (
	"flag"
    "fmt"

	common "github.com/dokku/dokku/plugins/common"
)

// shows docker images tags for app via command line
func main() {
	flag.Parse()
	appName := flag.Arg(1)
	if appName == "" {
		common.LogFail("Please specify an app to run the command on")
	}

	err := common.VerifyAppName(appName)
	if err != nil {
		common.LogFail(err.Error())
	}

	imageRepo := common.GetAppImageRepo(appName)

	common.LogInfoQuiet2(fmt.Sprintf("Image tags for %s", imageRepo))
	dockerArgs := []string{"images", imageRepo}
	dockerCmd := common.NewShellCmdWithArgs("docker", dockerArgs)
	dockerCmd.Execute()
}
